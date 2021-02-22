package cmb

import (
	"ark/hub/collector/cmb/dto"
	"ark/hub/collector/cmb/task"
	"ark/hub/collector/cmb/ws"
	"ark/hub/domain"
	"ark/hub/util"
	"ark/util/sysevent"
	"net/http"

	flow "github.com/trustmaster/goflow"
)

var soapServer *http.Server
var graph *ProcessGraph

//Ready prepare for start
func Ready() {
	initGateways()
	soapServer = ws.NewSOAPServer(":8081")
	graph = NewGraph()
	graph.SetInPort("In", ws.DataIn)
	task.Ready()
	//sysevent.EventBus.SubscribeAsync("gateway:remove", removeGateway, false)
	sysevent.EventBus.Subscribe("gateway:remove", removeGateway)

}

func removeGateway(gatewayID int) {
	gs := dto.CMBGateways.Iter()
	gUUID := ""
	gID := 1

	for kv := range gs {
		cmbGateway := kv.Value.(*dto.CMBGateway)
		if cmbGateway.GatewayID == gatewayID {
			gID = cmbGateway.GatewayID
			gUUID = cmbGateway.ID
			break
		}
	}

	if len(gUUID) > 0 {
		dto.CMBGateways.Del(gUUID)

		css := domain.CoreSources.Iter()

		for cs := range css {
			coreSource := cs.Value.(*domain.CoreSource)
			if coreSource.GatewayID == gID {
				domain.CoreSources.Del(coreSource.CoreSourceID)
			}
		}

		cps := domain.CorePoints.Iter()

		for cp := range cps {
			corePoint := cp.Value.(*domain.CorePoint)
			if corePoint.GatewayID == gID {
				domain.CorePoints.Del(corePoint.CorePointID)
			}
		}

		domain.Gateways.Del(gID)

		domain.CascadeDeleteGateway(gID)
	}

}

//初始化加载gateway
func initGateways() {
	gateways := util.DeepLoadGatewayByCollector("CMB")
	for _, gateway := range gateways {
		domain.Gateways.Insert(gateway.ID, &gateway)
		cmbGateway := &dto.CMBGateway{}
		cmbGateway.GatewayID = gateway.ID
		cmbGateway.ID = gateway.UUID
		cmbGateway.IP = gateway.IP
		cmbGateway.Devices = make([]dto.Device, 0)
		dto.CMBGateways.Insert(cmbGateway.ID, cmbGateway)

		css := gateway.CoreSources.Iter()

		for kv := range css {
			cs := kv.Value.(*domain.CoreSource)
			dev := dto.Device{}
			dev.Signals.Signal = make([]dto.Signal, 0)
			dev.DeviceID = cs.UniqueID
			dev.DeviceName = cs.SourceName
			cmbGateway.Devices = append(cmbGateway.Devices, dev)

			//make map
			dto.DeviceCoreSourceMap[cs.UniqueID] = cs.CoreSourceID

			ps := cs.CorePoints.Iter()

			for kv2 := range ps {
				cp := kv2.Value.(*domain.CorePoint)
				sig := dto.Signal{}
				sig.ID = cp.UniqueID
				sig.NMAlarmID = cp.OriginStandardID
				dev.Signals.Signal = append(dev.Signals.Signal, sig)

				//make map
				dto.SignalCorePointMap[cp.UniqueID] = cp.CorePointID
			}
		}
	}
}

//Start collector begin collect
func Start() {
	go flow.RunNet(graph)
	go task.Run()
	soapServer.ListenAndServe()
	<-graph.Wait()
}
