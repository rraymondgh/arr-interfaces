package client

import (
	"context"
	"errors"
	"slices"

	"github.com/rraymondgh/arr-interface/internal/servarr"
)

type servarrClient struct {
	commonClient
}

func (c servarrClient) setIndexerEnabled(arr servarr.ServarrSpecific, ctx context.Context, name string, indexers []*servarr.ArrIndexer, enable bool) error {
	var ids []int64

	for _, i := range indexers {
		if i.EnableInteractiveSearch && i.Name != name {
			ids = append(ids, i.ID)
		}
	}

	if len(ids) > 0 {
		return arr.UpdateIndexers(ctx, ids, enable)
	}

	return nil

}

func (c *servarrClient) downloadOne(ctx context.Context, content *Content, clientId string) error {
	d, err := servarr.New(c.config, content)
	if err != nil {
		return err
	}
	indexers, indexer, err := servarr.GetIndexers(d.Arr, ctx, d.IndexerName)
	if err != nil {
		return err
	}

	if c.config.OnlySearchBitmagnet {
		// any indexers that are disabled, defer re-enablement
		defer c.setIndexerEnabled(d.Arr, ctx, d.IndexerName, indexers, true)
		err = c.setIndexerEnabled(d.Arr, ctx, d.IndexerName, indexers, false)
		if err != nil {
			return err
		}

	}

	releases, err := d.Arr.SearchRelease(ctx, content, indexer)
	if err != nil {
		return err
	}

	i := slices.IndexFunc(releases, func(r *servarr.ArrRelease) bool { return r.GUID == content.InfoHash.String() })
	if i == -1 {
		return errors.New("download not found")
	}
	err = d.Arr.GrabRelease(ctx, indexer, releases[i])

	return err

}

func (c *servarrClient) AddInfoHashes(ctx context.Context, req AddInfoHashesRequest) error {
	return AddInfoHashes(ctx, req, c.downloadOne, c.config)
}
