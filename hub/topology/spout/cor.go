package spout

import (
	"ark/hub/domain"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//CORSpout COR数据发送器
type CORSpout struct {
	flow.Component
	//BatchSize control collection of dto to optimize performance
	BatchSize int
	batchCORs []*domain.COR
	In        chan *domain.COR
	COROut1   chan<- []*domain.COR
	COROut2   chan<- []*domain.COR
}

//Init init func for flow component
func (spout *CORSpout) Init() {
	// Your initialization code here
	spout.BatchSize = 1
}

//OnIn batch size control broadcast control
func (spout *CORSpout) OnIn(cor *domain.COR) {
	metrics.AppMetrics.HubBusCORCounter.Inc(1)
	if spout.batchCORs == nil {
		spout.batchCORs = make([]*domain.COR, 0)
	}

	spout.batchCORs = append(spout.batchCORs, cor)
	if len(spout.batchCORs) >= spout.BatchSize {
		spout.COROut1 <- spout.batchCORs
		spout.COROut2 <- spout.batchCORs
		spout.batchCORs = nil
	}
}
