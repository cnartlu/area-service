package tests

import (
	"github.com/cnartlu/area-service/internal/cron/job"
)

type CronJob struct {
	Github *job.Github
}

func NewCronJob(github *job.Github) *CronJob {
	return &CronJob{
		Github: github,
	}
}
