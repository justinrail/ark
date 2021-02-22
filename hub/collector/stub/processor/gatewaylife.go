package processor

import (
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/hub/govern"
	"ark/util/log"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//GatewayLifeSupervisor manage config status of gateway
type GatewayLifeSupervisor struct {
	flow.Component
	In     chan *domain.COG
	COGOut chan<- *domain.COG
}

//Init init func for flow component
func (supervisor *GatewayLifeSupervisor) Init() {
	// Your initialization code here
}

//OnIn batch size control broadcast control
func (supervisor *GatewayLifeSupervisor) OnIn(cog *domain.COG) {
	metrics.AppMetrics.HubCollectorStubGatewayLifeSupervisorCounter.Inc(1)

	if govern.CheckFuse(govern.FuseGatewaylifeSupervisor) {
		if cog.Flag >= enum.GatewayFlagUnkown && cog.Flag <= enum.GatewayFlagSetInfoUpdateIntervalAck {

			supervisor.COGOut <- cog

		} else {
			log.Error("unkown message flag:" + enum.GetEnumString(cog.Flag))
		}
	}
}
