package app

import (
	"github.com/rraymondgh/arr-interfaces/internal/app/appfx"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/cli/hooks"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/logging/loggingfx"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func New() *fx.App {
	return fx.New(
		appfx.New(),
		loggingfx.WithLogger(),
		fx.Invoke(func(
			logger *zap.SugaredLogger,
			_ *cli.App,
			_ hooks.AttachedHooks,
		) {
			logger.Debug("app invoked")
		}),
	)
}
