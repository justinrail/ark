package repo

import (
	"ark/hub/domain"
	"ark/store/jsonstore"
	"ark/util/exp"
	"encoding/json"
)

//GetAllHisComplexIndexs 获取所有历史指标
func GetAllHisComplexIndexs() []*domain.HisComplexIndex {
	items := make([]*domain.HisComplexIndex, 0)

	strs := jsonstore.ReadAllRecords("hiscomplexindex")

	for _, str := range strs {
		hisIndex := domain.HisComplexIndex{}
		err := json.Unmarshal([]byte(str), &hisIndex)
		if err == nil {
			items = append(items, &hisIndex)
		}
		exp.CheckError(err)
	}

	return items
}
