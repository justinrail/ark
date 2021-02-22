package repo

import (
	"ark/hub/domain"
	"ark/store/jsonstore"
	"container/list"
	"encoding/json"
)

//AppendHisComplexIndexs 批量写文件，减少文件被同时访问机率
func AppendHisComplexIndexs(hisComplexIndexs *list.List) {
	jsonstore.AppendRecords("hiscomplexindex", hisComplexIndexs)
}

//GetAllComplexIndexs 获取所有历史数据
func GetAllComplexIndexs() []domain.HisComplexIndex {
	items := make([]domain.HisComplexIndex, 0)

	strs := jsonstore.ReadAllRecords("hiscomplexindex")

	for _, str := range strs {
		hisComplexIndex := domain.HisComplexIndex{}

		err := json.Unmarshal([]byte(str), &hisComplexIndex)
		if err == nil {
			items = append(items, hisComplexIndex)
		}

	}

	return items
}
