package util

import (
	"ark/hub/domain"
	"ark/store/entity"
	"ark/store/mysql/repo"

	"gopkg.in/jeevatkm/go-model.v1"
)

//DeepLoadGatewayByCollector load and build domain gateway
func DeepLoadGatewayByCollector(collectorName string) []domain.Gateway {
	gateways := repo.GetGatewaysByCollector(collectorName)

	res := make([]domain.Gateway, len(gateways))
	for i, gateway := range gateways {
		domainGateway := domain.NewGateway()
		model.Copy(&domainGateway, gateway)
		res[i] = domainGateway
		deepLoadCoreSourceByGateway(&domainGateway)
	}
	return res
}

//DeepLoadGateway DeepLoadGateway
func DeepLoadGateway(gateway *entity.Gateway) *domain.Gateway {
	domainGateway := domain.NewGateway()
	model.Copy(&domainGateway, gateway)
	deepLoadCoreSourceByGateway(&domainGateway)

	return &domainGateway
}

func deepLoadCoreSourceByGateway(gateway *domain.Gateway) {
	coresources := repo.GetCoreSourceByGateway(gateway.ID)

	for _, cr := range coresources {
		domainCoresource := domain.NewCoreSource()
		model.Copy(&domainCoresource, cr)
		gateway.CoreSources.Insert(cr.CoreSourceID, &domainCoresource)
		domain.CoreSources.Insert(cr.CoreSourceID, &domainCoresource)
		deepLoadCorePointByCoreSource(&domainCoresource)
	}
}

func deepLoadCorePointByCoreSource(coresource *domain.CoreSource) {
	corepoints := repo.GetCorePointByByCoreSource(coresource.CoreSourceID)

	for _, cp := range corepoints {
		domainCorePoint := domain.NewCorePoint(coresource)
		model.Copy(domainCorePoint, cp)
		coresource.CorePoints.Insert(cp.CorePointID, domainCorePoint)
		domain.CorePoints.Insert(cp.CorePointID, domainCorePoint)
	}
}
