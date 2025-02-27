package tmdbproxy

import (
	"context"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"

	api "github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/oapi"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/requestworker"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/searchmodel"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/similarity"

	"go.uber.org/zap"
)

type TmdbAPI struct {
	api.ServerInterface
	Log      *zap.SugaredLogger
	Analysis *requestworker.TmdbAnalysis
}

func (a TmdbAPI) analysis() *requestworker.TmdbAnalysis {
	return a.Analysis
}

func (a TmdbAPI) meilisearch(ctx context.Context, query string, mediaType string, year int32) ([]*api.FindBy, error) {
	if len(query) == 0 {
		return nil, nil
	}

	var results []*api.FindBy
	err := searchmodel.TmdbSearch{Query: query, Year: int(year), MediaType: mediaType}.Search(ctx, a.analysis().MeiliClient).Decode(&results)
	if err != nil {
		return nil, err
	}

	sim := similarity.SimilarityHelper{}
	idx := sim.BestMatch(query, mediaType, year, results)
	if idx >= 0 {
		return results[idx : idx+1], nil
	}

	return nil, nil
}

func (a TmdbAPI) findById(ctx context.Context, externalId string, params api.FindByIdParams) (*api.FindById200JSONResponse, error) {
	type res struct {
		Id        int64  `json:"id"`
		MediaType string `json:"media_type"`
	}
	var results []res
	err := searchmodel.TmdbExternalId{ExternalId: externalId, ExternalSource: string(params.ExternalSource)}.Search(ctx, a.analysis().MeiliClient).Decode(&results)
	if err != nil {
		return nil, err
	}
	var movies []*api.FindBy
	var tv []*api.FindBy
	for _, found := range results {
		var findby api.FindBy
		err = searchmodel.TmdbSearch{ID: int(found.Id), MediaType: found.MediaType}.FindOne(ctx, a.analysis().MeiliClient, a.analysis().TmdbSearchCache, &findby)
		if err != nil {
			return nil, err
		}
		if found.MediaType == "movie" {
			movies = append(movies, &findby)
		} else if found.MediaType == "tv" {
			tv = append(tv, &findby)
		}
	}

	resp := api.FindById200JSONResponse{
		MovieResults: movies,
		TvResults:    tv,
	}
	return &resp, nil

}

func (a TmdbAPI) FindById(c *gin.Context, externalId string, params api.FindByIdParams) {
	statusCode := http.StatusOK

	resp, err := a.findById(c, externalId, params)
	if err != nil {
		a.Log.Warn(err)
		statusCode = http.StatusMethodNotAllowed
	}

	c.JSON(statusCode, resp)
}

func (a TmdbAPI) searchMedia(c *gin.Context, mediaType string, query string, year int32) (*api.Search200JSONResponse, error) {
	var findby []*api.FindBy
	webstrip := regexp.MustCompile(`^www [0-9a-zA-Z]+ [a-zA-Z]{2,5}[ ]+-`)

	log := searchmodel.UrlLog{}.New(c)
	log, err := log.FindOne(a.analysis().UrlLogCache)
	if err != nil {
		return nil, err
	}
	if webstrip.MatchString(query) {
		query = webstrip.ReplaceAllString(query, "")
		log.Query = query
		log.QueryChanged = true
	}
	if log.TmdbStatus == searchmodel.StatusFound {
		var res api.FindBy
		err = searchmodel.TmdbSearch{
			MediaType: mediaType,
			ID:        int(*log.TmdbId),
		}.FindOne(c, a.analysis().MeiliClient, a.analysis().TmdbSearchCache, &res)
		if err != nil {
			return nil, err
		}
		findby = append(findby, &res)

	} else if log.TmdbStatus == searchmodel.StatusNotFound {
		err = log.UpdateOne(a.analysis().UrlLogCache)
		if err != nil {
			return nil, err
		}

		a.analysis().SubmitTmdb(log)

	} else {
		findby, err = a.meilisearch(c, query, mediaType, year)
		if err != nil {
			return nil, err
		}

		if len(findby) == 0 {
			log.TmdbStatus = searchmodel.StatusNotFound
		} else {
			log.TmdbStatus = searchmodel.StatusFound
			log.TmdbId = &findby[0].Id
		}
		err = log.UpdateOne(a.analysis().UrlLogCache)
		if err != nil {
			return nil, err
		}
	}

	resp := api.Search200JSONResponse{
		Results:      findby,
		Page:         1,
		TotalPages:   1,
		TotalResults: len(findby),
	}

	return &resp, nil
}

func (a TmdbAPI) SearchMovie(c *gin.Context, params api.SearchMovieParams) {
	statusCode := http.StatusOK
	var year int
	if params.Year != nil {
		year, _ = strconv.Atoi(*params.Year)
	} else {
		year = -1
	}
	resp, err := a.searchMedia(c, "movie", params.Query, int32(year))
	if err != nil && err != searchmodel.ErrNoDocuments {
		a.Log.Warn(err)
		statusCode = http.StatusMethodNotAllowed
	}

	c.JSON(statusCode, resp)
}

func (a TmdbAPI) SearchTv(c *gin.Context, params api.SearchTvParams) {
	statusCode := http.StatusOK

	var year int32
	if params.FirstAirDateYear != nil {
		year = *params.FirstAirDateYear
	} else {
		year = -1
	}
	resp, err := a.searchMedia(c, "tv", params.Query, year)
	if err != nil && err != searchmodel.ErrNoDocuments {
		a.Log.Warn(err)
		statusCode = http.StatusMethodNotAllowed
	}
	c.JSON(statusCode, resp)
}

// Details
// (GET /3/movie/{movie_id})
func (a TmdbAPI) MovieDetails(c *gin.Context, movieId int32, params api.MovieDetailsParams) {
	statusCode := http.StatusOK

	var resp api.MovieDetails
	err := searchmodel.TmdbSearch{ID: int(movieId), MediaType: "movie"}.FindOne(c, a.analysis().MeiliClient, a.analysis().TmdbSearchCache, &resp)

	if err != nil {
		a.Log.Warn(c.Request.URL.Path, err)
		statusCode = http.StatusMethodNotAllowed
	}

	c.JSON(statusCode, resp)
}

// Details
// (GET /3/tv/{series_id})
func (a TmdbAPI) TvSeriesDetails(c *gin.Context, seriesId int32, params api.TvSeriesDetailsParams) {
	statusCode := http.StatusOK

	var resp api.TvSeriesDetails
	err := searchmodel.TmdbSearch{ID: int(seriesId), MediaType: "tv"}.FindOne(c, a.analysis().MeiliClient, a.analysis().TmdbSearchCache, &resp)
	if err != nil {
		a.Log.Warn(c.Request.URL.Path, err)
		statusCode = http.StatusMethodNotAllowed
	}
	c.JSON(statusCode, resp)

}

// Validate Key
// (GET /3/authentication)
func (a TmdbAPI) AuthenticationValidateKey(c *gin.Context) {
	c.JSON(http.StatusOK, api.AuthenticationValidateKey200JSONResponse{
		StatusCode:    0,
		Success:       true,
		StatusMessage: "",
	})
}

func (a TmdbAPI) Logs(c *gin.Context) {
	a.Log.Infof("cache len: %v", a.analysis().UrlLogCache.Len())
	items, err := a.analysis().UrlLogCache.Items()
	statusCode := http.StatusOK
	if err != nil {
		a.Log.Warn(c.Request.URL.Path, err)
		statusCode = http.StatusMethodNotAllowed
	}

	c.JSON(statusCode, items)
}
