package config

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/rraymondgh/arr-interface/internal/gql"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	ResolverRoot gql.ResolverRoot
}

func New(p Params) graphql.ExecutableSchema {

	return gql.NewExecutableSchema(gql.Config{Resolvers: p.ResolverRoot})

}
