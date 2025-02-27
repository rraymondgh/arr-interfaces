package tmdbproxyfx

import (
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/config/configfx"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/config"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/httpserver"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/requestworker/requestworkerfx"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"tmdbproxy",
		configfx.NewConfigModule[config.Config]("tmdbproxy", config.NewDefaultConfig()),
		fx.Provide(
			httpserver.New,
			requestworkerfx.New,
		),
	)
}
