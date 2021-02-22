package spout

import (
	"ark/hub/domain"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//COVSpout COV数据发送器
type COVSpout struct {
	flow.Component
	In      chan []*domain.COV
	COVOut1 chan<- []*domain.COV
	COVOut2 chan<- []*domain.COV
}

//OnIn batch size control broadcast control
func (spout *COVSpout) OnIn(covs []*domain.COV) {
	metrics.AppMetrics.HubBusCOVCounter.Inc(int64(len(covs)))
	spout.COVOut1 <- covs
	spout.COVOut2 <- covs

}
