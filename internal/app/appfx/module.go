package appfx

import (
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/app/boilerplateappfx"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/httpserver/httpserverfx"
	"github.com/rraymondgh/arr-interfaces/internal/database/databasefx"
	"github.com/rraymondgh/arr-interfaces/internal/quartz/quartzfx"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/tmdbproxyfx"
	"github.com/rraymondgh/arr-interfaces/internal/version/versionfx"
	"github.com/rraymondgh/arr-interfaces/internal/webhook/webhookfx"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"app",
		boilerplateappfx.New(),
		httpserverfx.New(),
		versionfx.New(),
		databasefx.New(),
		tmdbproxyfx.New(),
		webhookfx.New(),
		quartzfx.New(),
	)
}
