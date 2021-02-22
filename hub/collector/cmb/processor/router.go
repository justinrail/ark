package processor

import (
	"ark/hub/collector/cmb/ws"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//Router processor for dispatch
type Router struct {
	flow.Component
	In            chan *ws.Conversation
	LoginOut      chan<- *ws.Conversation
	SendConfigOut chan<- *ws.Conversation
	SendAlarmOut  chan<- *ws.Conversation
}

//OnIn Request handler
func (router *Router) OnIn(cvs *ws.Conversation) {
	metrics.AppMetrics.HubCollectorCMBRouterCounter.Inc(1)

	switch cvs.Request.PKType.Name {
	case "LOGIN":
		router.LoginOut <- cvs
	case "SEND_DEV_CONF_DATA":
		router.SendConfigOut <- cvs
	case "SEND_ALARM":
		router.SendAlarmOut <- cvs
	}
}
