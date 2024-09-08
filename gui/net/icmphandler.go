package net

import (
	"errors"
	"metrom/util"
	"net"
	"os"
	"sync"
	"time"
)

type IcmpHandler struct {
	icmpListener  IcmpListener
	timeout       time.Duration
	Mailbox       chan *IcmpAnswer
	cancelRequest chan bool
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

func (i *IcmpHandler) listen(group *sync.WaitGroup) {
	defer group.Done()

	i.cancelRequest = make(chan bool)
	util.Logger.Log("icmphandler:listen")
	for {
		select {
		case <-i.cancelRequest:
			i.icmpListener.packetConn.Close()
			util.Logger.Log("icmphandler:listen:cancel:receive")
			return
		default:
			readBytes := make([]byte, 1500)
			i.icmpListener.packetConn.SetReadDeadline(time.Now().Add(3 * time.Second))
			_, sAddr, connErr := i.icmpListener.packetConn.ReadFrom(readBytes)
			// Origina Port: offset from ICMP header included in IP header
			originPort := int(readBytes[28])*256 + int(readBytes[29]) // lol ntohs says hello
			candidate := false
			if readBytes[0] == 0x03 { // Found target maybe -- it will be dst unreachable, not TTL
				candidate = true
			}

			if connErr != nil {
				if errors.Is(connErr, os.ErrDeadlineExceeded) {
					continue
				}
			}
			remoteIp := net.ParseIP(sAddr.String())
			if remoteIp == nil {
				continue // TODO
			}

			i.Mailbox <- &IcmpAnswer{
				originPort: originPort,
				ip:         remoteIp,
				name:       []string{},
				candidate:  candidate}
		}
	}
}
func (i *IcmpHandler) cancel() {
	util.Logger.Log("icmphandler:listen:cancel:send")
	i.cancelRequest <- true
}
