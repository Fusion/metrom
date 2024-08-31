package net

import (

	//"golang.org/x/net/icmp"

	"fmt"
	"sync"
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

type Value interface {
	_isValue()
}
type VHost struct {
	Value string
}
type VMaxHops struct {
	Value int
}

func (v VHost) _isValue()    {}
func (v VMaxHops) _isValue() {}

type Options struct {
	Host    string
	MaxHops int
}

type NetOption func(*Options)

func WithOption(value Value) NetOption {
	switch value.(type) {
	case VHost:
		return func(h *Options) {
			h.Host = value.(VHost).Value
		}
	case VMaxHops:
		return func(h *Options) {
			h.MaxHops = value.(VMaxHops).Value
		}
	default:
		return func(h *Options) {} // TODO Warning
	}
}

type Data struct {
	TopHop    int
	HopStatus []HopStatus
}

type NetController struct {
	HopsLock     sync.Mutex
	data         Data
	SettingsLock sync.Mutex
	settings     map[string]string
	hopHandlers  map[string]*HopHandler
}

func NewNetController() NetController {
	return NetController{
		settings: make(map[string]string),
	}
}

func (n *NetController) LockData() {
	n.HopsLock.Lock()
}

func (n *NetController) UnlockData() {
	n.HopsLock.Unlock()
}

func (n *NetController) Run(options ...NetOption) error {
	runOptions := &Options{}
	for _, opt := range options {
		opt(runOptions)
	}

	const port = 33434 // Common starting port used by traceroute tools

	remote, err := NewRemote(runOptions.Host, port)
	if err != nil {
		return err
	}

	local, err := NewLocal()
	if err != nil {
		return err
	}

	listener, err := NewIcmpListener(*local)
	if err != nil {
		return err
	}

	pause := time.Duration(float64(3) * float64(time.Second.Nanoseconds()))
	timeout := time.Duration(float64(3) * float64(time.Second.Nanoseconds()))
	quiescent := time.Duration(float64(5) * float64(time.Second.Nanoseconds()))

	n.data.HopStatus = make([]HopStatus, runOptions.MaxHops+1)
	n.hopHandlers = make(map[string]*HopHandler)

	reader := NewIcmpHandler(*listener, timeout)
	for hop := 1; hop <= runOptions.MaxHops; hop++ {
		hopHandler := NewHopHandler(*local, *remote, ConnBehavior{
			pause:     pause,
			timeout:   timeout,
			quiescent: quiescent,
			retries:   0,
			ttl:       hop})
		n.runHandler(&hopHandler, 0 /* delay */)
	}

	foundTarget := false

	for {
		//fmt.Println("about to listen")
		answer, err := reader.listen()
		//fmt.Println("done listening")
		if err != nil {
			_, ok := err.(*TimeoutError)
			if ok {
				// TODO timeout
				continue
			}
			return err
		}
		hopHandler, ok := n.hopHandlers[string(answer.originPort)]
		if !ok {
			continue // TODO stray packet?
		}
		hopHandler.linger <- true // cancel timeout
		delete(n.hopHandlers, string(answer.originPort))
		if foundTarget {
			// Are we past our target?
			if hopHandler.connbehavior.ttl > n.data.TopHop {
				continue
			}
		}
		elapsed := time.Since(hopHandler.start).Milliseconds()
		if hopHandler.HopStats.ElapsedMin == -1 || elapsed < hopHandler.HopStats.ElapsedMin {
			hopHandler.HopStats.ElapsedMin = elapsed
		}
		if hopHandler.HopStats.ElapsedMax == -1 || elapsed > hopHandler.HopStats.ElapsedMax {
			hopHandler.HopStats.ElapsedMax = elapsed
		}
		jitter, jitterMin, jitterMax := int64(0), int64(0), int64(0)
		if hopHandler.HopStats.PingTotal >= 5 {
			jitter = hopHandler.GetJitter()
			if hopHandler.HopStats.JitterMin == -1 || jitter < hopHandler.HopStats.JitterMin {
				hopHandler.HopStats.JitterMin = jitter
			}
			jitterMin = hopHandler.HopStats.JitterMin
			if hopHandler.HopStats.JitterMax == -1 || jitter > hopHandler.HopStats.JitterMax {
				hopHandler.HopStats.JitterMax = jitter
			}
			jitterMax = hopHandler.HopStats.JitterMax
		}
		hopHandler.MemoryLatency(elapsed)
		/*
			fmt.Printf("%d: from %s (%s) reply to %d in %dms\n",
				hopHandler.connbehavior.ttl,
				answer.ip.String(),
				answer.name,
				answer.originPort,
				elapsed)
		*/
		n.HopsLock.Lock()
		n.data.HopStatus[hopHandler.connbehavior.ttl] = HopStatus{
			RemoteIp:   answer.ip.String(),
			RemoteDNS:  answer.name,
			Elapsed:    elapsed,
			ElapsedMin: hopHandler.HopStats.ElapsedMin,
			ElapsedMax: hopHandler.HopStats.ElapsedMax,
			PingTotal:  hopHandler.HopStats.PingTotal + 1,
			PingMiss:   hopHandler.HopStats.PingMiss,
			Jitter:     jitter,
			JitterMin:  jitterMin,
			JitterMax:  jitterMax,
		}
		n.HopsLock.Unlock()

		if answer.ip.String() == remote.ip.String() {
			newTop := false
			if foundTarget {
				// Always stop at the closest answer
				if hopHandler.connbehavior.ttl < n.data.TopHop {
					newTop = true
				}
			} else {
				foundTarget = true
				newTop = true
			}
			if newTop {
				n.data.TopHop = hopHandler.connbehavior.ttl
				if n.data.TopHop <= runOptions.MaxHops {
					for i := n.data.TopHop; i <= runOptions.MaxHops; i++ {
						n.data.HopStatus[i] = NewHopStatus()
					}
				}
			}
		} else {
			if foundTarget {
				if hopHandler.connbehavior.ttl > n.data.TopHop {
					continue
				}
			}
		}

		//updateDisplay(n.data.HopStatus)
		n.runHandler(hopHandler, hopHandler.connbehavior.pause-time.Duration(elapsed) /* delay */)
	}
}

func (n *NetController) runHandler(hopHandler *HopHandler, delay time.Duration) {
	go func() {
		if delay > 0 {
			time.Sleep(delay)
		}

		hopHandler.run()
		n.HopsLock.Lock()
		n.hopHandlers[string(hopHandler.udpwriter.sourcePort)] = hopHandler
		n.HopsLock.Unlock()

		hopHandler.linger = make(chan bool)
		timer := time.NewTimer(hopHandler.connbehavior.timeout)
		go func() {
			select {
			case <-timer.C:
				hopHandler.HopStats.PingMiss += 1
				n.runHandler(hopHandler, hopHandler.connbehavior.quiescent-hopHandler.connbehavior.timeout /* delay */)
			case <-hopHandler.linger:
			}
		}()
	}()
}

func (n *NetController) GetData() Data {
	return n.data
}

func (n *NetController) SetSetting(key string, value string) {
	n.SettingsLock.Lock()
	defer n.SettingsLock.Unlock()
	n.settings[key] = value
}

func (n *NetController) GetSetting(key string) string {
	n.SettingsLock.Lock()
	defer n.SettingsLock.Unlock()
	res := n.settings[key]
	return res
}

func updateDisplay(hopStatus []HopStatus) {
	for i := 1; i < len(hopStatus); i++ {
		fmt.Printf("%d: %s\n", i, hopStatus[i])
	}
}
