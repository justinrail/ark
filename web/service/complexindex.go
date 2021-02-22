package service

import (
	"ark/hub/domain"
	"ark/util/str"
	"ark/web/vm"
	"fmt"

	"gopkg.in/jeevatkm/go-model.v1"
)

//GetAllComplexIndexView get all complex indexs
func GetAllComplexIndexView() []vm.ComplexIndexView {
	items := make([]vm.ComplexIndexView, 0)

	cols := domain.ComplexIndexs.Iter()
	for kv := range cols {
		ci := kv.Value.(*domain.ComplexIndex)
		civ := vm.ComplexIndexView{}
		model.Copy(&civ, ci)
		civ.TimestampString = str.TimeToString(ci.Timestamp)
		civ.CurrentValueString = fmt.Sprintf("%.2f", ci.CurrentValue)
		civ.LastValueString = fmt.Sprintf("%.2f", ci.LastValue)
		civ.LastTimestampString = str.TimeToString(ci.LastTimestamp)
		items = append(items, civ)
	}

	return items
}

//GetAllComplexIndex get all complex indexs
func GetAllComplexIndex() []domain.ComplexIndex {
	items := make([]domain.ComplexIndex, 0)

	cols := domain.ComplexIndexs.Iter()
	for kv := range cols {
		ci := kv.Value.(*domain.ComplexIndex)
		items = append(items, *ci)
	}

	return items
}

//GetComplexIndexByID get complex index by ID
func GetComplexIndexByID(complexIndexID int) *domain.ComplexIndex {

	cp, ok := domain.ComplexIndexs.Get(complexIndexID)
	if ok {
		return cp.(*domain.ComplexIndex)
	}

	return nil

}
