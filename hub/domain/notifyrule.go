package domain

import (
	"ark/util/slide"
	"container/list"
)

const (
	//QueueLengthForMessage 有限长消息队列的长度
	QueueLengthForMessage = 1024
)

//NotifyRule 告警通知规则, 没有receiver，一个queue怎么用自己决定
type NotifyRule struct {
	NotifyRuleID int
	//rule的全局唯一ID
	UUID           string
	NotifyRuleName string
	//创建NotifyRule的应用实例
	Author string
	//0: pooling with dequeue= true&&size=top30 1:webhook no cache 2:webhook with queue cache
	ProvideType int
	//FilterExpression 过滤规则：coresourceid，corepointid， eventseverity， state，start，end，stateDuration，standardId
	//max eventseverity > ?
	//notifymessage <= select corepoint from activeevent where eventState in (1,2) and eventSeverity>3 and coresource in (???)
	//clean expression where state = 0 remove target from cache
	//target cache
	//notifyMessage <= select complexindex from complexindex[d:30] obj where obj.id = 232 and ci.val > 32
	//notifymessage <= select roomshadown from rooms rs where rs.id = 3 and rs.alarmgrade > 0
	//clean rule :target.val <=32
	//1 return ？返回对象加个type 查询确认后进行转换给rest api调用
	//2 时延缓冲都支持么？ci？暂时不考虑，时间
	//根据Matcher可能找到多个对象，作为多个target，如果有缓存，则有缓存的忽略
	TriggerClause string
	//清除条件，如果清除条件未达到，则保存Target，如果target存在则说明这个条件不需要重复，否则重新执行matcher
	CleanerClause string
	//重复条件，如果target存在，如果需要重复发送，那么这里存规则，给target加时间戳实现？
	//CleanAction？ clean and send some message ?
	RepeaterClause string
	MessageQueueID string
	WebHooker      string

	Trigger *NotifyTrigger

	Cleaner  *NotifyCleaner
	Repeater *NotifyTrigger

	TargetObjects *list.List
}

//Init 初始化规则
func (notifyRule *NotifyRule) Init() {
	notifyRule.TargetObjects = list.New()

	trigger, err := NewNotifyTrigger(notifyRule.TriggerClause)
	if err == nil {
		notifyRule.Trigger = trigger
	}

	cleaner, err2 := NewNotifyCleaner(notifyRule.CleanerClause)
	if err2 == nil {
		notifyRule.Cleaner = cleaner
	}
}

//Clean 清除规则执行
func (notifyRule *NotifyRule) Clean() {
	notifyRule.Cleaner.Check(notifyRule)
}

//Match 执行触发器check
func (notifyRule *NotifyRule) Match() {
	notifyRule.Trigger.Check(notifyRule)
}

//PushMessage 推送消息进入queue
func (notifyRule *NotifyRule) PushMessage(msg interface{}) {
	queue := MessageSinks[notifyRule.MessageQueueID]
	if queue == nil {
		MessageSinks[notifyRule.MessageQueueID] = slide.NewLengthWindow(QueueLengthForMessage)
	}
	MessageSinks[notifyRule.MessageQueueID].Enqueue(msg)

	// out, err := json.Marshal(msg)
	// if err == nil {
	// 	fmt.Println(string(out))
	// }
}
