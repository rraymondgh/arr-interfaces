package servarrfx

import (
	"github.com/rraymondgh/arr-interface/internal/boilerplate/config/configfx"
	"github.com/rraymondgh/arr-interface/internal/servarr"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"servarr",
		configfx.NewConfigModule[servarr.Config]("servarr", servarr.NewDefaultConfig()),
	)
}
