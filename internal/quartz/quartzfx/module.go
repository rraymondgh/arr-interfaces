package quartzfx

import (
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/config/configfx"
	"github.com/rraymondgh/arr-interfaces/internal/quartz"
	"github.com/rraymondgh/arr-interfaces/internal/quartz/config"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"quartz",
		configfx.NewConfigModule[config.Config]("quartz", config.NewDefaultConfig()),
		fx.Provide(
			quartz.New,
		),
	)
}
