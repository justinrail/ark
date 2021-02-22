package influxstore

import (
	"ark/hub/domain"
	"ark/util/cfg"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	client "github.com/influxdata/influxdb1-client"
)

//AppendHisCoreLiveEvent 保存历史告警
func AppendHisCoreLiveEvent(hisCoreLiveEvent *domain.HisCoreLiveEvent) {
	pts := make([]client.Point, 0)

	pt := client.Point{
		Measurement: "hiscoreliveevents",
		Tags: map[string]string{
			"GatewayID":     fmt.Sprintf("%d", hisCoreLiveEvent.GatewayID),
			"CoreSourceID":  fmt.Sprintf("%d", hisCoreLiveEvent.CoreSourceID),
			"CorePointID":   fmt.Sprintf("%d", hisCoreLiveEvent.CorePointID),
			"CorePointName": hisCoreLiveEvent.CorePointName,
			"EventSeverity": fmt.Sprintf("%d", hisCoreLiveEvent.EventSeverity),
			"StandardID":    fmt.Sprintf("%d", hisCoreLiveEvent.StandardID),
		},
		Fields: map[string]interface{}{
			"CurrentNumericValue": hisCoreLiveEvent.CurrentNumericValue,
			"CurrentStringValue":  hisCoreLiveEvent.CurrentStringValue,
			"EndTime":             hisCoreLiveEvent.EndTime,
		},
		Time:      time.Unix(hisCoreLiveEvent.StartTime, 0),
		Precision: "n",
	}

	pts = append(pts, pt)

	Write(pts, cfg.Read().App.InfluxDBServerDBName, "temp_flow")
}

//GetAllHisCoreLiveEvents 获取所有历史告警
func GetAllHisCoreLiveEvents() []domain.HisCoreLiveEvent {
	items := make([]domain.HisCoreLiveEvent, 0)

	hes := Query("SELECT * FROM hiscoreliveevents", cfg.Read().App.InfluxDBServerDBName)

	if len(hes[0].Series) == 0 {
		return items
	}

	for _, row := range hes[0].Series[0].Values {
		he := domain.HisCoreLiveEvent{}
		he.StartTime = getTimestamp(row[0])

		corePointID, err := strconv.Atoi(row[1].(string))
		if err == nil {
			he.CorePointID = int(corePointID)
		}

		he.CorePointName = row[2].(string)

		coreSourceID, err2 := strconv.Atoi(row[3].(string))
		if err2 == nil {
			he.CoreSourceID = int(coreSourceID)
		}

		nv, err3 := (row[4].(json.Number)).Float64()
		if err3 == nil {
			he.CurrentNumericValue = float32(nv)
		}

		he.CurrentStringValue = row[5].(string)

		endTime, err4 := (row[6].(json.Number)).Int64()
		if err4 == nil {
			he.EndTime = endTime
		}

		eventSeverity, err5 := strconv.Atoi(row[7].(string))
		if err5 == nil {
			he.EventSeverity = int(eventSeverity)
		}

		gatewayID, err6 := strconv.Atoi(row[8].(string))
		if err6 == nil {
			he.GatewayID = int(gatewayID)
		}

		standID, err7 := strconv.Atoi(row[9].(string))
		if err7 == nil {
			he.StandardID = int(standID)
		}

		items = append(items, he)
	}

	return items
}
