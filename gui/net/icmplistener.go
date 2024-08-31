//go:build !windows
// +build !windows

package net

import (
	"net"

	"golang.org/x/net/icmp"
)

type IcmpListener struct {
	packetConn net.PacketConn
}

func NewIcmpListener(local Local) (*IcmpListener, error) {
	packetConn, err := icmp.ListenPacket("ip4:icmp", local.address)
	if err != nil {
		return nil, err
	}
	return &IcmpListener{packetConn: packetConn}, nil
}
