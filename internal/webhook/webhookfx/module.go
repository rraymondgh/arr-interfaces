package webhookfx

import (
	"github.com/rraymondgh/arr-interfaces/internal/webhook/httpserver"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"webhook",
		fx.Provide(
			httpserver.New,
		),
	)
}
