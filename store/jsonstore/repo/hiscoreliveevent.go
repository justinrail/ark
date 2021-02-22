package repo

import (
	"ark/hub/domain"
	"ark/store/jsonstore"
	"ark/util/exp"
	"encoding/json"
)

//AppendHisCoreLiveEvent 数据库追加历史告警
func AppendHisCoreLiveEvent(liveEvent *domain.HisCoreLiveEvent) bool {
	return jsonstore.AppendRecord("hiscoreliveevent", liveEvent)
}

//GetAllHisCoreLiveEvent 获取所有历史告警
func GetAllHisCoreLiveEvent() []domain.HisCoreLiveEvent {
	items := make([]domain.HisCoreLiveEvent, 0)

	strs := jsonstore.ReadAllRecords("hiscoreliveevent")

	for _, str := range strs {
		hisEvent := domain.HisCoreLiveEvent{}
		//fmt.Println(str)
		err := json.Unmarshal([]byte(str), &hisEvent)
		if err == nil {
			items = append(items, hisEvent)
		}
		exp.CheckError(err)
	}

	return items
}
