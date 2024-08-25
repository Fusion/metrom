package main

import (

	//"golang.org/x/net/icmp"

	"fmt"
	"time"
)

/*
 * ICMP response format:
 *|type|code|checksum|unused|ipheader|
 *|0b|00|2380|00110000|45002000|
 * type (8b) = 0b when exceeded
 * code (8b) = 00 if ttl exceeded, 01 if re-assembly exceeded
 * checksum (16b)
 * unused (32b)
 * IPheader+first8 (32b)
 * Source port (16b) offset: +62

 */

/*
 * TODO
 * figure listener address, currently hardcoded in local to 0.0.0.0
 */
func main() {
	const port = 33434 // Common starting port used by traceroute tools
	pause := time.Duration(float64(3) * float64(time.Second.Nanoseconds()))
	timeout := time.Duration(float64(3) * float64(time.Second.Nanoseconds()))
	const maxhops = 30

	hopStatus := make([]HopStatus, maxhops+1)
	hopHandlers := make(map[string]HopHandler)

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

	listener, err := NewIcmpListener(*local)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := NewIcmpHandler(*listener, timeout)
	for hop := 1; hop <= maxhops; hop++ {
		hopHandler := NewHopHandler(*local, *remote, ConnBehavior{
			pause:   pause,
			timeout: timeout,
			retries: 0,
			ttl:     hop})
		hopHandler.run()
		hopHandlers[string(hopHandler.udpwriter.sourcePort)] = hopHandler
	}
	for {
		fmt.Println("about to listen")
		answer, err := reader.listen()
		fmt.Println("done listening")
		if err != nil {
			_, ok := err.(*TimeoutError)
			if ok {
				// TODO timeout
				continue
			}
			fmt.Println(err)
			return
		}
		hopHandler, ok := hopHandlers[string(answer.originPort)]
		if !ok {
			fmt.Println("-- stray packet? --")
			continue // TODO stray packet?
		}
		delete(hopHandlers, string(answer.originPort))
		elapsed := time.Since(hopHandler.start).Milliseconds()
		/*
			fmt.Printf("%d: from %s (%s) reply to %d in %dms\n",
				hopHandler.connbehavior.ttl,
				answer.ip.String(),
				answer.name,
				answer.originPort,
				elapsed)
		*/
		hopStatus[hopHandler.connbehavior.ttl] = HopStatus{
			remoteIp:  answer.ip.String(),
			remoteDNS: answer.name,
			elapsed:   elapsed,
		}
		updateDisplay(hopStatus)
		sleeper := make(chan bool)
		go func() {
			delay := hopHandler.connbehavior.pause - time.Duration(elapsed)
			if delay > 0 {
				time.Sleep(delay)
			}
			hopHandler.run()
			hopHandlers[string(hopHandler.udpwriter.sourcePort)] = hopHandler
			sleeper <- true
		}()
	}
}

func updateDisplay(hopStatus []HopStatus) {
	for i := 1; i < len(hopStatus); i++ {
		fmt.Printf("%d: %s\n", i, hopStatus[i])
	}
}
