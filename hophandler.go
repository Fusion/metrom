package main

import (
	"time"
)

type HopHandler struct {
	local          Local
	remote         Remote
	connbehavior   ConnBehavior
	firstroundtrip bool
	udpwriter      *UdpWriter
	start          time.Time
}

func NewHopHandler(local Local, remote Remote, connbehavior ConnBehavior) HopHandler {
	return HopHandler{
		local:          local,
		remote:         remote,
		connbehavior:   connbehavior,
		firstroundtrip: true}
}

func (h *HopHandler) run() error {
	writer, err := NewUdpWriter(h.connbehavior.ttl, h.remote)
	if err != nil {
		return err
	}
	h.start = time.Now()
	h.udpwriter = writer
	writer.poke()
	writer.cleanup()
	return nil
	/*
		nextPause, err := h.roundtrip()
		if err != nil {
			_, ok := err.(*TimeoutError)
			if ok {
				fmt.Println("Uuuuh timeout keep retrying perhaps?")
			}
			_, ok = err.(*FinalHopError)
			if ok {
				finalHop = true
			}
			fmt.Println(err)
		}
	*/
}

/*
func (h *HopHandler) roundtrip() (time.Duration, error) {
	writer, err := NewUdpWriter(h.remote)
	if err != nil {
		return 0, err
	}

	listener, err := NewIcmpListener(h.local)
	if err != nil {
		return 0, err
	}
	reader := NewIcmpHandler(*listener, h.connbehavior.timeout)

	start := time.Now()
	writer.poke(1, h.connbehavior.ttl)
	answer, err := reader.listen()
	if err != nil {
		return 0, err
	}
	latency := time.Since(start)
	reader.cleanup()
	writer.cleanup()

	nextPause := h.connbehavior.pause - latency
	if nextPause < 0 {
		nextPause = 0
	}
	fmt.Printf("%d: from %s (%s) in %dms\n", h.connbehavior.ttl, answer.ip.String(), answer.name, latency.Milliseconds())
	if h.firstroundtrip && answer.ip.Equal(h.remote.ip) {
		return nextPause, &FinalHopError{}
	}

	return nextPause, nil
}
*/
