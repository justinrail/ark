package domain

import (
	"gopkg.in/jeevatkm/go-model.v1"
)

//HisCoreLiveEvent 历史告警
type HisCoreLiveEvent struct {
	GatewayID     int
	CoreSourceID  int
	CorePointID   int
	CorePointName string
	EventSeverity int
	//当前数值值缓冲
	CurrentNumericValue float32
	//当前字符串值缓冲
	CurrentStringValue string
	//事件：开始，确认，结束
	EventState int
	//数据是否有效
	IsAvailabe bool
	//数据是否有效,0:ok,1:low limit exceed 2: up limit exceed
	LimitState int
	//常用的State开始对应的，开始时间，如果CurrentEventState可用，否则存的是上次的结束时间
	StartTime int64
	//常用的State结束对应的，结束时间，如果CurrentEndState可用，否则存的是上次的结束时间
	EndTime     int64
	StateRuleID int
	StandardID  int
}

//NewHisCoreLiveEvent create his event by coreliveevent
func NewHisCoreLiveEvent(liveEvent *CoreLiveEvent) *HisCoreLiveEvent {
	hisLiveEvent := HisCoreLiveEvent{}
	model.Copy(&hisLiveEvent, liveEvent)
	hisLiveEvent.EventState = liveEvent.CurrentEventState

	return &hisLiveEvent
}
