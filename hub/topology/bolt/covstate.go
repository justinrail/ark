package bolt

import (
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//COVStateUpdater 更新Corepoint的实际事件状态
type COVStateUpdater struct {
	flow.Component
	In                chan *domain.COV
	CoreLiteEventOut1 chan *domain.CoreLiteEvent
	CoreLiteEventOut2 chan *domain.CoreLiteEvent
}

//Init init func for flow component
func (updater *COVStateUpdater) Init() {
	// Your initialization code here

}

//OnIn Request handler
func (updater *COVStateUpdater) OnIn(cov *domain.COV) {
	metrics.AppMetrics.HubTopoCorePointStateUpdaterCounter.Inc(1)
	cp, ok := domain.CorePoints.Get(cov.CorePointID)

	if ok {
		corepoint := (cp.(*domain.CorePoint))
		if corepoint.StateRuleID == 2 {
			updater.handleStateRule2(corepoint, cov)
		}
		if corepoint.StateRuleID == 3 {
			updater.handleStateRule3(corepoint, cov)
		}
	}

}

//update point's event states
func (updater *COVStateUpdater) handleStateRule2(corepoint *domain.CorePoint, cov *domain.COV) {

	if cov.StateID == enum.CorePointFlagStart {
		if corepoint.CurrentEventState == 0 {

			coreLiteEvent := domain.NewCoreLiteEvent(corepoint)
			coreLiteEvent.EventChangeState = 1
			coreLiteEvent.StartTime = cov.Timestamp
			updater.CoreLiteEventOut1 <- coreLiteEvent
			updater.CoreLiteEventOut2 <- coreLiteEvent
			corepoint.CurrentEventState = 1
		}
	}

	if cov.StateID == enum.CorePointFlagEnd {
		if corepoint.CurrentEventState == 1 {
			coreLiteEvent := domain.NewCoreLiteEvent(corepoint)
			coreLiteEvent.EventChangeState = 3 //送出结束的liteevent
			//结束重置点开始时间（冗余这个开始时间用处不大，只是在搜索信号看告警时可以算得告警持续时间）
			corepoint.StartTime = 0
			coreLiteEvent.EndTime = cov.Timestamp

			updater.CoreLiteEventOut1 <- coreLiteEvent
			updater.CoreLiteEventOut2 <- coreLiteEvent
			corepoint.CurrentEventState = 0 //staterule = 2 时，告警状态只有有告警和无告警
		}
	}

}

//TODO
func (updater *COVStateUpdater) handleStateRule3(corepoint *domain.CorePoint, cov *domain.COV) {

}
