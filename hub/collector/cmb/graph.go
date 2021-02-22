package cmb

import (
	"ark/hub/collector/cmb/processor"

	flow "github.com/trustmaster/goflow"
)

//ProcessGraph cmb collector logic net
type ProcessGraph struct {
	flow.Graph
}

//NewGraph A constructor that creates network structure
func NewGraph() *ProcessGraph {
	// Create a new graph
	net := new(ProcessGraph)
	net.InitGraphState()
	// Add graph nodes
	net.Add(new(processor.Router), "router")

	net.Add(new(processor.LoginProcessor), "login_processor")
	net.Add(new(processor.SendConfigProcessor), "send_config_processor")
	net.Add(new(processor.SendAlarmProcessor), "send_alarm_processor")

	//net connections
	net.Connect("router", "LoginOut", "login_processor", "In")
	net.Connect("router", "SendConfigOut", "send_config_processor", "In")
	net.Connect("router", "SendAlarmOut", "send_alarm_processor", "In")

	// Network ports
	net.MapInPort("In", "router", "In")

	return net
}
