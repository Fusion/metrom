package main

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
	if err := i.icmpListener.packetConn.SetDeadline(time.Now().Add(i.timeout)); err != nil {
		return nil, err
	}
	readBytes := make([]byte, 1500)                                    // 1500 Bytes ethernet MTU
	_, sAddr, connErr := i.icmpListener.packetConn.ReadFrom(readBytes) // first return value (Code) might be useful

	if connErr != nil {
		if errors.Is(connErr, os.ErrDeadlineExceeded) {
			return nil, &TimeoutError{}
		}
	}
	var remoteIp net.IP
	remoteIp = net.ParseIP(sAddr.String())
	if remoteIp == nil {
		return nil, &TimeoutError{}
	}

	remoteDNS, _ := net.LookupAddr(remoteIp.String())
	return &IcmpAnswer{remoteIp, remoteDNS}, nil
}
