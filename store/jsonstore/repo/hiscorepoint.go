package repo

import (
	"ark/hub/domain"
	"ark/store/jsonstore"
	"container/list"
	"encoding/json"
)

//AppendHisCorePoint 数据库追加历史告警
func AppendHisCorePoint(hisPoint *domain.HisCorePoint) bool {
	return jsonstore.AppendRecord("hiscorepoint", hisPoint)
}

//AppendHisCorePoints 批量写文件，减少文件被同时访问机率
func AppendHisCorePoints(hisPoints *list.List) {
	jsonstore.AppendRecords("hiscorepoint", hisPoints)
}

//GetAllHisCorePoint 获取所有历史数据
func GetAllHisCorePoint() []domain.HisCorePoint {
	items := make([]domain.HisCorePoint, 0)

	strs := jsonstore.ReadAllRecords("hiscorepoint")

	for _, str := range strs {
		hisPoint := domain.HisCorePoint{}

		err := json.Unmarshal([]byte(str), &hisPoint)
		if err == nil {
			items = append(items, hisPoint)
		}

	}

	return items
}
