package hub

import (
	"ark/hub/collector/cmb"
	"ark/hub/collector/stub"
	"ark/hub/job"
	"ark/hub/topology"
	"ark/util/cfg"

	evbus "github.com/asaskevich/EventBus"
)

//EventBus EventBus
var EventBus evbus.Bus

//Run Prepare and Run
func Run() {
	EventBus = evbus.New()

	if cfg.Read().Hub.CollectorStubRunFuse {
		stub.Ready()
		go stub.Start()
	}

	if cfg.Read().Hub.CollectorCMBRunFuse {
		cmb.Ready()
		go cmb.Start()
	}

	topology.Ready()
	go topology.Start()

	job.Ready()
	job.Start()
}
