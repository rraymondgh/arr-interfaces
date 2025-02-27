package healthcheckfx

import (
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/healthcheck"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/healthcheck/httpserver"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"healthcheck",
		fx.Provide(healthcheck.New),
		fx.Provide(httpserver.New),
	)
}
