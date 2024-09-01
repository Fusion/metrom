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
	Mailbox      chan *IcmpAnswer
}

func NewIcmpHandler(icmpListener IcmpListener, timeout time.Duration) IcmpHandler {
	return IcmpHandler{
		icmpListener: icmpListener,
		timeout:      timeout,
		Mailbox:      make(chan *IcmpAnswer, 1000 /* buffer size*/)}
}

func (i IcmpHandler) cleanup() {
	i.icmpListener.packetConn.Close()
}

func (i IcmpHandler) listen() {
	for {
		readBytes := make([]byte, 1500)
		_, sAddr, connErr := i.icmpListener.packetConn.ReadFrom(readBytes)
		originPort := int(readBytes[28])*256 + int(readBytes[29]) // lol ntohs says hello

		if connErr != nil {
			if errors.Is(connErr, os.ErrDeadlineExceeded) {
				continue // TODO
			}
			// TODO
		}
		remoteIp := net.ParseIP(sAddr.String())
		if remoteIp == nil {
			continue // TODO
		}

		i.Mailbox <- &IcmpAnswer{
			originPort: originPort,
			ip:         remoteIp,
			name:       []string{}}
	}
}

func (i IcmpHandler) __listen() (*IcmpAnswer, error) {
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
