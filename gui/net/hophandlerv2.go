package net

import (
	"strconv"
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
	local           Local
	remote          Remote
	handlerMap      *sync.Map
	connbehavior    ConnBehavior
	attemptsleft    int
	udpwriter       *UdpWriter
	start           time.Time
	HopStats        HopStatsV2
	hhCancelRequest chan bool
	pokeRequest     chan bool
	poked           bool
	dead            bool
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
		attemptsleft:    3,
		hhCancelRequest: make(chan bool),
		pokeRequest:     make(chan bool)}
}

func (h *HopHandlerV2) Run(group *sync.WaitGroup) {
	defer group.Done()

	h.dead = false
	h.poke()

	for {
		timer := time.NewTimer(h.connbehavior.timeout)
		select {
		case <-h.hhCancelRequest:
			h.dead = true
			return
		case <-timer.C:
			if h.dead {
				return
			}
			//fmt.Println("DING! ttl = ", h.connbehavior.ttl, " id:", h.hhCancelRequest)
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
	var err error
	if h.udpwriter == nil {
		h.udpwriter, err = NewUdpWriter(0, h.connbehavior.ttl, h.remote)
	}
	h.handlerMap.Store(strconv.Itoa(h.udpwriter.sourcePort), h) // store handler indexed by just booked source port
	if err != nil {
		return err
	}
	h.start = time.Now()
	h.udpwriter.poke()
	//h.udpwriter.cleanup()
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
