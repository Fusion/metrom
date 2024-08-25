package main

import (
	"net"

	"golang.org/x/net/ipv4"
)

type UdpWriter struct {
	conn net.Conn
}

func NewUdpWriter(remote Remote) (*UdpWriter, error) {

	conn, err := net.Dial("udp4", remote.String())
	if err != nil {
		return nil, err
	}

	return &UdpWriter{conn: conn}, nil
}

func (r UdpWriter) cleanup() {
	r.conn.Close()
}

func (r UdpWriter) poke(ttl int) error {
	newConn := ipv4.NewConn(r.conn)
	if err := newConn.SetTTL(ttl); err != nil {
		return err
	}
	_, err := r.conn.Write([]byte("TABS"))
	if err != nil {
		return err
	}
	return nil
}
