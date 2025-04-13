package config

import (
	"github.com/robfig/cron/v3"
	"learn/services"
)

func InitCronJobs() *cron.Cron {
	cronScheduler := cron.New()

	cronScheduler.AddFunc("*/1 * * * *", func() {
		services.RunScheduler()
	})

	cronScheduler.Start()

	return cronScheduler
}
