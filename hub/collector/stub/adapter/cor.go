package adapter

import (
	"ark/hub/bus"
	"ark/hub/domain"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//CORAdapter processor for dispatch cog to bus
type CORAdapter struct {
	flow.Component
	In chan *domain.COR
}

//OnIn Request handler
func (adapter *CORAdapter) OnIn(cor *domain.COR) {
	metrics.AppMetrics.HubCollectorStubCORAdapterCounter.Inc(1)
	bus.CORBus <- cor
}
