package bolt

import (
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/util/log"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//CORStateUpdater 更新COR的实际状态，并进行修改和日志
type CORStateUpdater struct {
	flow.Component
	In chan []*domain.COR
}

//OnIn Request handler
func (updater *CORStateUpdater) OnIn(cors []*domain.COR) {
	metrics.AppMetrics.HubTopoCORStateUpdaterCounter.Inc(int64(len(cors)))
	for _, cor := range cors {
		if cor.Flag >= enum.CoreSourceFlagUnkown && cor.Flag <= enum.CoreSourceFlagConDown {
			updater.handleCORState(cor.Clone())
		} else {
			log.Error("unkown message flag:" + enum.GetEnumString(cor.Flag))
		}
	}
}

func (updater *CORStateUpdater) handleCORState(cor *domain.COR) {
	gateway, ok := domain.Gateways.Get(cor.GateWayID)

	if ok {

		corsource, exist := (gateway.(*domain.Gateway)).CoreSources.Get(cor.CoreSourceID)

		if exist {
			(corsource.(*domain.CoreSource)).UpdateState(cor)
		}
	}
}
