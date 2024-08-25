package main

import (
	"fmt"
	"net"
)

type Local struct {
	address string
}

func NewLocal() (*Local, error) {
	return &Local{address: "0.0.0.0"}, nil
}

func (l Local) enumerate() {
	fmt.Println("Finding mahhhh socks")
	socketAddr := [4]byte{0, 0, 0, 0}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if len(ipnet.IP.To4()) == net.IPv4len {
				copy(socketAddr[:], ipnet.IP.To4())
				fmt.Printf("if: %s\n", ipnet.String())
			}
		}
	}
}
