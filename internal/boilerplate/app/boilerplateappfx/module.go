package boilerplateappfx

import (
	configcmd "github.com/rraymondgh/arr-interface/internal/boilerplate/app/cmd/config"
	workercmd "github.com/rraymondgh/arr-interface/internal/boilerplate/app/cmd/worker"
	"github.com/rraymondgh/arr-interface/internal/boilerplate/boilerplatefx"
	"github.com/rraymondgh/arr-interface/internal/boilerplate/cli/hooks"
	"github.com/rraymondgh/arr-interface/internal/boilerplate/worker/workerfx"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"app_boilerplate",
		boilerplatefx.New(),
		workerfx.New(),
		fx.Provide(
			hooks.New,
			configcmd.New,
			workercmd.New,
		),
	)
}
