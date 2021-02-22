package service

import (
	"ark/hub/domain"
	"ark/util/str"
	"ark/web/vm"
	"fmt"

	"gopkg.in/jeevatkm/go-model.v1"
)

//GetLiveEvent get all active live events
func GetLiveEvent() []vm.LiveEventView {
	items := make([]vm.LiveEventView, 0)

	cols := domain.CoreLiveEvents.Iter()
	for kv := range cols {
		cle := kv.Value.(*domain.CoreLiveEvent)
		clv := vm.LiveEventView{}
		model.Copy(&clv, cle)
		clv.CurrentNumericValueString = fmt.Sprintf("%.2f", cle.CurrentNumericValue)
		clv.StartTimeString = str.TimeToString(cle.StartTime)
		clv.EndTimeString = str.TimeToString(cle.EndTime)
		items = append(items, clv)
	}

	return items
}
