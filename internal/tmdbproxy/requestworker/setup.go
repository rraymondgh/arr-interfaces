package requestworker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/config"
	api "github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/oapi"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/searchmodel"
)

func (a TmdbAnalysis) SourceAndStore(
	ctx context.Context,
	arr config.Arr,
	mediaType string,
) error {
	if arr.ApiKey == "" {
		return nil
	}
	type arrId struct {
		TmdbId int `json:"tmdbId"`
	}
	var inArr []arrId
	var toCache []arrId

	endpoint := "series"
	if mediaType == "movie" {
		endpoint = "movie"
	}

	url := fmt.Sprintf("%s/api/v3/%s", arr.Url, endpoint)
	resp, err := resty.New().R().
		SetHeader("X-Api-Key", arr.ApiKey).
		SetResult(&inArr).
		Get(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("[%d] (%s) %s", resp.StatusCode(), resp.Request.URL, resp.Status())
	}

	for _, id := range inArr {
		var v interface{}
		err = searchmodel.TmdbSearch{ID: id.TmdbId, MediaType: mediaType}.FindOne(ctx, a.MeiliClient, a.TmdbSearchCache, &v)
		if err == searchmodel.ErrNoDocuments && id.TmdbId > 0 {
			toCache = append(toCache, id)
		} else if err == searchmodel.ErrNoDocuments && id.TmdbId == 0 {
		} else if err != nil {
			return err
		}
	}

	a.Log.Debug(fmt.Sprintf("%s ids: %+v", mediaType, toCache))

	n := time.Now()

	for _, id := range toCache {
		pk := searchmodel.Tmdb{MediaType: mediaType, ID: id.TmdbId}.PrimaryKey()
		if mediaType == "movie" {
			err = persist(ctx, a, mediaType, id.TmdbId, searchmodel.UrlLog{ApiKey: a.Config.ApiKey}, api.MovieDetails{
				MediaType:  &mediaType,
				CreatedAt:  &n,
				PrimaryKey: pk,
			})
		} else {
			err = persist(ctx, a, mediaType, id.TmdbId, searchmodel.UrlLog{ApiKey: a.Config.ApiKey}, api.TvSeriesDetails{
				MediaType:  &mediaType,
				CreatedAt:  &n,
				PrimaryKey: pk,
			})

		}
		if err != nil {
			return err
		}
	}

	if len(toCache) > 0 {
		a.Log.Info("reseting NotFound UrlLogCache")
		err = searchmodel.UrlLog{TmdbStatus: searchmodel.StatusNotFound}.DeleteMany(a.UrlLogCache)
		if err != nil {
			return err
		}
	}

	return nil

}

func (a TmdbAnalysis) InitialDataAndIndexes(ctx context.Context, wg *sync.WaitGroup) {
	err := searchmodel.TmdbSearch{}.CreateIndexes(ctx, a.MeiliClient, a.Config, wg)
	if err != nil {
		a.Log.Warn(err)
	}
	err = searchmodel.TmdbExternalId{}.CreateIndexes(ctx, a.MeiliClient, a.Config, wg)
	if err != nil {
		a.Log.Warn(err)
	}
	err = a.SourceAndStore(ctx, a.Config.Sonarr, "tv")
	if err != nil {
		a.Log.Warn(err)
	}
	err = a.SourceAndStore(ctx, a.Config.Radarr, "movie")
	if err != nil {
		a.Log.Warn(err)
	}

	a.Log.Debug("done")
}
