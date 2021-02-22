package job

import (
	"ark/hub/domain"
	"ark/store/entity"
	"ark/store/mysql/repo"
	"time"

	"gopkg.in/jeevatkm/go-model.v1"
)

//ComplexIndexCalcJob 指标计算任务
type ComplexIndexCalcJob struct {
	Spec          string
	ComplexIndexs []*entity.ComplexIndex
}

//Run 加载配置
func (job *ComplexIndexCalcJob) Run() {
	timeNow := time.Now().Unix()

	for _, ci := range job.ComplexIndexs {
		dci, ok := domain.ComplexIndexs.Get(ci.ComplexIndexID)

		var domainCI *domain.ComplexIndex

		if !ok {
			domainCI = createComplexIndexInCache(ci)
		} else {
			domainCI = dci.(*domain.ComplexIndex)
		}

		domainCI.Calculate(timeNow)
		domainCI.Timestamp = timeNow
	}
}

//初始化complexindex在全局hashmap中
func createComplexIndexInCache(ci *entity.ComplexIndex) *domain.ComplexIndex {
	dci := domain.ComplexIndex{}
	model.Copy(&dci, ci)
	dci.Init()
	domain.ComplexIndexs.Insert(dci.ComplexIndexID, &dci)
	return &dci
}

func prepareComplexIndexCalcJob() []*TaskItem {
	complexIndexCalcJobs := make(map[string][]*entity.ComplexIndex)

	cis := repo.GetAllComplexIndex()

	for _, ci := range cis {
		if complexIndexCalcJobs[ci.CalcCron] == nil {
			complexIndexCalcJobs[ci.CalcCron] = make([]*entity.ComplexIndex, 0)
		}
		ps := complexIndexCalcJobs[ci.CalcCron]
		ps = append(ps, &ci)
		complexIndexCalcJobs[ci.CalcCron] = ps
	}

	jobs := make([]*ComplexIndexCalcJob, 0)

	for k, v := range complexIndexCalcJobs {
		calcJob := &ComplexIndexCalcJob{Spec: k, ComplexIndexs: v}
		jobs = append(jobs, calcJob)
	}

	jobItems := make([]*TaskItem, 0)

	for _, job := range jobs {
		jobItem := &TaskItem{Spec: job.Spec, Job: job}
		jobItems = append(jobItems, jobItem)
	}

	return jobItems
}
