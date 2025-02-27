package requestworker

import (
	"github.com/meilisearch/meilisearch-go"
	"github.com/munnik/uniqueue"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/config"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/searchmodel"
	"go.uber.org/zap"
)

type TmdbAnalysis struct {
	Config          config.Config
	Log             *zap.SugaredLogger
	MeiliClient     meilisearch.ServiceManager
	TmdbSearchCache *searchmodel.TmdbSearchCache
	UrlLogCache     *searchmodel.UrlLogCache
	UniqueQueue     *uniqueue.UQ[searchmodel.UrlLog]
}
