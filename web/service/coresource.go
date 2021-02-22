package service

import (
	"ark/hub/domain"
	"ark/util/str"
	"ark/web/vm"

	"gopkg.in/jeevatkm/go-model.v1"
)

//GetCorePointByCoreSource get all corepoint by coresource
func GetCorePointByCoreSource(gatewayID int, coresourceID int) []vm.CorePointView {
	items := make([]vm.CorePointView, 0)

	obj, ok := domain.Gateways.Get(gatewayID)

	if ok {
		gateway := obj.(*domain.Gateway)

		cor, exist := gateway.CoreSources.Get(coresourceID)

		if exist {
			cps := (cor.(*domain.CoreSource)).CorePoints.Iter()

			for kv := range cps {
				corepoint := kv.Value.(*domain.CorePoint)
				cpv := vm.CorePointView{}
				model.Copy(&cpv, corepoint)
				cpv.UpdateTimeString = str.TimeToString(corepoint.UpdateTime)
				items = append(items, cpv)
			}
		}
	}

	return items
}
