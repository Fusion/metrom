package net

import (
	"fmt"
	"net"
	"strconv"
)

type Remote struct {
	address string
	port    uint16
	ip      net.IP
}

func NewRemote(address string, port uint16) (*Remote, error) {
	ip := net.ParseIP(address)
	if ip == nil {
		ips, err := net.LookupIP(address)
		if err != nil {
			return nil, err
		}
		if len(ips) == 0 {
			return nil, fmt.Errorf("destination lookup failed")
		}
		ip = ips[0]
	}
	return &Remote{
		address: address,
		port:    port,
		ip:      ip}, nil
}

func (r Remote) String() string {
	return r.ip.String() + ":" + strconv.Itoa(int(r.port))
}

func (r Remote) StringWithSeq(seq int) string {
	return r.ip.String() + ":" + strconv.Itoa(int(r.port)+seq)
}
