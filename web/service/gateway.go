package service

import (
	"ark/hub/domain"
	"ark/store/mysql/repo"
	"ark/util/sysevent"
)

//GetAllGateway Get all loaded gateway
func GetAllGateway() []domain.Gateway {
	items := make([]domain.Gateway, 0)

	cols := domain.Gateways.Iter()
	for kv := range cols {
		gateway := kv.Value.(*domain.Gateway)
		items = append(items, *gateway)
	}

	return items
}

//GetGatewayByID GetGatewayByID
func GetGatewayByID(id int) *domain.Gateway {
	gw, ok := domain.Gateways.Get(id)
	if ok {
		return gw.(*domain.Gateway)
	}
	return nil
}

//GetGatewaysByCollector GetGatewaysByCollector
func GetGatewaysByCollector(collectorName string) []domain.Gateway {
	items := make([]domain.Gateway, 0)

	cols := domain.Gateways.Iter()
	for kv := range cols {
		gateway := kv.Value.(*domain.Gateway)
		if gateway.Collector == collectorName {
			items = append(items, *gateway)
		}
	}

	return items
}

//DeleteGatewayByGatewayID DeleteGatewayByGatewayID
func DeleteGatewayByGatewayID(gatewayID int) {
	sysevent.EventBus.Publish("gateway:remove", gatewayID)
}

//UpdateGatewayJoined UpdateGatewayJoined
func UpdateGatewayJoined(gatewayID int, joined bool) *domain.Gateway {
	gw, ok := domain.Gateways.Get(gatewayID)
	if ok {
		gw.(*domain.Gateway).Joined = joined
		repo.UpdateGatewayJoined(gatewayID, joined)
	}

	return gw.(*domain.Gateway)
}
