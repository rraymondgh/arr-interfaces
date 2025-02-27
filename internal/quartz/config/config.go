package config

import "github.com/rraymondgh/arr-interfaces/internal/quartz/jobs"

type schedItem struct {
	Name     jobs.JobName
	CronExpr string
}

type Config struct {
	TestRun  bool
	Schedule []schedItem
}

func NewDefaultConfig() Config {
	return Config{
		TestRun: true,
		Schedule: []schedItem{
			{Name: jobs.SonarrClassifierJob, CronExpr: "0 0 3 * * *"},
			{Name: jobs.RadarrClassifierJob, CronExpr: "0 0 3 1,4 * *"},
		},
	}
}
