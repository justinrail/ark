package job

import "github.com/robfig/cron"

//TaskItem item of job
type TaskItem struct {
	Spec string
	Job  cron.Job
}
