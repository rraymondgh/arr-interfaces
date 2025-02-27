package requestworkerfx

import (
	"context"
	"sync"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/meilisearch/meilisearch-go"
	"github.com/munnik/uniqueue"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/cachestruct"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/lazy"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/config"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/requestworker"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/searchmodel"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In
	Log         *zap.SugaredLogger
	Meiliclient lazy.Lazy[meilisearch.ServiceManager]
	Config      config.Config
}

type Result struct {
	fx.Out
	TmdbAnalysis lazy.Lazy[*requestworker.TmdbAnalysis]
	AppHook      fx.Hook `group:"app_hooks"`
}

func New(p Params) (Result, error) {
	stopped := make(chan struct{})
	var closewg sync.WaitGroup
	var server *requestworker.TmdbAnalysis
	lazyAnalysis := lazy.New(func() (*requestworker.TmdbAnalysis, error) {
		tmdbSearchCache := cachestruct.New[string, any](p.Config.CompressCache)
		urlLogCache := cachestruct.New[uint64, searchmodel.UrlLog](
			p.Config.CompressCache,
			ttlcache.WithTTL[uint64, []byte](time.Duration(p.Config.LogCacheTTL)*time.Minute),
		)
		mc, err := p.Meiliclient.Get()
		if err != nil {
			return nil, err
		}
		server = &requestworker.TmdbAnalysis{
			Config:          p.Config,
			Log:             p.Log,
			MeiliClient:     mc,
			TmdbSearchCache: tmdbSearchCache,
			UrlLogCache:     urlLogCache,
			UniqueQueue:     uniqueue.NewUQ[searchmodel.UrlLog](2),
		}
		server.UniqueQueue.AutoRemoveConstraint = true

		go func() {
			<-stopped
			p.Log.Debug("shutting down analysis")
			close(server.UniqueQueue.Back())
			server.UrlLogCache.Stop()
			closewg.Done()
		}()

		ctx := context.Background()
		var initialStartup sync.WaitGroup
		initialStartup.Add(2)
		go server.InitialDataAndIndexes(ctx, &initialStartup)
		go server.PrimeCache(ctx, &initialStartup)
		go server.FetchTmdb(ctx)
		go server.UrlLogCache.Start()

		return server, nil
	})

	return Result{
		TmdbAnalysis: lazyAnalysis,
		AppHook: fx.Hook{
			OnStop: func(ctx context.Context) error {
				if server != nil {
					closewg.Add(1)
				}
				close(stopped)
				closewg.Wait()
				return nil
			},
		},
	}, nil

}
