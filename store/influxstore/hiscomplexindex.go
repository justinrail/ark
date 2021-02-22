package influxstore

import (
	"ark/hub/domain"
	"ark/util/cfg"
	"ark/util/str"
	"container/list"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	client "github.com/influxdata/influxdb1-client"
)

//GetHisComplexIndexByIDsANDTimeRange GetHisComplexIndexByIDsANDTimeRange
func GetHisComplexIndexByIDsANDTimeRange(complexIndexIDs []int, startTime int64, endTime int64) []domain.HisComplexIndex {

	idexp := " ("
	for _, id := range complexIndexIDs {
		idexp = idexp + " ComplexIndexID='" + strconv.Itoa(id) + "' OR"
	}

	idexp = idexp[0:len(idexp)-2] + ")"

	timeexp := " AND time>='" + str.TimeToString(startTime) + "' AND time <='" + str.TimeToString(endTime) + "'"

	sql := "SELECT * FROM hiscomplexindexs WHERE" + idexp + timeexp + "  tz('Local')"
	cis := Query(sql, cfg.Read().App.InfluxDBServerDBName)

	return getHisComplexIndexFromResult(cis)
}

func getHisComplexIndexFromResult(cis []client.Result) []domain.HisComplexIndex {
	items := make([]domain.HisComplexIndex, 0)

	if len(cis[0].Series) == 0 {
		return items
	}

	for _, row := range cis[0].Series[0].Values {
		hci := domain.HisComplexIndex{}
		hci.Timestamp = getTimestamp(row[0])

		hci.BusinessID = row[1].(string)

		calcType, err2 := (row[2].(json.Number)).Int64()
		if err2 == nil {
			hci.CalcType = int(calcType)
		}

		hci.Category = row[3].(string)

		ciID, err7 := strconv.Atoi(row[4].(string))
		if err7 == nil {
			hci.ComplexIndexID = int(ciID)
		}

		hci.ComplexIndexName = row[5].(string)

		curValue, err3 := (row[6].(json.Number)).Float64()

		if err3 == nil {
			hci.CurrentValue = curValue
		}

		grID, err := strconv.Atoi(row[7].(string))
		if err == nil {
			hci.GlobalResourceID = grID
		}
		hci.Label = row[8].(string)

		lastTS, err4 := (row[9].(json.Number)).Int64()
		if err4 == nil {
			hci.LastTimestamp = lastTS
		}

		lastVal, err5 := row[10].(json.Number).Float64()
		if err5 == nil {
			hci.LastValue = lastVal
		}

		objID, err9 := strconv.Atoi(row[11].(string))
		if err9 == nil {
			hci.ObjectTypeID = int(objID)
		}

		hci.Title = row[12].(string)

		items = append(items, hci)
	}

	return items
}

//GetHisComplexIndexByIDs GetHisComplexIndexByIDs
func GetHisComplexIndexByIDs(complexIndexIDs []int) []domain.HisComplexIndex {

	idexp := " ("
	for _, id := range complexIndexIDs {
		idexp = idexp + " ComplexIndexID='" + strconv.Itoa(id) + "' OR"
	}

	idexp = idexp[0:len(idexp)-2] + ")"

	cis := Query("SELECT * FROM hiscomplexindexs WHERE"+idexp, cfg.Read().App.InfluxDBServerDBName)
	return getHisComplexIndexFromResult(cis)

}

//AppendHisComplexIndexs 批量写
func AppendHisComplexIndexs(hisComplexIndexs *list.List) {
	pts := make([]client.Point, 0)

	for e := hisComplexIndexs.Front(); e != nil; e = e.Next() {
		ci := e.Value.(*domain.HisComplexIndex)
		pt := client.Point{
			Measurement: "hiscomplexindexs",
			Tags: map[string]string{
				"GlobalResourceID": fmt.Sprintf("%d", ci.GlobalResourceID),
				"ComplexIndexID":   fmt.Sprintf("%d", ci.ComplexIndexID),
				"ComplexIndexName": ci.ComplexIndexName,
				"Category":         ci.Category,
				"Label":            ci.Label,
				"Title":            ci.Title,
				"ObjectTypeID":     fmt.Sprintf("%d", ci.ObjectTypeID),
				"BusinessID":       ci.BusinessID,
			},
			Fields: map[string]interface{}{
				"CurrentValue":  ci.CurrentValue,
				"LastValue":     ci.LastValue,
				"LastTimestamp": ci.LastTimestamp,
				"CalcType":      ci.CalcType,
			},
			Time:      time.Now(),
			Precision: "n",
		}
		//fmt.Println(str.TimeToString(time.Unix(ci.Timestamp, 0).Unix()))
		//fmt.Println(str.NowString())
		pts = append(pts, pt)
	}

	Write(pts, cfg.Read().App.InfluxDBServerDBName, "temp_flow")
}

//GetAllHisComplexIndexs 获取所有历史数据
func GetAllHisComplexIndexs() []domain.HisComplexIndex {
	cis := Query("SELECT * FROM hiscomplexindexs", cfg.Read().App.InfluxDBServerDBName)
	return getHisComplexIndexFromResult(cis)
}

func getTimestamp(obj interface{}) int64 {
	t, err := time.Parse(time.RFC3339, obj.(string))
	if err != nil {
		return 0
	}

	return t.Unix()
}
