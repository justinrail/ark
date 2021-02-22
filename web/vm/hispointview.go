package vm

//HisPointView hispoint view model
type HisPointView struct {
	CoreSourceID int
	CorePointID  int
	//当前数值值缓冲
	CurrentNumericValueString string
	//当前字符串值缓冲
	CurrentStringValue string
	//数据是否有效,0:ok,1:low limit exceed 2: up limit exceed
	LimitState      int
	TimestampString string
	StandardID      int
}
