package service

import "ark/hub/domain"

//GetCoreSourceByGateway get all coresoruce by gateway
func GetCoreSourceByGateway(gatewayID int) []domain.CoreSource {
	items := make([]domain.CoreSource, 0)

	obj, ok := domain.Gateways.Get(gatewayID)

	if ok {
		gateway := obj.(*domain.Gateway)

		crs := gateway.CoreSources.Iter()

		for kv := range crs {
			coreSource := kv.Value.(*domain.CoreSource)
			items = append(items, *coreSource)
		}
	}

	return items
}
