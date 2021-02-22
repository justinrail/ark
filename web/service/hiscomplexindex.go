package service

import (
	"ark/hub/domain"
	"ark/store/influxstore"
	"ark/util/str"
	"ark/web/vm"
	"fmt"

	"gopkg.in/jeevatkm/go-model.v1"
)

//GetHisComplexIndexViewByID get history complex indexs
func GetHisComplexIndexViewByID(complexIndexID int) []vm.ComplexIndexView {
	items := make([]vm.ComplexIndexView, 0)

	hisComplexIndexs := influxstore.GetAllHisComplexIndexs()

	for _, hci := range hisComplexIndexs {
		if hci.ComplexIndexID == complexIndexID {
			civ := vm.ComplexIndexView{}
			model.Copy(&civ, hci)
			civ.TimestampString = str.TimeToString(hci.Timestamp)
			civ.CurrentValueString = fmt.Sprintf("%.2f", hci.CurrentValue)
			civ.LastValueString = fmt.Sprintf("%.2f", hci.LastValue)
			civ.LastTimestampString = str.TimeToString(hci.LastTimestamp)
			items = append(items, civ)
		}
	}

	return items
}

//GetHisComplexIndexByIDs get history complex indexs
func GetHisComplexIndexByIDs(complexIndexIDs []int) []domain.HisComplexIndex {

	return influxstore.GetHisComplexIndexByIDs(complexIndexIDs)

}

//GetHisComplexIndexByIDsANDTimeRange GetHisComplexIndexByIDsANDTimeRange
func GetHisComplexIndexByIDsANDTimeRange(complexIndexIDs []int, startTime int64, endTime int64) []domain.HisComplexIndex {
	return influxstore.GetHisComplexIndexByIDsANDTimeRange(complexIndexIDs, startTime, endTime)
}
