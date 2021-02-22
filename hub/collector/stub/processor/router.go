package processor

import (
	"ark/hub/collector/stub/dto"
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//Router processor for dispatch
type Router struct {
	flow.Component
	In     chan *dto.Packet
	COGOut chan<- *domain.COG
	COROut chan<- *domain.COR
	COVOut chan<- []*domain.COV
}

//OnIn Request handler
func (router *Router) OnIn(obj *dto.Packet) {
	metrics.AppMetrics.HubCollectorStubRouterCounter.Inc(1)

	switch obj.MessageType {
	case enum.PacketCOG:
		cog := obj.Body.(domain.COG)
		router.COGOut <- &cog
	case enum.PacketCOR:
		cor := obj.Body.(domain.COR)
		router.COROut <- &cor
	case enum.PacketCOV:
		covs := obj.Body.([]*domain.COV)
		router.COVOut <- covs
	}
}
