package bolt

import (
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//COVStateNormalizer 过滤乱序重复的cov
type COVStateNormalizer struct {
	flow.Component
	In       chan *domain.COV
	StateOut chan *domain.COV
}

//Init component init functions
func (normalizer *COVStateNormalizer) Init() {

}

//OnIn Request handler
func (normalizer *COVStateNormalizer) OnIn(cov *domain.COV) {
	metrics.AppMetrics.HubTopoCOVStateNormalizerCounter.Inc(1)
	switch cov.StateID {
	case enum.CorePointFlagStart:
		normalizer.handleStart(cov)
	case enum.CorePointFlagConfirm:
		normalizer.handleConfirm(cov)
	case enum.CorePointFlagEnd:
		normalizer.handleEnd(cov)
	}
}

func (normalizer *COVStateNormalizer) handleStart(cov *domain.COV) {
	cp, ok := domain.CorePoints.Get(cov.CorePointID)
	if ok {
		//如果无告警
		if (cp.(*domain.CorePoint)).CurrentEventState == 0 {
			(cp.(*domain.CorePoint)).StartTime = cov.Timestamp
			normalizer.StateOut <- cov
		}
	}
}

func (normalizer *COVStateNormalizer) handleConfirm(cov *domain.COV) {

	cp, exist := domain.CorePoints.Get(cov.CorePointID)
	if exist {
		//如果staterule为3，并且告警已经开始,而且非确认状态
		if (cp.(*domain.CorePoint)).StateRuleID == 3 && (cp.(*domain.CorePoint)).CurrentEventState == 1 && (cp.(*domain.CorePoint)).CurrentEventState != 2 {
			normalizer.StateOut <- cov
		}
	}
}

func (normalizer *COVStateNormalizer) handleEnd(cov *domain.COV) {

	cp, exist := domain.CorePoints.Get(cov.CorePointID)
	if exist {
		//如果告警发生则结束它
		if (cp.(*domain.CorePoint)).StateRuleID == 2 && (cp.(*domain.CorePoint)).CurrentEventState == 1 {
			(cp.(*domain.CorePoint)).EndTime = cov.Timestamp
			normalizer.StateOut <- cov

			//如果告警规则为3，而且未结束
		} else if (cp.(*domain.CorePoint)).StateRuleID == 3 && (cp.(*domain.CorePoint)).CurrentEventState != 3 {
			normalizer.StateOut <- cov
		}
	}
}
