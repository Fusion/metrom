package main

import (

	//"golang.org/x/net/icmp"

	"fmt"
	"time"
)

/*
 * TODO
 * figure listener address, currently hardcoded in local to 0.0.0.0
 */
func main() {
	const port = 33434 // Common starting port used by traceroute tools
	pause := time.Duration(float64(3) * float64(time.Second.Nanoseconds()))
	timeout := time.Duration(float64(3) * float64(time.Second.Nanoseconds()))
	const maxhops = 30

	remote, err := NewRemote("red.voilaweb.com", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	local, err := NewLocal()
	if err != nil {
		fmt.Println(err)
		return
	}

	for hop := 1; hop <= maxhops; hop++ {
		hopHandler := NewHopHandler(*local, *remote, ConnBehavior{
			pause:   pause,
			timeout: timeout,
			retries: 0,
			ttl:     hop})
		hopHandler.run()
	}
}
