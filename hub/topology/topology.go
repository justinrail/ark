package topology

import (
	"ark/hub/bus"
	"ark/hub/topology/graph"
	"ark/hub/topology/notifier"
	"ark/util/cfg"

	flow "github.com/trustmaster/goflow"
)

var topologyGraph *graph.TopologyGraph

//Ready init topology graph and connect to bus
func Ready() {
	topologyGraph = graph.NewGraph()

	//link
	topologyGraph.SetInPort("cog_in", bus.COGBus)
	topologyGraph.SetInPort("cor_in", bus.CORBus)
	topologyGraph.SetInPort("cov_in", bus.COVBus)

	if cfg.Read().Hub.NotifyRunFuse {
		notifier.Ready()
		go notifier.Start()
	}
}

//Start to run topology
func Start() {
	flow.RunNet(topologyGraph)
}
