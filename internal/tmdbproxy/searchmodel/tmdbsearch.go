package searchmodel

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"sync"
	"time"

	"github.com/meilisearch/meilisearch-go"
	"github.com/rraymondgh/arr-interfaces/internal/boilerplate/cachestruct"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const indexUID = "tmdb"

type TmdbSearchCache = cachestruct.Cache[string, any]

func (o TmdbSearch) pk() string {
	return Tmdb{MediaType: o.MediaType, ID: o.ID}.PrimaryKey()
}

func (o TmdbSearch) FindOne(ctx context.Context, client meilisearch.ServiceManager, cache *TmdbSearchCache, v interface{}) error {
	if cache.Has(o.pk()) {
		return cache.Get(o.pk(), &v)
	} else {
		var rawDoc map[string]interface{}
		err := client.Index(indexUID).GetDocumentWithContext(ctx, o.pk(), &meilisearch.DocumentQuery{}, &rawDoc)
		if err != nil {
			merr, ok := err.(*meilisearch.Error)
			if ok && merr.MeilisearchApiError.Code == "document_not_found" {
				err = ErrNoDocuments
			}
			return err
		}
		cache.Set(o.pk(), rawDoc)
		return cache.Get(o.pk(), &v)
	}
}

func (o TmdbSearch) Search(ctx context.Context, client meilisearch.ServiceManager) *SearchResponse {
	searchAttr := []string{"title", "original_title"}
	if o.MediaType == "tv" {
		searchAttr = []string{"name", "original_name", "origin_country"}
	}
	m, err := client.Index(indexUID).SearchWithContext(ctx, o.Query, &meilisearch.SearchRequest{
		ShowRankingScore:        true,
		ShowRankingScoreDetails: true,
		Filter:                  []string{fmt.Sprintf("media_type = %s", o.MediaType)},
		AttributesToSearchOn:    searchAttr,
		Limit:                   3,
		RankingScoreThreshold:   .75,
	})

	return &SearchResponse{m, err, true}

}

func (o TmdbSearch) InsertOne(ctx context.Context, client meilisearch.ServiceManager, content interface{}) error {
	_, err := client.Index(indexUID).AddDocumentsWithContext(ctx, content)
	if err != nil {
		return err
	}

	return err
}

func (o TmdbSearch) CreateIndexes(ctx context.Context, client meilisearch.ServiceManager, config config.Config, wg *sync.WaitGroup) error {
	idx, err := client.GetIndex(indexUID)
	if err != nil || idx.UID != indexUID {
		ti, err := client.CreateIndexWithContext(ctx, &meilisearch.IndexConfig{
			Uid:        indexUID,
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

	tasks := make([]*meilisearch.TaskInfo, 4)
	current, _ := client.Index(indexUID).GetSearchableAttributesWithContext(ctx)
	searchable := []string{"title", "original_title", "name", "original_name", "origin_country"}
	if !reflect.DeepEqual(*current, searchable) {
		tasks[0], err = client.Index(indexUID).UpdateSearchableAttributesWithContext(ctx, &searchable)
		if err != nil {
			return err
		}
	}

	current, _ = client.Index(indexUID).GetFilterableAttributesWithContext(ctx)
	filterable := []string{"media_type", "id"}
	slices.Sort(filterable)
	slices.Sort(*current)
	if !slices.Equal(*current, filterable) {
		tasks[1], err = client.Index(indexUID).UpdateFilterableAttributesWithContext(ctx, &filterable)
		if err != nil {
			return err
		}
	}

	synonyms, _ := client.Index(indexUID).GetSynonymsWithContext(ctx)
	if !reflect.DeepEqual(*synonyms, config.Meilisearch.Synonyms) {
		tasks[2], err = client.Index(indexUID).UpdateSynonymsWithContext(ctx, &config.Meilisearch.Synonyms)
		if err != nil {
			return err
		}
	}

	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	stopWords := config.Meilisearch.StopWords
	for _, word := range stopWords {
		if isAlpha(word) {
			stopWords = append(stopWords, cases.Title(language.English, cases.Compact).String(word))
		}
	}
	current, _ = client.Index(indexUID).GetStopWordsWithContext(ctx)
	slices.Sort(*current)
	slices.Sort(stopWords)
	if !slices.Equal(stopWords, *current) {
		tasks[3], err = client.Index(indexUID).UpdateStopWords(&stopWords)
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
