package main

import (
	"metrom/components"
	"metrom/models"
	"metrom/net"
	"metrom/util"
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
		resolver:   resolver}
}

func (h *HopsCollection) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", "no-store")
	hops := []models.FrontendHop{} // If I knew size in advance I couldl use make() to right-size this guy

	if error := h.controller.GetState("error"); error != "" {
		openModal(w, r, "Error", error)
		h.controller.SetState("error", "")
		return
	}
	util.Logger.Log("hops / refresh")
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
		// TODO configurable
		if data.HopStatus[i].Candidate || data.HopStatus[i].RemoteIp == data.RemoteIp { // Found target?
			break
		}
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
	if r.FormValue("cb-resolve") == "on" {
		util.Logger.Log("hops /toggleresolve set resolve to on")
		h.controller.SetSetting("resolve", "on")
		models.SetPreference(AppPreferences, "resolve", true)
	} else {
		util.Logger.Log("hops /toggleresolve set resolve to off")
		h.controller.SetSetting("resolve", "off")
		models.SetPreference(AppPreferences, "resolve", false)
	}
	models.SavePreferences(AppPreferences)
}

func (h *HopsCollection) SaveMaxHops(w http.ResponseWriter, r *http.Request) {
	util.Logger.Log("hops /savemaxhops save maxhops preference")
	value, _ := strconv.Atoi(r.FormValue("maxhopsslider"))
	models.SetPreference(AppPreferences, "maxhops", value)
	models.SavePreferences(AppPreferences)
}

func (h *HopsCollection) SaveTimeout(w http.ResponseWriter, r *http.Request) {
	util.Logger.Log("hops /savetimeout save timeout preference")
	value, _ := strconv.Atoi(r.FormValue("timeoutslider"))
	models.SetPreference(AppPreferences, "timeout", value)
	models.SavePreferences(AppPreferences)
}

func (h *HopsCollection) SaveProbes(w http.ResponseWriter, r *http.Request) {
	util.Logger.Log("hops /saveprobes save probecount preference")
	value, _ := strconv.Atoi(r.FormValue("probesslider"))
	models.SetPreference(AppPreferences, "probecount", value)
	models.SavePreferences(AppPreferences)
}

func (h *HopsCollection) SaveJitter(w http.ResponseWriter, r *http.Request) {
	util.Logger.Log("hops /savejitter save jitttersamples preference")
	value, _ := strconv.Atoi(r.FormValue("jitterslider"))
	models.SetPreference(AppPreferences, "jittersamples", value)
	models.SavePreferences(AppPreferences)
}

func (h *HopsCollection) ResetSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", "no-store")
	w.Header().Set("HX-Retarget", "#actioncontainer")
	if h.controller.IsBusy() {
		util.Logger.Log("hops /resetsearch check busy state")
		component := components.OOBButton("busy")
		component.Render(r.Context(), w)
	} else {
		util.Logger.Log("hops /resetsearch set ready state")
		component := components.OOBButton("")
		component.Render(r.Context(), w)
	}
}

func (h *HopsCollection) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", "no-store")

	if h.controller.Cancel() { // I was running!
		util.Logger.Log("hops / (post) search cancelation")
		w.Header().Set("HX-Retarget", "#actioncontainer")
		component := components.OOBButton("busy")
		component.Render(r.Context(), w)
		return
	}

	util.Logger.Log("hops / (post) starting search")
	subject := r.FormValue("subject")
	if subject == "" {
		openModal(w, r, "Hold on", "Please enter a host name or ip address")
		return
	}

	h.resolver.Cleanup()
	go func() {
		h.controller.Run(
			net.WithOption(net.VHost{Value: r.FormValue("subject")}),
			net.WithOption(net.VMaxHops{Value: 30}),
		)
	}()

	util.Logger.Log("hops / (post) set search state")
	w.Header().Set("HX-Retarget", "#actioncontainer")
	component := components.OOBButton("search")
	component.Render(r.Context(), w)
}
