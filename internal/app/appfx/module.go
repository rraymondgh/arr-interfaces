package appfx

import (
	"github.com/rraymondgh/arr-interface/internal/boilerplate/app/boilerplateappfx"
	"github.com/rraymondgh/arr-interface/internal/boilerplate/httpserver/httpserverfx"
	"github.com/rraymondgh/arr-interface/internal/gql/gqlfx"
	"github.com/rraymondgh/arr-interface/internal/version/versionfx"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"app",
		boilerplateappfx.New(),
		gqlfx.New(),
		httpserverfx.New(),
		versionfx.New(),
	)
}
