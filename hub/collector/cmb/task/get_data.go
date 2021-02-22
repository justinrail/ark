package task

import (
	"ark/hub/bus"
	"ark/hub/collector/cmb/dto"
	"ark/hub/domain"
	"ark/hub/enum"
	"ark/util/metrics"
	"ark/util/str"
)

func handleGetDataResponse(resp *dto.GetDataResponse) {
	covs := make([]*domain.COV, 0)

	if resp.Result == 0 {
		return
	}

	cg, ok := dto.CMBGateways.GetStringKey(resp.FSUID)
	if ok == false {
		return
	}

	gatewayID := (cg.(*dto.CMBGateway)).GatewayID
	gatewayUUID := (cg.(*dto.CMBGateway)).ID

	for _, dev := range resp.Values.DeviceAckList.Device {

		csID := dto.DeviceCoreSourceMap[gatewayUUID+"."+dev.DeviceID]

		if csID > 0 {
			for _, sig := range dev.TSemaphore {

				sigID := dto.SignalCorePointMap[gatewayUUID+"."+dev.DeviceID+"."+sig.ID+"."+sig.SignalNumber]
				//fmt.Println(sigID, gatewayUUID+"."+dev.DeviceID+"."+sig.ID+"."+sig.SignalNumber)
				if sigID > 0 {
					cov := &domain.COV{}
					cov.GateWayID = gatewayID
					cov.CoreSourceID = csID
					cov.CorePointID = sigID
					cov.CurrentNumericValue = sig.MeasuredVal
					cov.Timestamp = str.TimstampStringToTime(sig.Time)
					if sig.Status == 1 {
						cov.IsValid = false
					} else {
						cov.IsValid = true
					}
					cov.StateID = enum.CorePointFlagData

					covs = append(covs, cov)
				}
			}
		}

	}
	metrics.AppMetrics.HubCollectorCMBCOVCounter.Inc(int64(len(covs)))
	bus.COVBus <- covs
}
