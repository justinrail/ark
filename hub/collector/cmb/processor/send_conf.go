package processor

import (
	"ark/hub/bus"
	"ark/hub/collector/cmb/dto"
	"ark/hub/collector/cmb/ws"
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/hub/util"
	"ark/store/mysql/repo"
	"ark/util/metrics"
	"time"

	flow "github.com/trustmaster/goflow"
)

//SendConfigProcessor SendConfigProcessor
type SendConfigProcessor struct {
	flow.Component
	In chan *ws.Conversation
}

func (sendConfigProcessor *SendConfigProcessor) saveSignal(coreSourceID int, sigs []dto.Signal, prefix string) {
	// DI＝4	数字输入量（包含多态数字输入量），遥信
	// AI＝3	模拟输入量，遥测
	// DO＝1	数字输出量，遥控
	// AO＝2	模拟输出量，遥调
	// ALARM= 0 告警
	for _, sig := range sigs {

		dataType := 1
		writable := false
		readable := false
		stateRule := 0
		eventSeverity := 0

		switch sig.Type {
		case 0:
			dataType = 1
			readable = false
			writable = false
			stateRule = 2
			eventSeverity = sig.AlarmLevel
		case 1:
			dataType = 1
			readable = true
			writable = false
		case 2:
			dataType = 2
			readable = true
			writable = false
		case 3:
			dataType = 2
			readable = true
			writable = true
		case 4:
			dataType = 1
			readable = true
			writable = true
		}

		uuid := prefix + "." + sig.ID + "." + sig.SignalNumber
		corePointID := repo.AddCorePoint(coreSourceID, sig.SignalName, uuid, dataType, eventSeverity, sig.NMAlarmID, readable, writable, stateRule)
		//dto.SignalCorePointMap[sig.ID] = corePointID
		dto.SignalCorePointMap[uuid] = corePointID
		//sigID := dto.SignalCorePointMap[gatewayUUID+"."+dev.DeviceID+"."+sig.ID+"."+sig.SignalNumber]
		//fmt.Println(sigID, gatewayUUID+"."+dev.DeviceID+"."+sig.ID+"."+sig.SignalNumber)
	}
}

//OnIn Request handler
func (sendConfigProcessor *SendConfigProcessor) OnIn(cvs *ws.Conversation) {
	metrics.AppMetrics.HubCollectorCMBSendConfigProcessorCounter.Inc(1)

	req := &dto.SendConfigRequest{}
	cvs.GetRequest(req)

	//从现有的Gateways中查找这个gateway的UUID==FSUID
	_, ok := dto.CMBGateways.GetStringKey(req.FSUID)
	if ok == false {
		//如果不存在则创建CMBGateway
		cmbGateway := &dto.CMBGateway{}
		cmbGateway.ID = req.FSUID
		cmbGateway.Devices = req.Values.Device
		cmbGateway.IP = cvs.GetClientIP()
		dto.CMBGateways.Insert(cmbGateway.ID, cmbGateway)

		// //如果gateway在数据库已经存在则进行加载后进行映射表初始化
		// eg := repo.FindGatewayByUUID(req.FSUID)
		// if eg.ID > 0 {
		// 	dg, exist := domain.Gateways.Get(eg.ID)
		// 	if exist {
		// 		cmbGateway.GatewayID = (dg.(*domain.Gateway)).ID
		// 	}
		// } else {

		// }

		//根据CMBGateway的配置刷入Gateway及下级配置到数据
		gateway := repo.AddGateway(cmbGateway.ID, cmbGateway.ID, "CMB", cvs.GetClientIP())
		cmbGateway.GatewayID = gateway.ID

		//循环增加设备和信号
		for _, dev := range cmbGateway.Devices {
			coreSourceID := repo.AddCoreSource(gateway.ID, dev.DeviceName, cmbGateway.ID+"."+dev.DeviceID)
			dto.DeviceCoreSourceMap[cmbGateway.ID+"."+dev.DeviceID] = coreSourceID
			sendConfigProcessor.saveSignal(coreSourceID, dev.Signals.Signal, cmbGateway.ID+"."+dev.DeviceID)
		}
		//添加gateway到内存
		dg := util.DeepLoadGateway(gateway)
		domain.Gateways.Insert(dg.ID, dg)
	} else {
		// 目前没有规则确定是否要配置同步，如果有规则则可设置SynState，标志配置是否同步，先不做细节检查，太麻烦。
		// 如果需要同步，则进行同步，暂不开发
	}

	cg, good := dto.CMBGateways.GetStringKey(req.FSUID)
	if good {
		// 发送COG到bus
		cog := &domain.COG{}
		cog.ID = (cg.(*dto.CMBGateway)).GatewayID
		cog.Name = req.FSUID
		cog.Timestamp = time.Now().Unix()
		cog.Flag = enum.GatewayFlagSendConfig
		metrics.AppMetrics.HubCollectorCMBCOGCounter.Inc(1)
		bus.COGBus <- cog
	}

	scRep := &dto.SendConfigResponse{}
	scRep.FSUID = req.FSUID
	scRep.Result = 1
	scRep.FailureCause = "NULL"

	cvs.SendResponse("SEND_DEV_CONF_DATA_ACK", scRep)

}
