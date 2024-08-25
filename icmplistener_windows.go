//go:build windows
// +build windows

package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"syscall"
	"unsafe"
)

type IcmpListener struct {
	packetConn net.PacketConn
}

func NewIcmpListener(local Local) (*IcmpListener, error) {
	dialedConn, err := net.Dial("ip4:icmp", local.address)
	if err != nil {
		return nil, err
	}
	localAddr := dialedConn.LocalAddr()
	dialedConn.Close()
	var socketHandle syscall.Handle
	cfg := net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			return c.Control(func(s uintptr) {
				socketHandle = syscall.Handle(s)
			})
		},
	}
	packetConn, err := cfg.ListenPacket(context.Background(), "ip4:icmp", localAddr.String())
	if err != nil {
		return nil, err
	}
	unused := uint32(0) // Documentation states that this is unused, but WSAIoctl fails without it.
	flag := uint32(3)   // IPLEVEL
	size := uint32(unsafe.Sizeof(flag))
	err = syscall.WSAIoctl(socketHandle, syscall.IOC_IN|syscall.IOC_VENDOR|1, (*byte)(unsafe.Pointer(&flag)), size, nil, 0, &unused, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to set socket to listen to all packests: %s\n", os.NewSyscallError("WSAIoctl", err))
	}

	return &IcmpListener{packetConn: packetConn}, nil
}
