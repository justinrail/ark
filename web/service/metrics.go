package service

import (
	"ark/hub/bus"
	"ark/hub/sink"
	"ark/util/metrics"
	"fmt"
)

//GetMetrics get all metrics
func GetMetrics() map[string]string {
	items := make(map[string]string)

	items["HubBusCOGCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubBusCOGCounter.Count())
	items["HubBusCOGPercent"] = fmt.Sprintf("%.2f", float32((len(bus.COGBus)/bus.QueueSize))*100)
	items["HubBusCORCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubBusCORCounter.Count())
	items["HubBusCORPercent"] = fmt.Sprintf("%.2f", float32((len(bus.CORBus)/bus.QueueSize))*100)
	items["HubBusCOVCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubBusCOVCounter.Count())
	items["HubBusCOVPercent"] = fmt.Sprintf("%.2f", float32((len(bus.COVBus)/bus.QueueSize))*100)

	items["HubSinkHisCorePointCounter"] = fmt.Sprintf("%d", sink.HisCorePointSinkQueue.Total)
	items["HubSinkHisCorePointLostCounter"] = fmt.Sprintf("%d", sink.HisCorePointSinkQueue.Lost)
	items["HubSinkHisCorePointPercent"] = fmt.Sprintf("%.2f", float32(sink.HisCorePointSinkQueue.Count()/sink.HisCorePointSinkQueue.Length)*100)

	items["HubTopoCOGStateUpdaterCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubTopoCOGStateUpdaterCounter.Count())
	items["HubTopoCORStateUpdaterCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubTopoCORStateUpdaterCounter.Count())
	items["HubTopoCOVDataUpdaterCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubTopoCOVDataUpdaterCounter.Count())
	items["HubTopoCOVStateNormalizerCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubTopoCOVStateNormalizerCounter.Count())
	items["HubTopoCorePointStateUpdaterCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubTopoCorePointStateUpdaterCounter.Count())
	items["HubTopoCoreActiveEventUpdaterCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubTopoCoreActiveEventUpdaterCounter.Count())
	items["HubTopoHisCoreActiveEventAppenderCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubTopoHisCoreActiveEventAppenderCounter.Count())

	items["HubCollectorStubRouterCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubCollectorStubRouterCounter.Count())
	items["HubCollectorStubGatewayLifeSupervisorCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubCollectorStubGatewayLifeSupervisorCounter.Count())
	items["HubCollectorStubCoreSourceLifeSupervisorCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubCollectorStubCoreSourceLifeSupervisorCounter.Count())
	items["HubCollectorStubCOGAdapterCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubCollectorStubCOGAdapterCounter.Count())
	items["HubCollectorStubCORAdapterCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubCollectorStubCORAdapterCounter.Count())
	items["HubCollectorStubCOVAdapterCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubCollectorStubCOVAdapterCounter.Count())

	items["HubSinkHisComplexIndexCounter"] = fmt.Sprintf("%d", sink.HisComplexIndexSinkQueue.Total)
	items["HubSinkHisComplexIndexLostCounter"] = fmt.Sprintf("%d", sink.HisComplexIndexSinkQueue.Lost)
	items["HubSinkHisComplexIndexPercent"] = fmt.Sprintf("%.2f", float32(sink.HisComplexIndexSinkQueue.Count()/sink.HisComplexIndexSinkQueue.Length)*100)

	items["HubCollectorCMBLoginProcessorCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubCollectorCMBLoginProcessorCounter.Count())
	items["HubCollectorCMBRouterCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubCollectorCMBRouterCounter.Count())
	items["HubCollectorCMBCOGCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubCollectorCMBCOGCounter.Count())
	items["HubCollectorCMBSendDataProcessorCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubCollectorCMBSendDataProcessorCounter.Count())
	items["HubCollectorCMBSendConfigProcessorCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubCollectorCMBSendConfigProcessorCounter.Count())
	items["HubCollectorCMBGetDataTaskCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubCollectorCMBGetDataTaskCounter.Count())
	items["HubCollectorCMBCOVCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubCollectorCMBCOVCounter.Count())
	items["HubCollectorCMBSendAlarmProcessorCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubCollectorCMBSendAlarmProcessorCounter.Count())

	items["HubTopoPhoenixCOVCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubTopoPhoenixCOVCounter.Count())
	items["HubTopoPhoenixCORCounter"] = fmt.Sprintf("%d", metrics.AppMetrics.HubTopoPhoenixCORCounter.Count())

	return items
}
