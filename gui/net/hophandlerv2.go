package net

import (
	"sync"
	"time"
)

type HopStatsV2 struct {
	ElapsedMin        int64
	ElapsedMax        int64
	History           [5]int64
	HistoryWriterHead int
	HistoryReaderHead int
	JitterMin         int64
	JitterMax         int64
	PingTotal         int64
	PingMiss          int64
	ProbeMiss         int
}
type HopHandlerV2 struct {
	local         Local
	remote        Remote
	handlerMap    *sync.Map
	connbehavior  ConnBehavior
	attemptsleft  int
	udpwriter     *UdpWriter
	start         time.Time
	HopStats      HopStatsV2
	cancelRequest chan struct{}
	pokeRequest   chan bool
	poked         bool
}

func NewHopHandlerV2(local Local, remote Remote, handlerMap *sync.Map, connbehavior ConnBehavior) HopHandlerV2 {
	return HopHandlerV2{
		local:        local,
		remote:       remote,
		handlerMap:   handlerMap,
		connbehavior: connbehavior,
		HopStats: HopStatsV2{
			ElapsedMin: -1,
			ElapsedMax: -1,
			JitterMin:  -1,
			JitterMax:  -1,
		},
		attemptsleft:  3,
		cancelRequest: make(chan struct{}),
		pokeRequest:   make(chan bool)}
}

// TODO add 3 seconds timeout after which if not reset by a pokeQuest we try again
func (h *HopHandlerV2) Run() {
	h.poke()
	for {
		timer := time.NewTimer(h.connbehavior.timeout)
		select {
		case <-h.cancelRequest:
			return
		case <-timer.C:
			if h.HopStats.PingTotal%3 == 0 { // Initial ping in a sequence of n probes
				h.HopStats.ProbeMiss = 0
			}
			if !h.poked {
				h.HopStats.ProbeMiss += 1
			} else {
				h.poked = false
			}
			if h.HopStats.PingTotal%3 == 2 { // Final ping in sequence
				if h.HopStats.ProbeMiss == 3 {
					h.HopStats.PingMiss += 1
				}
			}
			h.HopStats.PingTotal += 1
			h.poke()
		case <-h.pokeRequest:
			h.poked = true
		}
	}
}

func (h *HopHandlerV2) poke() error {
	writer, err := NewUdpWriter(0, h.connbehavior.ttl, h.remote)
	h.handlerMap.Store(string(writer.sourcePort), h) // store handler indexed by just booked source port
	if err != nil {
		return err
	}
	h.start = time.Now()
	h.udpwriter = writer
	writer.poke()
	writer.cleanup()
	return nil
}

func (h *HopHandlerV2) MemoryLatency(latency int64) {
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

func (h *HopHandlerV2) GetJitter() int64 {
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
