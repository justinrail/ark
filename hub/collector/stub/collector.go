package stub

import (
	"ark/hub/collector/stub/entry"
	"ark/hub/collector/stub/graph"
	"ark/hub/domain"
	"ark/hub/util"
)

var stubGraph *graph.StubGraph

func loadGateways() {
	gateways := util.DeepLoadGatewayByCollector("stub")
	for _, gateway := range gateways {
		domain.Gateways.Insert(gateway.ID, &gateway)
	}
}

//Ready prepare for start
func Ready() {
	loadGateways()
	stubGraph = graph.NewGraph()

	entry.Link(stubGraph)
	entry.Ready()
}

//Start collector begin collect
func Start() {
	entry.Start(stubGraph)
}
