package adapter

import (
	"ark/hub/bus"
	"ark/hub/domain"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//COGAdapter processor for dispatch cog to bus
type COGAdapter struct {
	flow.Component
	In chan *domain.COG
}

//OnIn Request handler
func (adapter *COGAdapter) OnIn(cog *domain.COG) {
	metrics.AppMetrics.HubCollectorStubCOGAdapterCounter.Inc(1)
	bus.COGBus <- cog
}
