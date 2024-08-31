package models

type FrontendHop struct {
	Hop        string
	Host       string
	Loss       string
	LatencyAvg string
	LatencyMin string
	LatencyMax string
	JitterAvg  string
	JitterMin  string
	JitterMax  string
}
