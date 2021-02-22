package domain

//HisCorePoint 历史数据
type HisCorePoint struct {
	CoreSourceID int
	CorePointID  int
	//当前数值值缓冲
	CurrentNumericValue float32
	//当前字符串值缓冲
	CurrentStringValue string
	LimitState         int
	Timstamp           int64
	StandardID         int
}
