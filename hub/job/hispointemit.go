package job

import (
	"ark/hub/domain"
	"ark/hub/sink"
	"ark/store/entity"
	"ark/store/mysql/repo"
	"time"

	"gopkg.in/jeevatkm/go-model.v1"
)

//HisPointEmitJob 历史数据任务
type HisPointEmitJob struct {
	Spec       string
	CorePoints []*entity.CorePoint
}

//Run 加载配置
func (job *HisPointEmitJob) Run() {
	timeNow := time.Now().Unix()

	hisPoints := make([]*domain.HisCorePoint, 0)

	for _, cp := range job.CorePoints {
		hisPoint := domain.HisCorePoint{}
		model.Copy(&hisPoint, cp)
		hisPoint.StandardID = cp.StandardID
		hisPoint.Timstamp = timeNow
		pointCache, ok := domain.CorePoints.Get(hisPoint.CorePointID)
		if ok {
			p := pointCache.(*domain.CorePoint)
			if p.DataIsReady() {
				hisPoint.CurrentNumericValue = p.CurrentNumericValue
				hisPoint.CurrentStringValue = p.CurrentStringValue
				hisPoint.LimitState = p.LimitState

				hisPoints = append(hisPoints, &hisPoint)
			}
		}
	}

	job.savePoints(hisPoints)
}

func (job *HisPointEmitJob) savePoints(hisPoints []*domain.HisCorePoint) {
	for _, hp := range hisPoints {
		sink.HisCorePointSinkQueue.Enqueue(hp)
	}

}

func prepareHisPointEmitJob() []*TaskItem {
	pointJobs := make(map[string][]*entity.CorePoint)

	corePoints := repo.GetCorePointHasCron()

	for _, cp := range corePoints {
		if pointJobs[cp.Cron] == nil {
			pointJobs[cp.Cron] = make([]*entity.CorePoint, 0)
		}
		ps := pointJobs[cp.Cron]
		ps = append(ps, &cp)
		pointJobs[cp.Cron] = ps
	}

	jobs := make([]*HisPointEmitJob, 0)

	for k, v := range pointJobs {
		hisPointJob := &HisPointEmitJob{Spec: k, CorePoints: v}
		jobs = append(jobs, hisPointJob)
	}

	jobItems := make([]*TaskItem, 0)

	for _, job := range jobs {
		jobItem := &TaskItem{Spec: job.Spec, Job: job}
		jobItems = append(jobItems, jobItem)
	}

	return jobItems
}
