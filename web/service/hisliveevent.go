package service

import (
	"ark/store/influxstore"
	"ark/util/str"
	"ark/web/vm"
	"fmt"

	"gopkg.in/jeevatkm/go-model.v1"
)

//GetHisLiveEvent get all hsitory live events
func GetHisLiveEvent() []vm.LiveEventView {
	items := make([]vm.LiveEventView, 0)

	events := influxstore.GetAllHisCoreLiveEvents()

	for _, event := range events {
		clv := vm.LiveEventView{}
		model.Copy(&clv, event)
		clv.CurrentNumericValueString = fmt.Sprintf("%.2f", event.CurrentNumericValue)
		clv.StartTimeString = str.TimeToString(event.StartTime)
		clv.EndTimeString = str.TimeToString(event.EndTime)
		clv.CurrentEventState = event.EventState
		items = append(items, clv)
	}
	return items
}
