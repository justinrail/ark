package bolt

import (
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//COVDataUpdater 更新COV的实际状态，并进行修改和日志
type COVDataUpdater struct {
	flow.Component
	In       chan []*domain.COV
	StateOut chan *domain.COV
}

//OnIn Request handler
func (updater *COVDataUpdater) OnIn(covs []*domain.COV) {
	metrics.AppMetrics.HubTopoCOVDataUpdaterCounter.Inc(int64(len(covs)))
	for _, cov := range covs {
		if cov.StateID >= enum.CorePointFlagUnkown && cov.StateID <= enum.CorePointFlagEnd {
			updater.handleCOVData(cov.Clone())
		}
	}
}

func (updater *COVDataUpdater) handleCOVData(cov *domain.COV) {

	corepoint, here := domain.CorePoints.Get(cov.CorePointID)

	if here {
		(corepoint.(*domain.CorePoint)).UpdateData(cov)
		//实时计算可能性能 开销比较大，如果数据量非常大，可以使用定时巡检，而非主动更新，所以独立为函数方便屏蔽
		(corepoint.(*domain.CorePoint)).UpdateLimitState(cov)
		//过滤出事件，提供后来者进行处理
		updater.filterCOVState((corepoint.(*domain.CorePoint)), cov)
	}
}

func (updater *COVDataUpdater) filterCOVState(corePoint *domain.CorePoint, cov *domain.COV) {

	if cov.IsValid == true {
		switch cov.StateID {
		case enum.CorePointFlagStart:
			updater.StateOut <- cov.Clone()
		case enum.CorePointFlagConfirm:
			updater.StateOut <- cov.Clone()
		case enum.CorePointFlagEnd:
			updater.StateOut <- cov.Clone()
		}
	}

	corePoint.UpdateStateLogs(cov)
}
