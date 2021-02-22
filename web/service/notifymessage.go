package service

import (
	"ark/hub/domain"
	"ark/web/vm"
	"fmt"
	"reflect"
)

//GetNotifyMessage get history complex indexs
func GetNotifyMessage(queueID string) []vm.LiteItem {
	items := make([]vm.LiteItem, 0)

	que := domain.MessageSinks[queueID]
	if que != nil {
		que.ForEach(func(msg interface{}) {
			item := vm.LiteItem{}
			item.ItemName = reflect.TypeOf(msg).String()

			// out, err := json.Marshal(msg)
			// if err == nil {
			// 	item.ItemValue = string(out)
			// }
			item.ItemValue = fmt.Sprint(msg)

			items = append(items, item)
		})
	}

	return items
}
