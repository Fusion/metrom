package net

import (
	"time"
)

type HopStats struct {
	ElapsedMin        int64
	ElapsedMax        int64
	History           [5]int64
	HistoryWriterHead int
	HistoryReaderHead int
	JitterMin         int64
	JitterMax         int64
	PingTotal         int64
	PingMiss          int64
}
type HopHandler struct {
	local        Local
	remote       Remote
	connbehavior ConnBehavior
	attemptsleft int
	udpwriter    *UdpWriter
	start        time.Time
	donotlinger  chan bool
	HopStats     HopStats
}

func NewHopHandler(local Local, remote Remote, connbehavior ConnBehavior) HopHandler {
	return HopHandler{
		local:        local,
		remote:       remote,
		connbehavior: connbehavior,
		HopStats: HopStats{
			ElapsedMin: -1,
			ElapsedMax: -1,
			JitterMin:  -1,
			JitterMax:  -1,
		},
		attemptsleft: 3}
}

func (h *HopHandler) run(seq int) error {
	h.HopStats.PingTotal += 1
	writer, err := NewUdpWriter(seq, h.connbehavior.ttl, h.remote)
	if err != nil {
		return err
	}
	h.start = time.Now()
	h.udpwriter = writer
	writer.poke()
	writer.cleanup()
	return nil
}

func (h *HopHandler) MemoryLatency(latency int64) {
	h.HopStats.HistoryWriterHead += 1
	if h.HopStats.HistoryWriterHead == len(h.HopStats.History) {
		h.HopStats.HistoryWriterHead = 0
	}
	h.HopStats.History[h.HopStats.HistoryWriterHead] = latency

	/* debug
	if h.connbehavior.ttl == 4 {
		fmt.Println(h.HopStats.History)
	}
	*/
}

func (h *HopHandler) GetJitter() int64 {
	acc := int64(0)
	h.HopStats.HistoryReaderHead += 1
	if h.HopStats.HistoryReaderHead == len(h.HopStats.History) {
		h.HopStats.HistoryReaderHead = 0
	}
	prev := int64(-1)
	idx := h.HopStats.HistoryReaderHead
	for i := 0; i < 5; i++ {
		idx += 1
		if idx == len(h.HopStats.History) {
			idx = 0
		}
		if prev != -1 {
			jitter := h.HopStats.History[idx] - prev
			acc += (jitter ^ (jitter >> 63)) - (jitter >> 63) // abs, 64th bit is sign
		}
		prev = h.HopStats.History[idx]
	}
	return acc / 4
}
