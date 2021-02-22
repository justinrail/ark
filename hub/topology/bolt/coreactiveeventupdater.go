package bolt

import (
	"ark/hub/domain"
	"ark/util/metrics"
	"fmt"

	flow "github.com/trustmaster/goflow"
	"gopkg.in/jeevatkm/go-model.v1"
)

//CoreActiveEventUpdater 更新CoreActiveEvent的实际状态
type CoreActiveEventUpdater struct {
	flow.Component
	In                  chan *domain.CoreLiteEvent
	HisCoreLiveEventOut chan *domain.HisCoreLiveEvent
}

//OnIn Request handler only staterule =2
func (updater *CoreActiveEventUpdater) OnIn(coreLiteEvent *domain.CoreLiteEvent) {
	metrics.AppMetrics.HubTopoCoreActiveEventUpdaterCounter.Inc(1)
	activeEvent := domain.CoreLiveEvent{}
	model.Copy(&activeEvent, coreLiteEvent)
	activeEvent.CurrentEventState = coreLiteEvent.EventChangeState

	oldEvent, ok := domain.CoreLiveEvents.GetOrInsert(coreLiteEvent.CorePointID, &activeEvent)

	if ok {
		existEvent := oldEvent.(*domain.CoreLiveEvent)
		activeEvent.StartTime = existEvent.StartTime
		domain.CoreLiveEvents.Del(coreLiteEvent.CorePointID)
		updater.HisCoreLiveEventOut <- domain.NewHisCoreLiveEvent(&activeEvent)
	} else if activeEvent.CurrentEventState == 1 {
		domain.CoreLiveEvents.Insert(coreLiteEvent.CorePointID, &activeEvent)
	} else if activeEvent.CurrentEventState == 3 {
		fmt.Println(activeEvent)
	}
}
