package gqlfx

import (
	"github.com/rraymondgh/arr-interface/internal/gql/config"
	"github.com/rraymondgh/arr-interface/internal/gql/httpserver"
	"github.com/rraymondgh/arr-interface/internal/gql/resolvers"

	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"graphql",
		fx.Provide(
			config.New,
			httpserver.New,
			resolvers.New,
		),
	)
}
