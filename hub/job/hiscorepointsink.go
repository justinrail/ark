package job

import (
	"ark/hub/domain"
	"ark/hub/sink"
	"ark/store/entity"
	"ark/store/influxstore"
	"container/list"
)

const (
	//HisCorePointSinkBatchSize 历史数据每次存储的数据量大小，用于减少数据库服务器调用次数，减少开销
	HisCorePointSinkBatchSize = 10
	//HisCorePointSinkMaxBatchCount 每个Job内历史数据最多存储次数，避免超大数据导致任务超时执行和数据存储的平稳（数据上送过快，依然会丢）
	HisCorePointSinkMaxBatchCount = 5
)

//HisPointSinkJob 历史数据存储任务
type HisPointSinkJob struct {
	Spec       string
	CorePoints []*entity.CorePoint
}

//Run cron调度的func
func (job *HisPointSinkJob) Run() {
	storeItems := list.New()
	for j := 0; j < HisCorePointSinkMaxBatchCount; j++ {
		//如果数据一次没取完，继续取
		if sink.HisCorePointSinkQueue.Count() > 0 {
			//如果未满一批，则先取一批
			for i := 0; i < HisCorePointSinkBatchSize; i++ {
				hp := sink.HisCorePointSinkQueue.Dequeue()
				if hp != nil {
					p := hp.(*domain.HisCorePoint)
					storeItems.PushBack(p)
				} else {
					break
				}
			}
			influxstore.AppendHisCorePoints(storeItems)
			storeItems = list.New()
		} else {
			break
		}
	}
}
