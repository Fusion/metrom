package main

import (
	"gui/components"
	"gui/models"
	"gui/net"
	"net/http"
	"strconv"
)

type HopsCollection struct {
	controller BackendController[net.Data, net.NetOption]
	resolver   net.Resolver
}

func NewHopsCollection(controller BackendController[net.Data, net.NetOption], resolver net.Resolver) HopsCollection {
	return HopsCollection{
		controller: controller,
		resolver:   net.NewResolver()}
}

func (h *HopsCollection) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", "no-store")
	hops := []models.FrontendHop{} // If I knew size in advance I couldl use make() to right-size this guy

	h.controller.LockData()
	data := h.controller.GetData()
	for i := 1; i < len(data.HopStatus); i++ {
		if data.HopStatus[i].PingTotal == 0 {
			// TODO
			// Decide what the best UX would be: display all these "***"
			// or simply not display them, and have a less bloated experience?
			continue
		}
		var host string
		if h.controller.GetSetting("resolve") == "on" {
			if len(data.HopStatus[i].RemoteDNS) == 0 {
				data.HopStatus[i].RemoteDNS = h.resolver.ResolveAddress(data.HopStatus[i].RemoteIp)
			}
			host = data.HopStatus[i].RemoteDNS[0]
		} else {
			host = data.HopStatus[i].RemoteIp
		}
		hops = append(hops, models.FrontendHop{
			Hop:        strconv.Itoa(i),
			Host:       host,
			Loss:       strconv.FormatInt(data.HopStatus[i].PingMiss*100/data.HopStatus[i].PingTotal, 10),
			LatencyAvg: strconv.FormatInt(data.HopStatus[i].Elapsed, 10),
			LatencyMin: strconv.FormatInt(data.HopStatus[i].ElapsedMin, 10),
			LatencyMax: strconv.FormatInt(data.HopStatus[i].ElapsedMax, 10),
			JitterAvg:  strconv.FormatInt(data.HopStatus[i].Jitter, 10),
			JitterMin:  strconv.FormatInt(data.HopStatus[i].JitterMin, 10),
			JitterMax:  strconv.FormatInt(data.HopStatus[i].JitterMax, 10),
		})
	}
	h.controller.UnlockData()
	component := components.HopTable(hops)
	/*
		h.hops = append(h.hops, models.FrontendHop{
			Hop:     "1",
			Host:    "192.168.1.5",
			Loss:    "0",
			Latency: "0",
			Jitter:  "0",
		})
		component := components.HopTable(h.hops)
	*/
	component.Render(r.Context(), w)
}

func (h *HopsCollection) ToggleResolve(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", "no-store")
	if r.FormValue("cb-resolve") == "on" {
		h.controller.SetSetting("resolve", "on")
	} else {
		h.controller.SetSetting("resolve", "off")
	}
}

func (h *HopsCollection) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", "no-store")
	subject := r.FormValue("subject")
	if subject == "" {
		openModal(w, r, "Hold on", "Please enter a host name or ip address")
		return
	}
	err := h.controller.Run(
		net.WithOption(net.VHost{Value: r.FormValue("subject")}),
		net.WithOption(net.VMaxHops{Value: 30}),
	)
	if err != nil {
		openModal(w, r, "Error", err.Error())
		return
	}
}
