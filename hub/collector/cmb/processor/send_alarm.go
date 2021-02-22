package processor

import (
	"ark/hub/bus"
	"ark/hub/collector/cmb/dto"
	"ark/hub/collector/cmb/ws"
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/util/metrics"
	"ark/util/str"

	flow "github.com/trustmaster/goflow"
)

//SendAlarmProcessor SendAlarmProcessor
type SendAlarmProcessor struct {
	flow.Component
	In chan *ws.Conversation
}

func (sendAlarmProcessor *SendAlarmProcessor) sendAlarm() {

}

//OnIn Request handler
func (sendAlarmProcessor *SendAlarmProcessor) OnIn(cvs *ws.Conversation) {

	req := &dto.SendAlarmRequest{}
	cvs.GetRequest(req)

	covs := make([]*domain.COV, 0)

	//根据TAlarm生成COV，发送到Bus
	for _, alarm := range req.Values.TAlarmList.TAlarm {
		metrics.AppMetrics.HubCollectorCMBSendAlarmProcessorCounter.Inc(1)
		metrics.AppMetrics.HubCollectorCMBCOVCounter.Inc(1)

		cov := &domain.COV{}

		g, exist := dto.CMBGateways.Get(req.FSUID)
		if exist == false {
			continue
		}

		csID := dto.DeviceCoreSourceMap[req.FSUID+"."+alarm.DeviceID]

		if csID < 0 {
			continue
		}

		cpID := dto.SignalCorePointMap[req.FSUID+"."+alarm.DeviceID+"."+alarm.ID+"."+alarm.SignalNumber]

		if cpID < 0 {
			continue
		}

		cov.GateWayID = (g.(*dto.CMBGateway)).GatewayID
		cov.CoreSourceID = csID
		cov.CorePointID = cpID
		cov.CurrentNumericValue = alarm.EventValue
		cov.IsValid = true
		cov.Timestamp = str.TimstampStringToTime(alarm.AlarmTime)
		if alarm.AlarmFlag == "BEGIN" {
			cov.StateID = enum.CorePointFlagStart
		}
		if alarm.AlarmFlag == "END" {
			cov.StateID = enum.CorePointFlagEnd
		}
		covs = append(covs, cov)
	}

	bus.COVBus <- covs

	// fmt.Println(cvs.Request.Info.InnerText)
	// fmt.Println(req.Values.TAlarmList.TAlarm[0])

	rep := &dto.SendAlarmResponse{}
	rep.FSUID = req.FSUID
	rep.Result = 1
	rep.FailureCause = "NULL"

	cvs.SendResponse("SEND_ALARM_ACK", rep)

}
