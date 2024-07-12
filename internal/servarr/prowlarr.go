package servarr

import (
	"context"

	"golift.io/starr"
	"golift.io/starr/prowlarr"
)

type Prowlarr struct {
	*ServarrInstance
}

func (arr Prowlarr) GetIndexers(ctx context.Context) ([]*ArrIndexer, error) {
	indexers, err := arr.Prowlarr.GetIndexersContext(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]*ArrIndexer, len(indexers))
	err = RecastIndexerSlice(indexers, out)
	return out, err
}

func (arr Prowlarr) SearchRelease(ctx context.Context, content *Content, indexer *ArrIndexer) ([]*ArrRelease, error) {
	rel, err := arr.Prowlarr.SearchContext(ctx, prowlarr.SearchInput{
		Query:      content.InfoHash.String(),
		IndexerIDs: []int64{indexer.ID},
		Type:       "search",
	})
	if err != nil {
		return nil, err
	}

	return rel, err

}

func (arr Prowlarr) GrabRelease(ctx context.Context, indexer *ArrIndexer, release *ArrRelease) error {
	_, err := arr.Prowlarr.GrabSearchContext(ctx, &ArrRelease{
		GUID:      release.GUID,
		IndexerID: indexer.ID,
	})
	return err
}

func (arr Prowlarr) UpdateIndexers(ctx context.Context, indexerIds []int64, enabled bool) error {
	return nil
}

func (arr Prowlarr) GetDownloadClients(ctx context.Context) ([]*ArrClient, error) {
	return arr.Prowlarr.GetDownloadClientsContext(ctx)
}

func NewProwlarr(content *Content, config *Config) *ServarrInstance {
	if config.Prowlarr.ApiKey != "private" {
		dl := &ServarrInstance{
			IndexerName: config.IndexerName,
			Prowlarr:    prowlarr.New(starr.New(config.Prowlarr.ApiKey, config.Prowlarr.Url, 0)),
		}
		dl.Arr = Prowlarr{dl}
		return dl
	}
	return nil
}
