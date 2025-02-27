package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/httpserver"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/lazy"

	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/config"
	t "github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/oapi"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/requestworker"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In
	Logger      *zap.SugaredLogger
	ProxyConfig config.Config
	Analysis    lazy.Lazy[*requestworker.TmdbAnalysis]
}

type Result struct {
	fx.Out
	Option httpserver.Option `group:"http_server_options"`
}

func New(p Params) Result {
	return Result{
		Option: &builder{
			logger:      p.Logger.Named("tmdbproxy"),
			proxyConfig: p.ProxyConfig,
			analysis:    p.Analysis,
		},
	}
}

const ImportIdHeader = "x-import-id"

type builder struct {
	logger      *zap.SugaredLogger
	proxyConfig config.Config
	analysis    lazy.Lazy[*requestworker.TmdbAnalysis]
}

func (builder) Key() string {
	return "tmdbproxy"
}

func (b builder) Apply(e *gin.Engine) error {
	analysis, err := b.analysis.Get()
	if err != nil {
		return err
	}
	server := tmdbproxy.TmdbAPI{
		Log:      b.logger,
		Analysis: analysis,
	}
	t.RegisterHandlersWithOptions(e, server, t.GinServerOptions{BaseURL: "/tmdb"})
	e.GET("/logs", server.Logs)
	for _, r := range e.Routes() {
		b.logger.Debug(r.Path)
	}

	return nil
}
