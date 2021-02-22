package entry

import (
	"ark/hub/collector/stub/dto"
	"ark/hub/collector/stub/graph"

	"github.com/robfig/cron"
	flow "github.com/trustmaster/goflow"
)

var dataIn chan *dto.Packet
var cr *cron.Cron

func init() {
	dataIn = make(chan *dto.Packet)
}

//Link link in ports of graph
func Link(net *graph.StubGraph) {
	net.SetInPort("In", dataIn)
}

//Ready load and init configs
func Ready() {
	cr = cron.New()
	//rndCOG := RandomCOG()
	cr.AddFunc("@every 1s", func() {
		dataIn <- RandomCOG()
		dataIn <- RandomCORPacket()
		dataIn <- RandomCOVPacket()
	})
}

//Start start simulate
func Start(net *graph.StubGraph) {
	flow.RunNet(net)
	cr.Start()
	<-net.Wait()
}
