package bolt

import (
	"ark/hub/domain"
	"ark/util/cfg"
	"ark/util/log"
	"ark/util/metrics"
	"ark/util/str"
	"fmt"
	"strconv"

	"github.com/go-resty/resty"
	flow "github.com/trustmaster/goflow"
)

//PhoenixCOVHooker 更新Phoenix 的COV数据
type PhoenixCOVHooker struct {
	flow.Component
	In chan *domain.CoreLiteEvent
}

//OnIn Request handler only staterule =2
func (hooker *PhoenixCOVHooker) OnIn(coreLiteEvent *domain.CoreLiteEvent) {

	if cfg.Read().Phoenix.BoltSendCOV == false {
		return
	}

	metrics.AppMetrics.HubTopoPhoenixCOVCounter.Inc(1)

	birthTime := ""

	switch coreLiteEvent.EventChangeState {
	case 1:
		birthTime = str.TimeToString(coreLiteEvent.StartTime)
	case 3:
		birthTime = str.TimeToString(coreLiteEvent.EndTime)
	}

	url := fmt.Sprintf("http://%s:%s/api/covs", cfg.Read().Phoenix.PhoenixServerIP, cfg.Read().Phoenix.PhoenixServerPort)
	resp, err := resty.R().
		SetBody(map[string]interface{}{
			"corePointId":  strconv.Itoa(coreLiteEvent.CorePointID),
			"coreSourceId": strconv.Itoa(coreLiteEvent.CoreSourceID),
			"engineId":     "100",
			"qos":          "1",
			"value":        fmt.Sprintf("%f", coreLiteEvent.CurrentNumericValue),
			"state":        strconv.Itoa(coreLiteEvent.EventChangeState),
			"birthTime":    birthTime,
		}).
		Post(url)

	if err != nil {
		fmt.Printf("\nError: %v", err)
		fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
		fmt.Printf("\nResponse Status: %v", resp.Status())
		fmt.Printf("\nResponse Time: %v", resp.Time())
		fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
		fmt.Printf("\nResponse Body: %v", resp)
		log.Error(err)
	}
}
