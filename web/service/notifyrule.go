package service

import (
	"ark/hub/domain"
	"ark/hub/topology/notifier"
)

//GetAllNotifyRule get all complex indexs
func GetAllNotifyRule() []*domain.NotifyRule {
	return notifier.GetAllNotifyRule()
}
