package spout

import (
	"ark/hub/domain"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//COGSpout COG数据发送器
type COGSpout struct {
	flow.Component
	//BatchSize control collection of dto to optimize performance
	BatchSize int
	batchCOGs []*domain.COG
	In        chan *domain.COG
	COGOut1   chan<- []*domain.COG
	COGOut2   chan<- []*domain.COG
}

//Init init func for flow component
func (spout *COGSpout) Init() {
	// Your initialization code here
	spout.BatchSize = 1
}

//OnIn batch size control broadcast control
func (spout *COGSpout) OnIn(cog *domain.COG) {
	metrics.AppMetrics.HubBusCOGCounter.Inc(1)
	if spout.batchCOGs == nil {
		spout.batchCOGs = make([]*domain.COG, 0)
	}

	spout.batchCOGs = append(spout.batchCOGs, cog)
	if len(spout.batchCOGs) >= spout.BatchSize {
		spout.COGOut1 <- spout.batchCOGs
		spout.COGOut2 <- spout.batchCOGs
		spout.batchCOGs = nil
	}
}
