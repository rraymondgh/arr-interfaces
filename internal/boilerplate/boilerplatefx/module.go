package boilerplatefx

import (
	"github.com/rraymondgh/arr-interface/internal/boilerplate/cli/clifx"
	"github.com/rraymondgh/arr-interface/internal/boilerplate/config/configfx"
	"github.com/rraymondgh/arr-interface/internal/boilerplate/healthcheck/healthcheckfx"
	"github.com/rraymondgh/arr-interface/internal/boilerplate/logging/loggingfx"
	"github.com/rraymondgh/arr-interface/internal/boilerplate/validation/validationfx"
	"github.com/rraymondgh/arr-interface/internal/servarr/servarrfx"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"boilerplate",
		clifx.New(),
		servarrfx.New(),
		configfx.New(),
		healthcheckfx.New(),
		loggingfx.New(),
		validationfx.New(),
	)
}
