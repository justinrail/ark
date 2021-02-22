package bolt

import (
	"ark/hub/domain"
	"ark/store/influxstore"
	"ark/util/cfg"
	"ark/util/metrics"

	flow "github.com/trustmaster/goflow"
)

//HisCoreLiveEventAppender 存历史告警
type HisCoreLiveEventAppender struct {
	flow.Component
	In chan *domain.HisCoreLiveEvent
}

//OnIn Request HisCoreLiveEventAppender
func (appender *HisCoreLiveEventAppender) OnIn(hisLiveEvent *domain.HisCoreLiveEvent) {
	if cfg.Read().Hub.BoltHisEventRunFuse {
		metrics.AppMetrics.HubTopoHisCoreActiveEventAppenderCounter.Inc(1)
		influxstore.AppendHisCoreLiveEvent(hisLiveEvent)
	}
}
