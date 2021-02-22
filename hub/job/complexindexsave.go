package job

import (
	"ark/hub/domain"
	"ark/hub/sink"
	"ark/store/entity"
	"ark/store/mysql/repo"
	"time"

	"gopkg.in/jeevatkm/go-model.v1"
)

//HisComplexIndexSaveJob 历史指标数据存储任务
type HisComplexIndexSaveJob struct {
	Spec          string
	ComplexIndexs []*entity.ComplexIndex
}

//Run 加载配置
func (job *HisComplexIndexSaveJob) Run() {
	timeNow := time.Now().Unix()

	hisComplexIndexs := make([]*domain.HisComplexIndex, 0)

	for _, cp := range job.ComplexIndexs {
		hisIndex := domain.HisComplexIndex{}
		model.Copy(&hisIndex, cp)
		hisIndex.Timestamp = timeNow
		indexCache, ok := domain.ComplexIndexs.Get(hisIndex.ComplexIndexID)
		if ok {
			ci := indexCache.(*domain.ComplexIndex)
			if ci.IsValid {
				model.Copy(&hisIndex, ci)
				hisComplexIndexs = append(hisComplexIndexs, &hisIndex)
			}
		}
	}

	job.saveIndexs(hisComplexIndexs)
}

func (job *HisComplexIndexSaveJob) saveIndexs(hisComplexIndexs []*domain.HisComplexIndex) {
	for _, hp := range hisComplexIndexs {
		sink.HisComplexIndexSinkQueue.Enqueue(hp)
	}
}

func prepareHisComplexIndexSaveJob() []*TaskItem {
	idxSaveJobs := make(map[string][]*entity.ComplexIndex)

	complexIndexs := repo.GetAllComplexIndex()

	for _, ci := range complexIndexs {
		if len(ci.SaveCron) > 0 {
			if idxSaveJobs[ci.SaveCron] == nil {
				idxSaveJobs[ci.SaveCron] = make([]*entity.ComplexIndex, 0)
			}

			ps := idxSaveJobs[ci.SaveCron]

			ps = append(ps, &ci)
			idxSaveJobs[ci.SaveCron] = ps
		}
	}

	jobs := make([]*HisComplexIndexSaveJob, 0)

	for k, v := range idxSaveJobs {
		hisIndexJob := &HisComplexIndexSaveJob{Spec: k, ComplexIndexs: v}
		jobs = append(jobs, hisIndexJob)
	}

	jobItems := make([]*TaskItem, 0)

	for _, job := range jobs {
		jobItem := &TaskItem{Spec: job.Spec, Job: job}
		jobItems = append(jobItems, jobItem)
	}

	return jobItems
}
