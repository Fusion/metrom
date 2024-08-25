package main

import "net"

type IcmpAnswer struct {
	originPort int
	ip         net.IP
	name       []string
}
