package net

import "net"

type IcmpAnswer struct {
	originPort int
	ip         net.IP
	name       []string
	candidate  bool
}
