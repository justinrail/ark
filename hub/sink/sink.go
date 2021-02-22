package sink

import "ark/util/slide"

const (
	//SinkQueueSize 默认的定长窗口长度
	SinkQueueSize = 1024
)

//HisCorePointSinkQueue 历史数据暂存定长窗口
var HisCorePointSinkQueue *slide.LengthWindow

//HisComplexIndexSinkQueue 历史指标暂存定长窗口
var HisComplexIndexSinkQueue *slide.LengthWindow

func init() {
	HisCorePointSinkQueue = slide.NewLengthWindow(SinkQueueSize)
	HisComplexIndexSinkQueue = slide.NewLengthWindow(SinkQueueSize)
}
