package vm

//LiveEventView liveevent view model
type LiveEventView struct {
	GatewayID     int
	CoreSourceID  int
	CorePointID   int
	CorePointName string
	//当前数值值缓冲
	CurrentNumericValueString string
	//当前字符串值缓冲
	CurrentStringValue string
	EventSeverity      int
	//事件：开始，确认，结束
	CurrentEventState int
	//数据是否有效
	IsAvailabe bool
	//数据是否有效,0:ok,1:low limit exceed 2: up limit exceed
	LimitState int
	//常用的State开始对应的，开始时间，如果CurrentEventState可用，否则存的是上次的结束时间
	StartTimeString string
	//常用的State结束对应的，结束时间，如果CurrentEndState可用，否则存的是上次的结束时间
	EndTimeString string
	StateRuleID   int
	StandardID    int
}
