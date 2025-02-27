package clifx

import (
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/cli"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/cli/args"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"cli",
		fx.Provide(args.New),
		fx.Provide(cli.New),
	)
}
