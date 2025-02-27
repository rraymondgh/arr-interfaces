package meiliclient

import (
	"context"

	"github.com/meilisearch/meilisearch-go"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/lazy"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	Config Config
}

type Result struct {
	fx.Out
	Client  lazy.Lazy[meilisearch.ServiceManager]
	AppHook fx.Hook `group:"app_hooks"`
}

func New(p Params) (Result, error) {
	stopped := make(chan struct{})
	lazyClient := lazy.New(func() (meilisearch.ServiceManager, error) {
		client := meilisearch.New(p.Config.Uri, meilisearch.WithAPIKey(p.Config.MasterKey))
		_, err := client.GetStats()
		go func() {
			<-stopped
			client.Close()
		}()

		return client, err
	})

	return Result{
		Client: lazyClient,
		AppHook: fx.Hook{
			OnStop: func(context.Context) error {
				close(stopped)
				return nil
			},
		},
	}, nil
}
