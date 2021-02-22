package graph

import (
	"ark/hub/collector/stub/adapter"
	"ark/hub/collector/stub/processor"

	flow "github.com/trustmaster/goflow"
)

//StubGraph stub collector logic net
type StubGraph struct {
	flow.Graph
}

//NewGraph A constructor that creates network structure
func NewGraph() *StubGraph {
	// Create a new graph
	net := new(StubGraph)
	net.InitGraphState()
	// Add graph nodes
	net.Add(new(processor.Router), "router")

	net.Add(new(processor.CoreSourceLifeSupervisor), "coresource_life_supervisor")
	net.Add(new(processor.GatewayLifeSupervisor), "gateway_life_supervisor")

	net.Add(new(adapter.COGAdapter), "cog_adapter")
	net.Add(new(adapter.CORAdapter), "cor_adapter")
	net.Add(new(adapter.COVAdapter), "cov_adapter")

	//net connections
	net.Connect("router", "COGOut", "gateway_life_supervisor", "In")
	net.Connect("gateway_life_supervisor", "COGOut", "cog_adapter", "In")

	net.Connect("router", "COROut", "coresource_life_supervisor", "In")
	net.Connect("coresource_life_supervisor", "COROut", "cor_adapter", "In")

	net.Connect("router", "COVOut", "cov_adapter", "In")
	// Network ports
	net.MapInPort("In", "router", "In")

	//net.MapInPort("CORIn", "coresource_life_supervisor", "In")
	return net
}
