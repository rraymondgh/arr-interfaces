package searchmodel

import (
	"context"
	"slices"
	"sync"
	"time"

	"github.com/meilisearch/meilisearch-go"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/config"
)

const idIndexUID = "tmdb_external"

func (o TmdbExternalId) Search(ctx context.Context, client meilisearch.ServiceManager) *SearchResponse {
	searchAttr := []string{o.ExternalSource}
	m, err := client.Index(idIndexUID).SearchWithContext(ctx, o.ExternalId, &meilisearch.SearchRequest{
		AttributesToSearchOn:  searchAttr,
		Limit:                 3,
		RankingScoreThreshold: 1,
	})

	return &SearchResponse{m, err, true}

}

func (o TmdbExternalId) InsertOne(ctx context.Context, client meilisearch.ServiceManager, content interface{}) error {
	_, err := client.Index(idIndexUID).AddDocumentsWithContext(ctx, content)
	return err
}

func (o TmdbExternalId) CreateIndexes(ctx context.Context, client meilisearch.ServiceManager, config config.Config, wg *sync.WaitGroup) error {
	idx, err := client.GetIndex(idIndexUID)
	if err != nil || idx.UID != idIndexUID {

		ti, err := client.CreateIndexWithContext(ctx, &meilisearch.IndexConfig{
			Uid:        idIndexUID,
			PrimaryKey: "pk",
		})
		if err != nil {
			return err
		}
		_, err = client.WaitForTaskWithContext(ctx, ti.TaskUID, time.Duration(30)*time.Second)
		if err != nil {
			return err
		}
	}
	wg.Done()

	searchable := []string{
		"facebook_id",
		"freebase_id",
		"freebase_mid",
		"imdb_id",
		"instagram_id",
		"tvdb_id",
		"tvrage_id",
		"twitter_id",
		"wikidata_id",
	}

	tasks := make([]*meilisearch.TaskInfo, 3)

	current, _ := client.Index(idIndexUID).GetSearchableAttributesWithContext(ctx)
	slices.Sort(*current)
	slices.Sort(searchable)
	if !slices.Equal(*current, searchable) {
		tasks[0], err = client.Index(idIndexUID).UpdateSearchableAttributesWithContext(ctx, &searchable)
		if err != nil {
			return err
		}
	}

	current, _ = client.Index(idIndexUID).GetRankingRulesWithContext(ctx)
	rules := []string{"words"}
	if !slices.Equal(*current, rules) {
		tasks[1], err = client.Index(idIndexUID).UpdateRankingRulesWithContext(ctx, &rules)
		if err != nil {
			return err
		}
	}

	typo, _ := client.Index(idIndexUID).GetTypoToleranceWithContext(ctx)
	if typo.Enabled {
		typo.Enabled = false
		tasks[2], err = client.Index(idIndexUID).UpdateTypoToleranceWithContext(ctx, typo)
		if err != nil {
			return err
		}
	}

	for _, task := range tasks {
		if task != nil {
			_, err = client.WaitForTaskWithContext(ctx, task.TaskUID, time.Duration(30)*time.Second)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
