package main

import "net"

type IcmpAnswer struct {
	ip   net.IP
	name []string
}
