package jobs

import (
	"context"

	"github.com/reugn/go-quartz/job"
	"github.com/reugn/go-quartz/quartz"
	"github.com/rraymondgh/arr-interfaces/internal/webhook"
)

type JobName string
type jobs map[JobName]quartz.Job

const (
	SonarrClassifierJob JobName = "SonarrClassifier"
	RadarrClassifierJob JobName = "RadarrClassifier"
)

func DefinedJobs(w webhook.Arr) jobs {
	return jobs{
		SonarrClassifierJob: job.NewFunctionJobWithDesc(func(ctx context.Context) (int, error) {
			w.Log.Infof("%s", SonarrClassifierJob)
			return 0, w.SonarrClassifier(ctx)
		}, string(SonarrClassifierJob)),
		RadarrClassifierJob: job.NewFunctionJobWithDesc(func(ctx context.Context) (int, error) {
			w.Log.Infof("%s", RadarrClassifierJob)
			return 0, w.RadarrClassifier(ctx)
		}, string(RadarrClassifierJob)),
	}
}
