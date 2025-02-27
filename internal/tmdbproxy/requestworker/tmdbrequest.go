package requestworker

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/meilisearch/meilisearch-go"
	api "github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/oapi"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/searchmodel"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/similarity"
)

func persist[T api.MovieDetails | api.TvSeriesDetails](
	ctx context.Context,
	tmdb TmdbAnalysis,
	mediaType string,
	id int,
	req searchmodel.UrlLog,
	content T,
) error {
	url := fmt.Sprintf("%s/{mediaType}/{id}", tmdb.Config.BaseUrl)
	resp, err := resty.New().R().
		SetPathParam("mediaType", mediaType).
		SetPathParam("id", strconv.Itoa(id)).
		SetQueryParam("api_key", req.ApiKey).
		SetQueryParam("append_to_response", "external_ids").
		SetResult(&content).
		Get(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("[%d] %s", resp.StatusCode(), resp.Body())
	}

	if resp.StatusCode() == 204 {
		return fmt.Errorf("NO CONTENT %s", resp.Request.URL)
	}
	err = searchmodel.TmdbSearch{MediaType: mediaType, ID: id}.InsertOne(ctx, tmdb.MeiliClient, content)
	if err != nil {
		return err
	}

	var ids_content api.ExternalIdsResponse
	switch any(content).(type) {
	case api.MovieDetails:
		ids_content = any(content).(api.MovieDetails).ExternalIDs
	case api.TvSeriesDetails:
		ids_content = any(content).(api.TvSeriesDetails).ExternalIDs
	}

	now := time.Now()
	ids_content.MediaType = &mediaType
	ids_content.CreatedAt = &now
	ids_content.PrimaryKey = searchmodel.Tmdb{MediaType: mediaType, ID: id}.PrimaryKey()

	return searchmodel.TmdbExternalId{}.InsertOne(ctx, tmdb.MeiliClient, &ids_content)
}

func (a TmdbAnalysis) fetchTmdb(ctx context.Context, req searchmodel.UrlLog) {

	u, _ := url.Parse(a.Config.BaseUrl)
	url := fmt.Sprintf("%s://%s%s?%s", u.Scheme, u.Host, strings.TrimPrefix(req.Path, "/tmdb"), req.RawQuery)
	var tmdbres api.Search200JSONResponse
	_, err := resty.New().R().SetResult(&tmdbres).Get(url)
	if err != nil {
		a.Log.Warn(err)
		return
	}
	sim := similarity.SimilarityHelper{}

	var mediaType string
	var year int32
	switch req.Path {
	case "/tmdb/3/search/tv":
		mediaType = "tv"
		if req.FirstAirDateYear == nil {
			year = -1
		} else {
			year = int32(*req.FirstAirDateYear)
		}
	case "/tmdb/3/search/movie":
		mediaType = "movie"
		if req.Year == nil {
			year = -1
		} else {
			year = int32(*req.Year)
		}
	}

	idx := sim.BestMatch(req.Query, mediaType, year, tmdbres.Results)
	req.TmdbStatus = searchmodel.StatusUndefined
	now := time.Now()
	req.LastTmdbAttempt = &now
	if idx != -1 {
		req.TmdbId = &tmdbres.Results[idx].Id
		req.TmdbStatus = searchmodel.StatusFound
	}
	err = req.UpdateOne(a.UrlLogCache)
	if err != nil {
		a.Log.Warn(err)
		return
	}
	if idx == -1 {
		return
	}

	goodres := tmdbres.Results[idx]
	pk := fmt.Sprintf("%s%v", mediaType, goodres.Id)

	var name string

	switch mediaType {
	case "movie":
		name = *goodres.Title
		err = persist[api.MovieDetails](ctx, a, mediaType, int(goodres.Id), req, api.MovieDetails{MediaType: &mediaType, CreatedAt: &now, PrimaryKey: pk})
	case "tv":
		name = *goodres.Name
		err = persist[api.TvSeriesDetails](ctx, a, mediaType, int(goodres.Id), req, api.TvSeriesDetails{MediaType: &mediaType, CreatedAt: &now, PrimaryKey: pk})

	}
	if err != nil {
		a.Log.Warn(err)
		return
	}

	a.Log.Infof("%s [%v] %s (%s)", mediaType, goodres.Id, name, req.Query)

}

func (a TmdbAnalysis) FetchTmdb(ctx context.Context) {
	a.Log.Debug("starting fetcher")

	for {
		req, ok := <-a.UniqueQueue.Front()

		if !ok {
			a.Log.Debug("shutting down fetcher")
			return
		}

		a.fetchTmdb(ctx, req)
	}
}

func (a TmdbAnalysis) SubmitTmdb(log searchmodel.UrlLog) {
	if a.Config.Tmdb.FetchMissing &&
		log.Counter >= a.Config.Tmdb.MinRequests &&
		(log.LastTmdbAttempt == nil ||
			(*log.LastTmdbAttempt).Before(time.Now().Add(time.Duration(a.Config.Tmdb.BackoffMinutes*-1)*time.Minute))) {
		// try sourcing from tmdb
		now := time.Now()
		log.LastTmdbAttempt = &now
		log.UpdateOne(a.UrlLogCache)
		a.UniqueQueue.Back() <- log
	}

}

func (a TmdbAnalysis) PrimeCache(ctx context.Context, wg *sync.WaitGroup) {
	var r meilisearch.DocumentsResult

	wg.Wait()
	err := a.MeiliClient.Index("tmdb").GetDocumentsWithContext(ctx, &meilisearch.DocumentsQuery{
		Limit: 10000,
	}, &r)
	if err != nil {
		a.Log.Warn(err)
		return
	}
	for _, m := range r.Results {
		a.TmdbSearchCache.Set(m["pk"].(string), m)
	}
	a.Log.Debug("cache primed")

}
