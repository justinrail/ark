package domain

//CoreLiteEvent 内部采集层时间变化
type CoreLiteEvent struct {
	GatewayID     int
	CoreSourceID  int
	CorePointID   int
	CorePointName string
	EventSeverity int
	//当前数值值缓冲
	CurrentNumericValue float32
	//当前字符串值缓冲
	CurrentStringValue string
	//事件：开始1，确认2，结束3
	EventChangeState int
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

//NewCoreLiteEvent create coreliteevent from corepoint and cov
func NewCoreLiteEvent(corepoint *CorePoint) *CoreLiteEvent {
	liteEvent := &CoreLiteEvent{}
	liteEvent.GatewayID = corepoint.GatewayID
	liteEvent.CoreSourceID = corepoint.CoreSourceID
	liteEvent.CorePointID = corepoint.CorePointID
	liteEvent.CorePointName = corepoint.PointName
	liteEvent.EventSeverity = corepoint.EventSeverity
	liteEvent.CurrentNumericValue = corepoint.CurrentNumericValue
	liteEvent.CurrentStringValue = corepoint.CurrentStringValue
	liteEvent.StartTime = corepoint.StartTime
	liteEvent.EndTime = corepoint.EndTime
	liteEvent.IsAvailabe = corepoint.IsAvailabe
	liteEvent.LimitState = corepoint.LimitState
	liteEvent.StateRuleID = corepoint.StateRuleID
	liteEvent.StandardID = corepoint.StandardID
	return liteEvent
}
