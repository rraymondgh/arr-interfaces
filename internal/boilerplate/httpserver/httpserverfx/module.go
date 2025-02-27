package httpserverfx

import (
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/config/configfx"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/httpserver"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/httpserver/cors"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"http_server",
		configfx.NewConfigModule[httpserver.Config]("http_server", httpserver.NewDefaultConfig()),
		fx.Provide(
			httpserver.New,
			cors.New,
		),
	)
}
