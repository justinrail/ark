package service

import (
	"ark/store/influxstore"
	"ark/util/str"
	"ark/web/vm"
	"fmt"

	"gopkg.in/jeevatkm/go-model.v1"
)

//GetHisPoint get all history point
func GetHisPoint() []vm.HisPointView {
	items := make([]vm.HisPointView, 0)

	ps := influxstore.GetAllHisCorePoint()

	for _, p := range ps {
		vp := vm.HisPointView{}
		model.Copy(&vp, p)
		vp.CurrentNumericValueString = fmt.Sprintf("%.2f", p.CurrentNumericValue)
		vp.TimestampString = str.TimeToString(p.Timstamp)
		items = append(items, vp)
	}
	return items
}
