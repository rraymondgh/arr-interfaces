package workerfx

import (
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/worker"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"worker",
		fx.Provide(worker.NewRegistry),
	)
}
