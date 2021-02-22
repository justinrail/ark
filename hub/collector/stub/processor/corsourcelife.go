package processor

import (
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/hub/govern"
	"ark/util/log"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//CoreSourceLifeSupervisor manage config status of coreSource
type CoreSourceLifeSupervisor struct {
	flow.Component
	In     chan *domain.COR
	COROut chan<- *domain.COR
}

//Init init func for flow component
func (supervisor *CoreSourceLifeSupervisor) Init() {
	// Your initialization code here
}

//OnIn batch size control broadcast control
func (supervisor *CoreSourceLifeSupervisor) OnIn(cor *domain.COR) {
	metrics.AppMetrics.HubCollectorStubCoreSourceLifeSupervisorCounter.Inc(1)
	if govern.CheckFuse(govern.FuseCoreSourceLifeSupervisor) {
		if cor.Flag >= enum.CoreSourceFlagUnkown && cor.Flag <= enum.CoreSourceFlagConDown {
			supervisor.COROut <- cor
		} else {
			log.Error("unkown message flag:" + enum.GetEnumString(cor.Flag))
		}
	}
}
