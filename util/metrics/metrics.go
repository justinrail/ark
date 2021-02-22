package metrics

import "github.com/rcrowley/go-metrics"

//ArkMetrics 指标监控对象
type ArkMetrics struct {
	HubBusCOGCounter                                metrics.Counter
	HubBusCORCounter                                metrics.Counter
	HubBusCOVCounter                                metrics.Counter
	HubTopoCOGStateUpdaterCounter                   metrics.Counter
	HubTopoCORStateUpdaterCounter                   metrics.Counter
	HubTopoCOVStateUpdaterCounter                   metrics.Counter
	HubTopoCOVDataUpdaterCounter                    metrics.Counter
	HubTopoCOVStateNormalizerCounter                metrics.Counter
	HubTopoCorePointStateUpdaterCounter             metrics.Counter
	HubTopoCoreActiveEventUpdaterCounter            metrics.Counter
	HubTopoHisCoreActiveEventAppenderCounter        metrics.Counter
	HubCollectorStubRouterCounter                   metrics.Counter
	HubCollectorStubGatewayLifeSupervisorCounter    metrics.Counter
	HubCollectorStubCoreSourceLifeSupervisorCounter metrics.Counter
	HubCollectorStubCOGAdapterCounter               metrics.Counter
	HubCollectorStubCORAdapterCounter               metrics.Counter
	HubCollectorStubCOVAdapterCounter               metrics.Counter
	HubCollectorCMBLoginProcessorCounter            metrics.Counter
	HubCollectorCMBCOGCounter                       metrics.Counter
	HubCollectorCMBSendDataProcessorCounter         metrics.Counter
	HubCollectorCMBSendConfigProcessorCounter       metrics.Counter
	HubCollectorCMBGetDataTaskCounter               metrics.Counter
	HubCollectorCMBCOVCounter                       metrics.Counter
	HubCollectorCMBSendAlarmProcessorCounter        metrics.Counter
	HubCollectorCMBRouterCounter                    metrics.Counter
	HubTopoPhoenixCOVCounter                        metrics.Counter
	HubTopoPhoenixCORCounter                        metrics.Counter
}

//AppMetrics 系统的总的指标监控对象
var AppMetrics *ArkMetrics

func init() {
	AppMetrics = &ArkMetrics{}
	AppMetrics.HubBusCOGCounter = metrics.NewCounter()
	metrics.Register("hub.bus.cog.counter", AppMetrics.HubBusCOGCounter)
	AppMetrics.HubBusCORCounter = metrics.NewCounter()
	metrics.Register("hub.bus.cor.counter", AppMetrics.HubBusCORCounter)
	AppMetrics.HubBusCOVCounter = metrics.NewCounter()
	metrics.Register("hub.bus.cov.counter", AppMetrics.HubBusCOVCounter)

	AppMetrics.HubTopoCOGStateUpdaterCounter = metrics.NewCounter()
	metrics.Register("hub.topo.cog.stateupdater.counter", AppMetrics.HubTopoCOGStateUpdaterCounter)
	AppMetrics.HubTopoCORStateUpdaterCounter = metrics.NewCounter()
	metrics.Register("hub.topo.cor.stateupdater.counter", AppMetrics.HubTopoCORStateUpdaterCounter)
	AppMetrics.HubTopoCOVDataUpdaterCounter = metrics.NewCounter()
	metrics.Register("hub.topo.cov.dataupdater.counter", AppMetrics.HubTopoCOVDataUpdaterCounter)
	AppMetrics.HubTopoCOVStateNormalizerCounter = metrics.NewCounter()
	metrics.Register("hub.topo.cov.statenormalizer.counter", AppMetrics.HubTopoCOVStateNormalizerCounter)
	AppMetrics.HubTopoCorePointStateUpdaterCounter = metrics.NewCounter()
	metrics.Register("hub.topo.corepoint.stateupdater.counter", AppMetrics.HubTopoCorePointStateUpdaterCounter)
	AppMetrics.HubTopoCoreActiveEventUpdaterCounter = metrics.NewCounter()
	metrics.Register("hub.topo.coreactiveevent.updater.counter", AppMetrics.HubTopoCoreActiveEventUpdaterCounter)
	AppMetrics.HubTopoHisCoreActiveEventAppenderCounter = metrics.NewCounter()
	metrics.Register("hub.topo.hiscoreactiveevent.appender.counter", AppMetrics.HubTopoHisCoreActiveEventAppenderCounter)

	AppMetrics.HubCollectorStubRouterCounter = metrics.NewCounter()
	metrics.Register("hub.collector.stub.router.counter", AppMetrics.HubCollectorStubRouterCounter)
	AppMetrics.HubCollectorStubGatewayLifeSupervisorCounter = metrics.NewCounter()
	metrics.Register("hub.collector.stub.gateway.lifesupervisor.counter", AppMetrics.HubCollectorStubGatewayLifeSupervisorCounter)
	AppMetrics.HubCollectorStubCoreSourceLifeSupervisorCounter = metrics.NewCounter()
	metrics.Register("hub.collector.stub.coresource.lifesupervisor.counter", AppMetrics.HubCollectorStubCoreSourceLifeSupervisorCounter)
	AppMetrics.HubCollectorStubCOGAdapterCounter = metrics.NewCounter()
	metrics.Register("hub.collector.stub.cog.adapter.counter", AppMetrics.HubCollectorStubCOGAdapterCounter)
	AppMetrics.HubCollectorStubCORAdapterCounter = metrics.NewCounter()
	metrics.Register("hub.collector.stub.cor.adapter.counter", AppMetrics.HubCollectorStubCORAdapterCounter)
	AppMetrics.HubCollectorStubCOVAdapterCounter = metrics.NewCounter()
	metrics.Register("hub.collector.stub.cov.adapter.counter", AppMetrics.HubCollectorStubCOVAdapterCounter)

	AppMetrics.HubCollectorCMBLoginProcessorCounter = metrics.NewCounter()
	metrics.Register("hub.collector.cmb.login.processor.counter", AppMetrics.HubCollectorCMBLoginProcessorCounter)
	AppMetrics.HubCollectorCMBRouterCounter = metrics.NewCounter()
	metrics.Register("hub.collector.cmb.router.counter", AppMetrics.HubCollectorCMBRouterCounter)
	AppMetrics.HubCollectorCMBCOGCounter = metrics.NewCounter()
	metrics.Register("hub.collector.cmb.cog.counter", AppMetrics.HubCollectorCMBCOGCounter)
	AppMetrics.HubCollectorCMBSendDataProcessorCounter = metrics.NewCounter()
	metrics.Register("hub.collector.cmb.senddata.processor.counter", AppMetrics.HubCollectorCMBSendDataProcessorCounter)
	AppMetrics.HubCollectorCMBSendConfigProcessorCounter = metrics.NewCounter()
	metrics.Register("hub.collector.cmb.sendconfig.processor.counter", AppMetrics.HubCollectorCMBSendConfigProcessorCounter)
	AppMetrics.HubCollectorCMBGetDataTaskCounter = metrics.NewCounter()
	metrics.Register("hub.collector.cmb.getdata.task.counter", AppMetrics.HubCollectorCMBGetDataTaskCounter)
	AppMetrics.HubCollectorCMBCOVCounter = metrics.NewCounter()
	metrics.Register("hub.collector.cmb.cov.processor.counter", AppMetrics.HubCollectorCMBCOVCounter)
	AppMetrics.HubCollectorCMBSendAlarmProcessorCounter = metrics.NewCounter()
	metrics.Register("hub.collector.cmb.sendalarm.processor.counter", AppMetrics.HubCollectorCMBSendAlarmProcessorCounter)

	AppMetrics.HubTopoPhoenixCOVCounter = metrics.NewCounter()
	metrics.Register("hub.topo.phoenix.cov.counter", AppMetrics.HubTopoPhoenixCOVCounter)
	AppMetrics.HubTopoPhoenixCORCounter = metrics.NewCounter()
	metrics.Register("hub.topo.phoenix.cor.counter", AppMetrics.HubTopoPhoenixCORCounter)

}
