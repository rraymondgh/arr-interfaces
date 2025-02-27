package quartz

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/reugn/go-quartz/logger"
	"github.com/reugn/go-quartz/quartz"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/lazy"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/worker"
	"github.com/rraymondgh/arr-interfaces/internal/quartz/config"
	"github.com/rraymondgh/arr-interfaces/internal/quartz/jobs"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/requestworker"
	"github.com/rraymondgh/arr-interfaces/internal/webhook"
	slogzap "github.com/samber/slog-zap/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In
	Config       config.Config
	TmdbAnalysis lazy.Lazy[*requestworker.TmdbAnalysis]
	Log          *zap.SugaredLogger
}

type Result struct {
	fx.Out
	Worker worker.Worker `group:"workers"`
}

func New(p Params) (Result, error) {
	var scheduler quartz.Scheduler

	return Result{
		Worker: worker.NewWorker(
			"quartz",
			fx.Hook{
				OnStart: func(ctx context.Context) error {
					a, err := p.TmdbAnalysis.Get()
					if err != nil {
						return err
					}
					w := webhook.Arr{
						Analysis: a,
						Log:      p.Log,
					}

					slogLogger := slog.New(slogzap.Option{Level: slog.LevelDebug, Logger: p.Log.Desugar()}.NewZapHandler())
					scheduler, err := quartz.NewStdScheduler(quartz.WithLogger(logger.NewSlogLogger(ctx, slogLogger)))
					if err != nil {
						return err
					}
					scheduler.Start(ctx)

					for _, schedItem := range p.Config.Schedule {
						jobFunc, ok := jobs.DefinedJobs(w)[schedItem.Name]
						if !ok {
							return fmt.Errorf("[%s] no schedule job defined", schedItem.Name)
						}
						trig, err := quartz.NewCronTrigger(schedItem.CronExpr)
						if err != nil {
							return err
						}
						err = scheduler.ScheduleJob(
							quartz.NewJobDetail(jobFunc, quartz.NewJobKey(string(schedItem.Name))),
							trig,
						)
						if err != nil {
							return err
						}
						if p.Config.TestRun {
							err = scheduler.ScheduleJob(
								quartz.NewJobDetail(jobFunc, quartz.NewJobKey(string(schedItem.Name)+"RunOnce")),
								quartz.NewRunOnceTrigger(0*time.Second),
							)
							if err != nil {
								return err
							}
						}

					}

					return nil
				},
				OnStop: func(ctx context.Context) error {
					// there's no evidence this hook is being called...
					scheduler.Stop()
					scheduler.Wait(ctx)
					return nil
				},
			},
		),
	}, nil
}
