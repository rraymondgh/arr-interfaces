package servarr

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rraymondgh/arr-interface/internal/gqlclient"
	"golift.io/starr"
	"golift.io/starr/radarr"
)

type Radarr struct {
	*ServarrInstance
}

func (arr Radarr) GetIndexers(ctx context.Context) ([]*ArrIndexer, error) {
	indexers, err := arr.Radarr.GetIndexersContext(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]*ArrIndexer, len(indexers))
	err = RecastIndexerSlice(indexers, out)
	return out, err
}

func (arr Radarr) SearchRelease(ctx context.Context, content *Content, indexer *ArrIndexer) ([]*ArrRelease, error) {
	starr := arr.Radarr
	tmdbid, err := strconv.ParseInt(content.Content.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	movies, err := starr.GetMovieContext(ctx, &radarr.GetMovie{TMDBID: tmdbid})
	if err != nil {
		return nil, err
	}
	if len(movies) != 1 {
		return nil, fmt.Errorf("movie not found in radarr [%d]", tmdbid)
	}

	rel, err := starr.SearchReleaseContext(ctx, movies[0].ID)
	if err != nil {
		return nil, err
	}
	out := make([]*ArrRelease, len(rel))
	err = RecastReleaseSlice(rel, out)

	return out, err
}

func (arr Radarr) GrabRelease(ctx context.Context, indexer *ArrIndexer, release *ArrRelease) error {
	_, err := arr.Radarr.GrabReleaseContext(ctx, &radarr.Release{
		GUID:      release.GUID,
		IndexerID: indexer.ID,
	})
	return err
}

func (arr Radarr) UpdateIndexers(ctx context.Context, indexerIds []int64, enabled bool) error {
	_, err := arr.Radarr.UpdateIndexersContext(ctx, &starr.BulkIndexer{IDs: indexerIds, EnableInteractiveSearch: &enabled})
	return err
}

func (arr Radarr) GetDownloadClients(ctx context.Context) ([]*ArrClient, error) {
	clients, err := arr.Radarr.GetDownloadClientsContext(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]*ArrClient, len(clients))
	err = RecastClientDownloadSlice(clients, out)

	return out, err
}

func NewRadarr(content *Content, config *Config) *ServarrInstance {
	if content.ContentType == gqlclient.ContentTypeMovie &&
		config.Radarr.ApiKey != "private" {
		dl := &ServarrInstance{
			IndexerName: fmt.Sprintf("%s (Prowlarr)", config.IndexerName),
			Radarr:      radarr.New(starr.New(config.Radarr.ApiKey, config.Radarr.Url, 0)),
		}
		dl.Arr = Radarr{dl}
		return dl
	}
	return nil
}
