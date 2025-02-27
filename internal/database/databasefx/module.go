package databasefx

import (
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/config/configfx"
	"github.com/rraymondgh/arr-interfaces/internal/database/meiliclient"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"database",
		configfx.NewConfigModule[meiliclient.Config]("meiliclient", meiliclient.NewDefaultConfig()),
		fx.Provide(
			meiliclient.New,
		),
	)
}
