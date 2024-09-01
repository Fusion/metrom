package net

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"
)

type IcmpHandler struct {
	icmpListener  IcmpListener
	timeout       time.Duration
	Mailbox       chan *IcmpAnswer
	cancelRequest chan struct{}
}

func NewIcmpHandler(icmpListener IcmpListener, timeout time.Duration) IcmpHandler {
	return IcmpHandler{
		icmpListener: icmpListener,
		timeout:      timeout,
		Mailbox:      make(chan *IcmpAnswer, 1000 /* buffer size*/)}
}

func (i *IcmpHandler) cleanup() {
	i.icmpListener.packetConn.Close()
}

func (i *IcmpHandler) listen() {
	i.cancelRequest = make(chan struct{})
	for {
		select {
		case <-i.cancelRequest:
			fmt.Println("icmphandler canceling")
			i.icmpListener.packetConn.Close()
			return
		default:
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
}
func (i *IcmpHandler) cancel() {
	close(i.cancelRequest)
}
