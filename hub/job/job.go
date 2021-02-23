package job

import (
	"ark/util/cfg"

	"github.com/robfig/cron"
)

var cr *cron.Cron

//Ready prepare for jobs
func Ready() {
	cr = cron.New()

	if cfg.Read().Hub.JobHisPointRunFuse {
		registerJobs(prepareHisPointEmitJob())
		cr.AddJob("@every 10s", new(HisPointSinkJob))
	}

	if cfg.Read().Hub.JobComplexIndexRunFuse {
		registerJobs(prepareComplexIndexCalcJob())
	}

	if cfg.Read().Hub.JobHisComplexIndexRunFuse {
		registerJobs(prepareHisComplexIndexSaveJob())
		cr.AddJob("@every 10s", new(HisComplexIndexSinkJob))
	}

	//cr.AddJob("@every 10s", new(HisComplexIndexSinkJob))

	// if cfg.Read().Phoenix.JobSendCOR {
	// 	cr.AddFunc("@every 5s", sendCOR2Phoenix)
	// }

	//cr.AddFunc("@every 30s", freeSystemMemory) //go return mem to system default is 5min , it's for debug, disabled by default
}

//RegisterJobs 注册任务
func registerJobs(jobs []*TaskItem) {

	for _, job := range jobs {
		cr.AddJob(job.Spec, job.Job)
	}
}

//Start start jobs
func Start() {
	cr.Start()
}
