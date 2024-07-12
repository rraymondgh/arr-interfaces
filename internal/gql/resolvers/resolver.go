package resolvers

import (
	"github.com/rraymondgh/arr-interface/internal/gql"
	"github.com/rraymondgh/arr-interface/internal/servarr"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	log           *zap.SugaredLogger
	servarrConfig servarr.Config
}

func New(
	log *zap.SugaredLogger,
	config servarr.Config,
) gql.ResolverRoot {
	return &Resolver{
		log:           log,
		servarrConfig: config,
	}
}
