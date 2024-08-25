package main

import (
	"fmt"
	"time"
)

type HopHandler struct {
	local          Local
	remote         Remote
	connbehavior   ConnBehavior
	firstroundtrip bool
	start          time.Time
}

func NewHopHandler(local Local, remote Remote, connbehavior ConnBehavior) HopHandler {
	return HopHandler{
		local:          local,
		remote:         remote,
		connbehavior:   connbehavior,
		firstroundtrip: true}
}

func (h HopHandler) run() {
	finalHop := false
	for {
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
		if h.firstroundtrip {
			if !finalHop {
				nextHopHandler := NewHopHandler(h.local, h.remote, ConnBehavior{
					pause:   h.connbehavior.timeout,
					timeout: h.connbehavior.timeout,
					retries: h.connbehavior.retries,
					ttl:     h.connbehavior.ttl + 1})
				go nextHopHandler.run()
			}
			h.firstroundtrip = false
		}
		if nextPause > 0 {
			time.Sleep(nextPause)
		}
	}
}

func (h HopHandler) roundtrip() (time.Duration, error) {
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
	writer.poke(h.connbehavior.ttl)
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
