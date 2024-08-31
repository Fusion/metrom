package net

import (
	"errors"
	"net"
	"os"
	"time"
)

type IcmpHandler struct {
	icmpListener IcmpListener
	timeout      time.Duration
}

func NewIcmpHandler(icmpListener IcmpListener, timeout time.Duration) IcmpHandler {
	return IcmpHandler{
		icmpListener: icmpListener,
		timeout:      timeout}
}

func (i IcmpHandler) cleanup() {
	i.icmpListener.packetConn.Close()
}

func (i IcmpHandler) listen() (*IcmpAnswer, error) {
	/*
		if err := i.icmpListener.packetConn.SetDeadline(time.Now().Add(i.timeout)); err != nil {
			return nil, err
		}
	*/
	readBytes := make([]byte, 1500)
	_, sAddr, connErr := i.icmpListener.packetConn.ReadFrom(readBytes)
	originPort := int(readBytes[28])*256 + int(readBytes[29]) // lol ntohs says hello

	if connErr != nil {
		if errors.Is(connErr, os.ErrDeadlineExceeded) {
			return nil, &TimeoutError{}
		}
		// TODO
	}
	remoteIp := net.ParseIP(sAddr.String())
	if remoteIp == nil {
		return nil, &TimeoutError{}
	}

	return &IcmpAnswer{
		originPort: originPort,
		ip:         remoteIp,
		name:       []string{}}, nil
}
