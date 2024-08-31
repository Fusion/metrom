package net

import (
	"time"
)

type HopHandler struct {
	local        Local
	remote       Remote
	connbehavior ConnBehavior
	attemptsleft int
	udpwriter    *UdpWriter
	start        time.Time
	linger       chan bool
}

func NewHopHandler(local Local, remote Remote, connbehavior ConnBehavior) HopHandler {
	return HopHandler{
		local:        local,
		remote:       remote,
		connbehavior: connbehavior,
		attemptsleft: 3}
}

// TODO Setup a timeout retry guy
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
}
