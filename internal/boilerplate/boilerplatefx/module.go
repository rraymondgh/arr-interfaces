package boilerplatefx

import (
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/cli/clifx"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/config/configfx"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/healthcheck/healthcheckfx"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/logging/loggingfx"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/validation/validationfx"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"boilerplate",
		clifx.New(),
		configfx.New(),
		healthcheckfx.New(),
		loggingfx.New(),
		validationfx.New(),
	)
}
