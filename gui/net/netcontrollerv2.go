package net

import (
	"metrom/util"
	"strconv"
	"sync"
	"time"
)

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
	RemoteIp  string
}

type NetControllerV2 struct {
	HopsLock                sync.Mutex
	seq                     int
	data                    Data
	states                  sync.Map
	settings                sync.Map
	runOptions              *Options
	hopHandlers             sync.Map
	running                 bool
	controllerCancelRequest chan struct{}
	listeners               sync.WaitGroup
	writers                 sync.WaitGroup
}

func NewNetControllerV2() NetControllerV2 {
	return NetControllerV2{}
}

func (n *NetControllerV2) Run(options ...NetOption) error {
	n.runOptions = &Options{}
	for _, opt := range options {
		opt(n.runOptions)
	}

	const port = 33434 // Common starting port used by traceroute tools

	// Initialize/Reset
	n.LockData()
	n.data.TopHop = 0
	n.data.HopStatus = make([]HopStatus, n.runOptions.MaxHops+1)
	n.UnlockData()
	//

	remote, err := NewRemote(n.runOptions.Host, port)
	if err != nil {
		n.SetState("error", err.Error())
		return err
	}
	n.data.RemoteIp = remote.ip.String()

	local, err := NewLocal()
	if err != nil {
		n.SetState("error", err.Error())
		return err
	}

	listener, err := NewIcmpListener(*local)
	if err != nil {
		n.SetState("error", err.Error())
		return err
	}

	pause := time.Duration(float64(3) * float64(time.Second.Nanoseconds()))
	timeout := time.Duration(float64(3) * float64(time.Second.Nanoseconds()))
	quiescent := time.Duration(float64(5) * float64(time.Second.Nanoseconds()))

	// Instantiate global icmp listener
	reader := NewIcmpHandler(*listener, timeout)
	n.listeners.Add(1)
	go reader.listen(&n.listeners)

	// Instantiate one udp writer per hop
	for hop := 1; hop <= n.runOptions.MaxHops; hop++ {
		hopHandler := NewHopHandlerV2(
			*local,
			*remote,
			&n.hopHandlers,
			ConnBehavior{
				pause:     pause,
				timeout:   timeout,
				quiescent: quiescent,
				retries:   0,
				ttl:       hop})
		n.writers.Add(1)
		go hopHandler.Run(&n.writers)
	}

	n.controllerCancelRequest = make(chan struct{})

	n.running = true
	for answer := range reader.Mailbox {
		select {
		case <-n.controllerCancelRequest:
			n.hopHandlers.Range(func(key interface{}, value interface{}) bool {
				value.(*HopHandlerV2).hhCancelRequest <- true
				n.hopHandlers.Delete(key)
				return true
			})
			n.writers.Wait()
			util.Logger.Log("all writers terminated")
			reader.cancel()
			n.listeners.Wait()
			util.Logger.Log("all listeners terminated")
			n.running = false
			return nil
		default:
			hopHandlerVal, ok := n.hopHandlers.Load(strconv.Itoa(answer.originPort))
			if !ok {
				continue // TODO stray packet?
			}
			hopHandler := hopHandlerVal.(*HopHandlerV2)

			// Update stats
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

			n.LockData()
			n.data.HopStatus[hopHandler.connbehavior.ttl] = HopStatus{
				RemoteIp:   answer.ip.String(),
				RemoteDNS:  answer.name,
				Candidate:  answer.candidate,
				Elapsed:    elapsed,
				ElapsedMin: hopHandler.HopStats.ElapsedMin,
				ElapsedMax: hopHandler.HopStats.ElapsedMax,
				PingTotal:  hopHandler.HopStats.PingTotal + 1,
				PingMiss:   hopHandler.HopStats.PingMiss,
				Jitter:     jitter,
				JitterMin:  jitterMin,
				JitterMax:  jitterMax,
			}
			n.UnlockData()
			//
			hopHandler.pokeRequest <- true
		}
	}

	return nil
}

// Return true if it was acrtually running
func (n *NetControllerV2) Cancel() bool {
	if n.running {
		util.Logger.Log("netcontrollerv2:cancel")
		close(n.controllerCancelRequest)
		return true
	}
	return false
}

func (n *NetControllerV2) IsBusy() bool {
	return n.running
}

func (n *NetControllerV2) LockData() {
	n.HopsLock.Lock()
}

func (n *NetControllerV2) UnlockData() {
	n.HopsLock.Unlock()
}

func (n *NetControllerV2) GetData() Data {
	return n.data
}

func (n *NetControllerV2) SetSetting(key string, value string) {
	n.settings.Store(key, value)
}

func (n *NetControllerV2) GetSetting(key string) string {
	res, ok := n.settings.Load(key)
	if !ok {
		return ""
	}
	return res.(string)
}

func (n *NetControllerV2) SetState(key string, value string) {
	n.states.Store(key, value)
}

func (n *NetControllerV2) GetState(key string) string {
	res, ok := n.states.Load(key)
	if !ok {
		return ""
	}
	return res.(string)
}
