package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/httpserver"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/lazy"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/requestworker"
	"github.com/rraymondgh/arr-interfaces/internal/webhook"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In
	TmdbAnalysis lazy.Lazy[*requestworker.TmdbAnalysis]
	Logger       *zap.SugaredLogger
}

type Result struct {
	fx.Out
	Option httpserver.Option `group:"http_server_options"`
}

func New(p Params) Result {
	return Result{
		Option: &builder{
			analysis: p.TmdbAnalysis,
			logger:   p.Logger,
		},
	}
}

const ImportIdHeader = "x-import-id"

type builder struct {
	analysis lazy.Lazy[*requestworker.TmdbAnalysis]
	logger   *zap.SugaredLogger
}

func (builder) Key() string {
	return "webhook"
}

func (b builder) Apply(e *gin.Engine) error {
	a, err := b.analysis.Get()
	if err != nil {
		return err
	}
	e.POST("/webhook", webhook.Arr{Analysis: a, Log: b.logger}.Event)

	return nil
}
