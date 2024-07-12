package servarr

import (
	"context"
	"fmt"
	"slices"
	"strconv"

	"github.com/rraymondgh/arr-interface/internal/gqlclient"
	"golift.io/starr"
	"golift.io/starr/sonarr"
)

type Sonarr struct {
	*ServarrInstance
}

func (arr Sonarr) GetIndexers(ctx context.Context) ([]*ArrIndexer, error) {
	indexers, err := arr.Sonarr.GetIndexersContext(ctx)
	return indexers, err
}

func (arr Sonarr) SearchRelease(ctx context.Context, content *Content, indexer *ArrIndexer) ([]*ArrRelease, error) {
	starr := arr.Sonarr

	i := slices.IndexFunc(content.Content.Attributes, func(a attribute) bool { return a.Source == "tvdb" })
	if i == -1 {
		return nil, fmt.Errorf("no TVDB id")
	}

	tvdbid, err := strconv.ParseInt(content.Content.Attributes[i].Value, 10, 64)
	if err != nil {
		return nil, err
	}
	series, err := starr.GetSeriesContext(ctx, tvdbid)
	if err != nil {
		return nil, err
	}
	if len(series) != 1 {
		return nil, fmt.Errorf("failed to find tvdb ID [%d]", tvdbid)
	}
	episodes, err := starr.GetSeriesEpisodesContext(ctx, &sonarr.GetEpisode{SeriesID: series[0].ID})
	if err != nil {
		return nil, err
	}
	i = slices.IndexFunc(episodes, func(e *sonarr.Episode) bool {
		return e.SeasonNumber == int(content.Episodes.Seasons[0].Season) &&
			e.EpisodeNumber == int(content.Episodes.Seasons[0].Episodes[0])
	})
	if i == -1 {
		return nil, fmt.Errorf("episode not found")
	}

	rel, err := starr.SearchReleaseContext(ctx, &sonarr.SearchRelease{
		SeriesID:     series[0].ID,
		SeasonNumber: episodes[i].SeasonNumber,
		EpisodeID:    episodes[i].ID,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*ArrRelease, len(rel))
	err = RecastReleaseSlice(rel, out)

	return out, err
}

func (arr Sonarr) GrabRelease(ctx context.Context, indexer *ArrIndexer, release *ArrRelease) error {
	_, err := arr.Sonarr.GrabReleaseContext(ctx, &sonarr.Release{
		GUID:      release.GUID,
		IndexerID: indexer.ID,
	})
	return err
}

func (arr Sonarr) UpdateIndexers(ctx context.Context, indexerIds []int64, enabled bool) error {
	_, err := arr.Sonarr.UpdateIndexersContext(ctx, &starr.BulkIndexer{IDs: indexerIds, EnableInteractiveSearch: &enabled})
	return err
}

func (arr Sonarr) GetDownloadClients(ctx context.Context) ([]*ArrClient, error) {
	clients, err := arr.Sonarr.GetDownloadClientsContext(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]*ArrClient, len(clients))
	err = RecastClientDownloadSlice(clients, out)

	return out, err
}

func NewSonarr(content *Content, config *Config) *ServarrInstance {
	if content.ContentType == gqlclient.ContentTypeTvShow &&
		len(content.Episodes.Seasons) == 1 &&
		len(content.Episodes.Seasons[0].Episodes) == 1 &&
		config.Sonarr.ApiKey != "private" {
		dl := &ServarrInstance{
			IndexerName: fmt.Sprintf("%s (Prowlarr)", config.IndexerName),
			Sonarr:      sonarr.New(starr.New(config.Sonarr.ApiKey, config.Sonarr.Url, 0)),
		}
		dl.Arr = Sonarr{dl}
		return dl
	}
	return nil
}
