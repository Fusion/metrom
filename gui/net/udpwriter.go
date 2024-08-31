package net

import (
	"fmt"
	"net"

	"golang.org/x/net/ipv4"
)

type UdpWriter struct {
	conn       net.Conn
	sourcePort int
	ttl        int
}

func NewUdpWriter(ttl int, remote Remote) (*UdpWriter, error) {

	conn, err := net.Dial("udp4", remote.String())
	if err != nil {
		return nil, err
	}
	addr := conn.LocalAddr()
	udpAddr, ok := addr.(*net.UDPAddr)
	if !ok {
		return nil, fmt.Errorf("unable to get local address")
	}
	//fmt.Printf("For ttl=%d, new port = %d\n", ttl, udpAddr.Port)

	return &UdpWriter{
		conn:       conn,
		sourcePort: udpAddr.Port,
		ttl:        ttl}, nil
}

func (r UdpWriter) cleanup() {
	r.conn.Close()
}

func (r UdpWriter) poke() error {
	newConn := ipv4.NewConn(r.conn)
	if err := newConn.SetTTL(r.ttl); err != nil {
		return err
	}
	_, err := r.conn.Write([]byte("TABS"))
	if err != nil {
		return err
	}
	return nil
}
