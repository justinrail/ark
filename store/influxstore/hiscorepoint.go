package influxstore

import (
	"ark/hub/domain"
	"ark/util/cfg"
	"container/list"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	client "github.com/influxdata/influxdb1-client"
)

//AppendHisCorePoints 批量写文件，减少文件被同时访问机率
func AppendHisCorePoints(hisPoints *list.List) {
	pts := make([]client.Point, 0)

	for e := hisPoints.Front(); e != nil; e = e.Next() {
		hp := e.Value.(*domain.HisCorePoint)
		pt := client.Point{
			Measurement: "hiscorepoints",
			Tags: map[string]string{
				"CorePointID":  fmt.Sprintf("%d", hp.CorePointID),
				"CoreSourceID": fmt.Sprintf("%d", hp.CoreSourceID),
				"StandardID":   fmt.Sprintf("%d", hp.StandardID),
			},
			Fields: map[string]interface{}{
				"CurrentNumericValue": hp.CurrentNumericValue,
				"CurrentStringValue":  hp.CurrentStringValue,
				"LimitState":          hp.LimitState,
			},
			Time:      time.Unix(hp.Timstamp, 0),
			Precision: "n",
		}

		pts = append(pts, pt)
	}

	Write(pts, cfg.Read().App.InfluxDBServerDBName, "temp_flow")

}

//GetAllHisCorePoint 获取所有历史数据
func GetAllHisCorePoint() []domain.HisCorePoint {
	items := make([]domain.HisCorePoint, 0)

	hps := Query("SELECT * FROM hiscorepoints", cfg.Read().App.InfluxDBServerDBName)

	if len(hps[0].Series) == 0 {
		return items
	}

	for _, row := range hps[0].Series[0].Values {
		hp := domain.HisCorePoint{}
		hp.Timstamp = getTimestamp(row[0])

		corePointID, err := strconv.Atoi(row[1].(string))
		if err == nil {
			hp.CorePointID = int(corePointID)
		}

		coreSourceID, err2 := strconv.Atoi(row[2].(string))
		if err2 == nil {
			hp.CoreSourceID = int(coreSourceID)
		}

		nv, err3 := (row[3].(json.Number)).Float64()
		if err3 == nil {
			hp.CurrentNumericValue = float32(nv)
		}

		hp.CurrentStringValue = row[4].(string)

		limitState, err4 := (row[5].(json.Number)).Int64()
		if err4 == nil {
			hp.LimitState = int(limitState)
		}

		standID, err7 := strconv.Atoi(row[6].(string))
		if err7 == nil {
			hp.StandardID = int(standID)
		}

		items = append(items, hp)
	}

	return items
}
