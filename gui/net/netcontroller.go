package net

//"golang.org/x/net/icmp"

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
type ValueV1 interface {
	_isValue()
}
type VHostV1 struct {
	Value string
}
type VMaxHopsV1 struct {
	Value int
}

func (v VHostV1) _isValue()    {}
func (v VMaxHopsV1) _isValue() {}

type OptionsV1 struct {
	Host    string
	MaxHops int
}

type NetOptionV1 func(*Options)

func WithOptionV1(value Value) NetOption {
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

// TODO CHECK IF ALREADY REPLIED WITH A LOWER TTL!!!
type DataV1 struct {
	TopHop    int
	HopStatus []HopStatus
}

type NetController struct {
	HopsLock      sync.Mutex
	seq           int
	data          DataV1
	settings      sync.Map
	runOptions    *Options
	hopHandlers   sync.Map
	running       bool
	cancelRequest chan struct{}
}

func NewNetController() NetController {
	return NetController{}
}

func (n *NetController) LockData() {
	n.HopsLock.Lock()
}

func (n *NetController) UnlockData() {
	n.HopsLock.Unlock()
}

func (n *NetController) Run(options ...NetOption) error {
	n.runOptions = &Options{}
	for _, opt := range options {
		opt(n.runOptions)
	}

	const port = 33434 // Common starting port used by traceroute tools

	remote, err := NewRemote(n.runOptions.Host, port)
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

	n.running = true

	pause := time.Duration(float64(3) * float64(time.Second.Nanoseconds()))
	timeout := time.Duration(float64(3) * float64(time.Second.Nanoseconds()))
	quiescent := time.Duration(float64(5) * float64(time.Second.Nanoseconds()))

	// Cleanup cache from any previous run
	n.Reset()

	reader := NewIcmpHandler(*listener, timeout)
	for hop := 1; hop <= n.runOptions.MaxHops; hop++ {
		hopHandler := NewHopHandler(*local, *remote, ConnBehavior{
			pause:     pause,
			timeout:   timeout,
			quiescent: quiescent,
			retries:   0,
			ttl:       hop})
		n.runHandler(&hopHandler, 0)
	}

	foundTarget := false

	n.cancelRequest = make(chan struct{})
	go reader.listen()

	for answer := range reader.Mailbox {
		select {
		case <-n.cancelRequest:
			reader.cancel()
			fmt.Println("netcontroller canceling")
			n.running = false
			n.Reset()
			return nil
		default:
			// TODO CONCURRENT MAP R/W ERROR
			hopHandlerVal, ok := n.hopHandlers.Load(string(answer.originPort))
			if !ok {
				continue // TODO stray packet?
			}
			hopHandler := hopHandlerVal.(*HopHandler)
			hopHandler.donotlinger <- true
			n.hopHandlers.Delete(string(answer.originPort))
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
			n.LockData()
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
			n.UnlockData()

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
					n.LockData()
					n.data.TopHop = hopHandler.connbehavior.ttl
					if n.data.TopHop <= n.runOptions.MaxHops {
						for i := n.data.TopHop; i <= n.runOptions.MaxHops; i++ {
							n.data.HopStatus[i] = NewHopStatus()
						}
					}
					n.UnlockData()
				}
			} else {
				if foundTarget {
					if hopHandler.connbehavior.ttl > n.data.TopHop {
						continue
					}
				}
			}

			//updateDisplay(n.data.HopStatus)
			n.runHandler(hopHandler, hopHandler.connbehavior.pause-time.Duration(elapsed))
		}
	}

	return nil
}

func (n *NetController) Reset() {
	n.LockData()
	n.data.TopHop = 0
	n.hopHandlers.Range(func(key interface{}, value interface{}) bool {
		n.hopHandlers.Delete(key)
		return true
	})
	n.data.HopStatus = make([]HopStatus, n.runOptions.MaxHops+1)
	n.UnlockData()
}

// TODO return a 3-state enum
func (n *NetController) Cancel() bool {
	if n.running {
		close(n.cancelRequest)
		return true
	}
	return false
}

func (n *NetController) runHandler(hopHandler *HopHandler, delay time.Duration) {
	go func() {
		if delay > 0 {
			time.Sleep(delay)
		}
		n.LockData()
		n.seq = 0
		n.UnlockData()
		hopHandler.run(n.seq)
		n.hopHandlers.Store(string(hopHandler.udpwriter.sourcePort), hopHandler)

		hopHandler.donotlinger = make(chan bool)
		timer := time.NewTimer(hopHandler.connbehavior.timeout)
		go func() {
			select {
			case <-timer.C:
				hopHandler.HopStats.PingMiss += 1
				n.runHandler(hopHandler, hopHandler.connbehavior.quiescent-hopHandler.connbehavior.timeout)
			case <-hopHandler.donotlinger:
				return
			}
		}()
	}()
}

func (n *NetController) GetData() DataV1 {
	return n.data
}

func (n *NetController) SetSetting(key string, value string) {
	n.settings.Store(key, value)
}

func (n *NetController) GetSetting(key string) string {
	res, _ := n.settings.Load(key)
	return res.(string)
}

func updateDisplay(hopStatus []HopStatus) {
	for i := 1; i < len(hopStatus); i++ {
		fmt.Printf("%d: %s\n", i, hopStatus[i])
	}
}
*/
