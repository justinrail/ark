package notifier

import (
	"ark/hub/domain"
	"container/list"
	"time"
)

var notifyRules *list.List

func init() {
	notifyRules = list.New()
}

//Ready 准备
func Ready() {
	notifyRules = list.New()

	//load notifyrule
	rule := &domain.NotifyRule{NotifyRuleID: -1, NotifyRuleName: "mock", UUID: "test", Author: "tester", ProvideType: 0,
		TriggerClause:  "CoreLiteEvent le le.CorePointID=cp.CorePointID,le.EventChangeState=1,le.StartTime=cp.StartTime <= select CorePoint cp cp.CorePointID=cop.CorePointID from CoreLiveEvent cle where cle.CorePointID IN (1,2,3,4,26,29,32,51,52,53,54,35) && duration(cle.StartTime,now()) > 1 && cle.CurrentEventState == 1",
		CleanerClause:  "CoreLiteEvent le le.EventChangeState=3,le.EndTime=target.EndTime <= target.CurrentEventState == 0",
		RepeaterClause: "", MessageQueueID: "ms1", WebHooker: ""}

	// rule := &domain.NotifyRule{NotifyRuleID: -1, UUID: "test", Author: "tester", ProvideType: 0,
	// TriggerClause:  "domain.CoreLiteEvent le {le.EventChangeState=1,le.StartTime=cp.StartTime} <= select CorePoint cp from CoreLiveEvent cle where cle.CorePointID IN (1,2,51,35) && cle.CurrentEventState == 1 && durationSeconds(cle.StartTime) > 3",
	// CleanerClause:  "domain.CoreLiteEvent le {le.EventChangeState=3,le.EndTime=now()} <= target.CurrentEventState == 0",
	// RepeaterClause: "", MessageQueueID: "ms1", WebHooker: ""}

	notifyRules.PushBack(rule)

	//load rules to run
	for iter := notifyRules.Back(); iter != nil; iter = iter.Prev() {
		iter.Value.(*domain.NotifyRule).Init()
	}
}

//Start 启动定时计算工作
func Start() {

	for {
		time.Sleep(2 * time.Second)
		clean()
		match()
	}
}

func clean() {
	for iter := notifyRules.Back(); iter != nil; iter = iter.Prev() {
		iter.Value.(*domain.NotifyRule).Clean()
	}
}

func match() {
	for iter := notifyRules.Back(); iter != nil; iter = iter.Prev() {
		iter.Value.(*domain.NotifyRule).Match()
	}
}

//GetAllNotifyRule 获取所有的notifyrule
func GetAllNotifyRule() []*domain.NotifyRule {
	items := make([]*domain.NotifyRule, 0)
	for iter := notifyRules.Back(); iter != nil; iter = iter.Prev() {
		items = append(items, iter.Value.(*domain.NotifyRule))
	}
	return items
}
