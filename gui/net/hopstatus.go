package net

type HopStatus struct {
	RemoteIp   string
	RemoteDNS  []string
	Elapsed    int64
	ElapsedMin int64
	ElapsedMax int64
	PingTotal  int64
	PingMiss   int64
	Jitter     int64
	JitterMin  int64
	JitterMax  int64
}

func NewHopStatus() HopStatus {
	return HopStatus{}
}
