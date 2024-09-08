package net

import (
	"fmt"
	"metrom/util"
	"net"
)

type Local struct {
	address string
}

func NewLocal() (*Local, error) {
	// Determine preferred IP address
	// TODO Make it selectable, too
	var address string
	conn, err := net.Dial("udp", "1.1.1.1:80")
	if err != nil {
		// TODO make a note -- you may have bigger problems!
		address = "0.0.0.0"
	} else {
		defer conn.Close()
		address = conn.LocalAddr().(*net.UDPAddr).IP.String()
	}
	util.Logger.Log(fmt.Sprintf("Using local IP [%s]", address))
	return &Local{address: address}, nil
}

func (l Local) enumerate() {
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
