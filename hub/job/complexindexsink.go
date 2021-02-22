package job

import (
	"ark/hub/domain"
	"ark/hub/sink"
	"ark/store/entity"
	"ark/store/influxstore"
	"container/list"
)

const (
	//HisComplexIndexSaveBatchSize 历史指标每次存储的数据量大小，用于减少数据库服务器调用次数，减少开销
	HisComplexIndexSaveBatchSize = 10
	//HisComplexIndexSaveMaxBatchCount 每个Job内历史指标最多存储次数，避免超大数据导致任务超时执行和数据存储的平稳（数据上送过快，依然会丢）
	HisComplexIndexSaveMaxBatchCount = 5
)

//HisComplexIndexSinkJob 历史数据存储任务
type HisComplexIndexSinkJob struct {
	Spec          string
	ComplexIndexs []*entity.ComplexIndex
}

//Run cron调度的func
func (job *HisComplexIndexSinkJob) Run() {
	storeItems := list.New()
	for j := 0; j < HisComplexIndexSaveMaxBatchCount; j++ {
		//如果数据一次没取完，继续取
		if sink.HisComplexIndexSinkQueue.Count() > 0 {
			//如果未满一批，则先取一批
			for i := 0; i < HisComplexIndexSaveBatchSize; i++ {
				hp := sink.HisComplexIndexSinkQueue.Dequeue()
				if hp != nil {
					p := hp.(*domain.HisComplexIndex)
					storeItems.PushBack(p)
				} else {
					break
				}
			}
			influxstore.AppendHisComplexIndexs(storeItems)
			storeItems = list.New()
		} else {
			break
		}
	}
}
