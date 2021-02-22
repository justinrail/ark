package adapter

import (
	"ark/hub/bus"
	"ark/hub/domain"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//COVAdapter processor for dispatch cog to bus
type COVAdapter struct {
	flow.Component
	In chan []*domain.COV
}

//OnIn Request handler
func (adapter *COVAdapter) OnIn(covs []*domain.COV) {
	metrics.AppMetrics.HubCollectorStubCOVAdapterCounter.Inc(int64(len(covs)))
	bus.COVBus <- covs
}
