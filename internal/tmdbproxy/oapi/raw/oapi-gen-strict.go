// Package tmdbproxyoapimodel provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package tmdbproxyoapimodel

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

const (
	Sec0Scopes = "sec0.Scopes"
)

// Defines values for FindByIdParamsExternalSource.
const (
	Empty       FindByIdParamsExternalSource = ""
	FacebookId  FindByIdParamsExternalSource = "facebook_id"
	ImdbId      FindByIdParamsExternalSource = "imdb_id"
	InstagramId FindByIdParamsExternalSource = "instagram_id"
	TiktokId    FindByIdParamsExternalSource = "tiktok_id"
	TvdbId      FindByIdParamsExternalSource = "tvdb_id"
	TwitterId   FindByIdParamsExternalSource = "twitter_id"
	WikidataId  FindByIdParamsExternalSource = "wikidata_id"
	YoutubeId   FindByIdParamsExternalSource = "youtube_id"
)

// FindByIdParams defines parameters for FindById.
type FindByIdParams struct {
	ExternalSource FindByIdParamsExternalSource `form:"external_source" json:"external_source"`
	Language       *string                      `form:"language,omitempty" json:"language,omitempty"`
}

// FindByIdParamsExternalSource defines parameters for FindById.
type FindByIdParamsExternalSource string

// MovieDetailsParams defines parameters for MovieDetails.
type MovieDetailsParams struct {
	// AppendToResponse comma separated list of endpoints within this namespace, 20 items max
	AppendToResponse *string `form:"append_to_response,omitempty" json:"append_to_response,omitempty"`
	Language         *string `form:"language,omitempty" json:"language,omitempty"`
}

// SearchMovieParams defines parameters for SearchMovie.
type SearchMovieParams struct {
	Query              string  `form:"query" json:"query"`
	IncludeAdult       *bool   `form:"include_adult,omitempty" json:"include_adult,omitempty"`
	Language           *string `form:"language,omitempty" json:"language,omitempty"`
	PrimaryReleaseYear *string `form:"primary_release_year,omitempty" json:"primary_release_year,omitempty"`
	Page               *int32  `form:"page,omitempty" json:"page,omitempty"`
	Region             *string `form:"region,omitempty" json:"region,omitempty"`
	Year               *string `form:"year,omitempty" json:"year,omitempty"`
}

// SearchTvParams defines parameters for SearchTv.
type SearchTvParams struct {
	Query string `form:"query" json:"query"`

	// FirstAirDateYear Search only the first air date. Valid values are: 1000..9999
	FirstAirDateYear *int32  `form:"first_air_date_year,omitempty" json:"first_air_date_year,omitempty"`
	IncludeAdult     *bool   `form:"include_adult,omitempty" json:"include_adult,omitempty"`
	Language         *string `form:"language,omitempty" json:"language,omitempty"`
	Page             *int32  `form:"page,omitempty" json:"page,omitempty"`

	// Year Search the first air date and all episode air dates. Valid values are: 1000..9999
	Year *int32 `form:"year,omitempty" json:"year,omitempty"`
}

// TvSeriesDetailsParams defines parameters for TvSeriesDetails.
type TvSeriesDetailsParams struct {
	// AppendToResponse comma separated list of endpoints within this namespace, 20 items max
	AppendToResponse *string `form:"append_to_response,omitempty" json:"append_to_response,omitempty"`
	Language         *string `form:"language,omitempty" json:"language,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Validate Key
	// (GET /3/authentication)
	AuthenticationValidateKey(c *gin.Context)
	// Find By ID
	// (GET /3/find/{external_id})
	FindById(c *gin.Context, externalId string, params FindByIdParams)
	// Details
	// (GET /3/movie/{movie_id})
	MovieDetails(c *gin.Context, movieId int32, params MovieDetailsParams)
	// Movie
	// (GET /3/search/movie)
	SearchMovie(c *gin.Context, params SearchMovieParams)
	// TV
	// (GET /3/search/tv)
	SearchTv(c *gin.Context, params SearchTvParams)
	// Details
	// (GET /3/tv/{series_id})
	TvSeriesDetails(c *gin.Context, seriesId int32, params TvSeriesDetailsParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// AuthenticationValidateKey operation middleware
func (siw *ServerInterfaceWrapper) AuthenticationValidateKey(c *gin.Context) {

	c.Set(Sec0Scopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AuthenticationValidateKey(c)
}

// FindById operation middleware
func (siw *ServerInterfaceWrapper) FindById(c *gin.Context) {

	var err error

	// ------------- Path parameter "external_id" -------------
	var externalId string

	err = runtime.BindStyledParameterWithOptions("simple", "external_id", c.Param("external_id"), &externalId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter external_id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Sec0Scopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params FindByIdParams

	// ------------- Required query parameter "external_source" -------------

	if paramValue := c.Query("external_source"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument external_source is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "external_source", c.Request.URL.Query(), &params.ExternalSource)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter external_source: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "language" -------------

	err = runtime.BindQueryParameter("form", true, false, "language", c.Request.URL.Query(), &params.Language)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter language: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.FindById(c, externalId, params)
}

// MovieDetails operation middleware
func (siw *ServerInterfaceWrapper) MovieDetails(c *gin.Context) {

	var err error

	// ------------- Path parameter "movie_id" -------------
	var movieId int32

	err = runtime.BindStyledParameterWithOptions("simple", "movie_id", c.Param("movie_id"), &movieId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter movie_id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Sec0Scopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params MovieDetailsParams

	// ------------- Optional query parameter "append_to_response" -------------

	err = runtime.BindQueryParameter("form", true, false, "append_to_response", c.Request.URL.Query(), &params.AppendToResponse)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter append_to_response: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "language" -------------

	err = runtime.BindQueryParameter("form", true, false, "language", c.Request.URL.Query(), &params.Language)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter language: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.MovieDetails(c, movieId, params)
}

// SearchMovie operation middleware
func (siw *ServerInterfaceWrapper) SearchMovie(c *gin.Context) {

	var err error

	c.Set(Sec0Scopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchMovieParams

	// ------------- Required query parameter "query" -------------

	if paramValue := c.Query("query"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument query is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "query", c.Request.URL.Query(), &params.Query)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter query: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "include_adult" -------------

	err = runtime.BindQueryParameter("form", true, false, "include_adult", c.Request.URL.Query(), &params.IncludeAdult)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter include_adult: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "language" -------------

	err = runtime.BindQueryParameter("form", true, false, "language", c.Request.URL.Query(), &params.Language)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter language: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "primary_release_year" -------------

	err = runtime.BindQueryParameter("form", true, false, "primary_release_year", c.Request.URL.Query(), &params.PrimaryReleaseYear)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter primary_release_year: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", c.Request.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "region" -------------

	err = runtime.BindQueryParameter("form", true, false, "region", c.Request.URL.Query(), &params.Region)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter region: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "year" -------------

	err = runtime.BindQueryParameter("form", true, false, "year", c.Request.URL.Query(), &params.Year)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter year: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.SearchMovie(c, params)
}

// SearchTv operation middleware
func (siw *ServerInterfaceWrapper) SearchTv(c *gin.Context) {

	var err error

	c.Set(Sec0Scopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchTvParams

	// ------------- Required query parameter "query" -------------

	if paramValue := c.Query("query"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument query is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "query", c.Request.URL.Query(), &params.Query)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter query: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "first_air_date_year" -------------

	err = runtime.BindQueryParameter("form", true, false, "first_air_date_year", c.Request.URL.Query(), &params.FirstAirDateYear)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter first_air_date_year: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "include_adult" -------------

	err = runtime.BindQueryParameter("form", true, false, "include_adult", c.Request.URL.Query(), &params.IncludeAdult)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter include_adult: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "language" -------------

	err = runtime.BindQueryParameter("form", true, false, "language", c.Request.URL.Query(), &params.Language)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter language: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", c.Request.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "year" -------------

	err = runtime.BindQueryParameter("form", true, false, "year", c.Request.URL.Query(), &params.Year)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter year: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.SearchTv(c, params)
}

// TvSeriesDetails operation middleware
func (siw *ServerInterfaceWrapper) TvSeriesDetails(c *gin.Context) {

	var err error

	// ------------- Path parameter "series_id" -------------
	var seriesId int32

	err = runtime.BindStyledParameterWithOptions("simple", "series_id", c.Param("series_id"), &seriesId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter series_id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Sec0Scopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params TvSeriesDetailsParams

	// ------------- Optional query parameter "append_to_response" -------------

	err = runtime.BindQueryParameter("form", true, false, "append_to_response", c.Request.URL.Query(), &params.AppendToResponse)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter append_to_response: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "language" -------------

	err = runtime.BindQueryParameter("form", true, false, "language", c.Request.URL.Query(), &params.Language)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter language: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.TvSeriesDetails(c, seriesId, params)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/3/authentication", wrapper.AuthenticationValidateKey)
	router.GET(options.BaseURL+"/3/find/:external_id", wrapper.FindById)
	router.GET(options.BaseURL+"/3/movie/:movie_id", wrapper.MovieDetails)
	router.GET(options.BaseURL+"/3/search/movie", wrapper.SearchMovie)
	router.GET(options.BaseURL+"/3/search/tv", wrapper.SearchTv)
	router.GET(options.BaseURL+"/3/tv/:series_id", wrapper.TvSeriesDetails)
}

type AuthenticationValidateKeyRequestObject struct {
}

type AuthenticationValidateKeyResponseObject interface {
	VisitAuthenticationValidateKeyResponse(w http.ResponseWriter) error
}

type AuthenticationValidateKey200JSONResponse struct {
	StatusCode    *int    `json:"status_code,omitempty"`
	StatusMessage *string `json:"status_message,omitempty"`
	Success       *bool   `json:"success,omitempty"`
}

func (response AuthenticationValidateKey200JSONResponse) VisitAuthenticationValidateKeyResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AuthenticationValidateKey401JSONResponse struct {
	StatusCode    *int    `json:"status_code,omitempty"`
	StatusMessage *string `json:"status_message,omitempty"`
	Success       *bool   `json:"success,omitempty"`
}

func (response AuthenticationValidateKey401JSONResponse) VisitAuthenticationValidateKeyResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type FindByIdRequestObject struct {
	ExternalId string `json:"external_id"`
	Params     FindByIdParams
}

type FindByIdResponseObject interface {
	VisitFindByIdResponse(w http.ResponseWriter) error
}

type FindById200JSONResponse struct {
	MovieResults *[]struct {
		Adult            *bool    `json:"adult,omitempty"`
		BackdropPath     *string  `json:"backdrop_path,omitempty"`
		GenreIds         *[]int   `json:"genre_ids,omitempty"`
		Id               *int     `json:"id,omitempty"`
		MediaType        *string  `json:"media_type,omitempty"`
		OriginalLanguage *string  `json:"original_language,omitempty"`
		OriginalTitle    *string  `json:"original_title,omitempty"`
		Overview         *string  `json:"overview,omitempty"`
		Popularity       *float32 `json:"popularity,omitempty"`
		PosterPath       *string  `json:"poster_path,omitempty"`
		ReleaseDate      *string  `json:"release_date,omitempty"`
		Title            *string  `json:"title,omitempty"`
		Video            *bool    `json:"video,omitempty"`
		VoteAverage      *float32 `json:"vote_average,omitempty"`
		VoteCount        *int     `json:"vote_count,omitempty"`
	} `json:"movie_results,omitempty"`
	PersonResults    *[]interface{} `json:"person_results,omitempty"`
	TvEpisodeResults *[]interface{} `json:"tv_episode_results,omitempty"`
	TvResults        *[]interface{} `json:"tv_results,omitempty"`
	TvSeasonResults  *[]interface{} `json:"tv_season_results,omitempty"`
}

func (response FindById200JSONResponse) VisitFindByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type MovieDetailsRequestObject struct {
	MovieId int32 `json:"movie_id"`
	Params  MovieDetailsParams
}

type MovieDetailsResponseObject interface {
	VisitMovieDetailsResponse(w http.ResponseWriter) error
}

type MovieDetails200JSONResponse struct {
	Adult               *bool        `json:"adult,omitempty"`
	BackdropPath        *string      `json:"backdrop_path,omitempty"`
	BelongsToCollection *interface{} `json:"belongs_to_collection,omitempty"`
	Budget              *int         `json:"budget,omitempty"`
	Genres              *[]struct {
		Id   *int    `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	} `json:"genres,omitempty"`
	Homepage            *string  `json:"homepage,omitempty"`
	Id                  *int     `json:"id,omitempty"`
	ImdbId              *string  `json:"imdb_id,omitempty"`
	OriginalLanguage    *string  `json:"original_language,omitempty"`
	OriginalTitle       *string  `json:"original_title,omitempty"`
	Overview            *string  `json:"overview,omitempty"`
	Popularity          *float32 `json:"popularity,omitempty"`
	PosterPath          *string  `json:"poster_path,omitempty"`
	ProductionCompanies *[]struct {
		Id            *int    `json:"id,omitempty"`
		LogoPath      *string `json:"logo_path,omitempty"`
		Name          *string `json:"name,omitempty"`
		OriginCountry *string `json:"origin_country,omitempty"`
	} `json:"production_companies,omitempty"`
	ProductionCountries *[]struct {
		Iso31661 *string `json:"iso_3166_1,omitempty"`
		Name     *string `json:"name,omitempty"`
	} `json:"production_countries,omitempty"`
	ReleaseDate     *string `json:"release_date,omitempty"`
	Revenue         *int    `json:"revenue,omitempty"`
	Runtime         *int    `json:"runtime,omitempty"`
	SpokenLanguages *[]struct {
		EnglishName *string `json:"english_name,omitempty"`
		Iso6391     *string `json:"iso_639_1,omitempty"`
		Name        *string `json:"name,omitempty"`
	} `json:"spoken_languages,omitempty"`
	Status      *string  `json:"status,omitempty"`
	Tagline     *string  `json:"tagline,omitempty"`
	Title       *string  `json:"title,omitempty"`
	Video       *bool    `json:"video,omitempty"`
	VoteAverage *float32 `json:"vote_average,omitempty"`
	VoteCount   *int     `json:"vote_count,omitempty"`
}

func (response MovieDetails200JSONResponse) VisitMovieDetailsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type SearchMovieRequestObject struct {
	Params SearchMovieParams
}

type SearchMovieResponseObject interface {
	VisitSearchMovieResponse(w http.ResponseWriter) error
}

type SearchMovie200JSONResponse struct {
	Page    *int `json:"page,omitempty"`
	Results *[]struct {
		Adult            *bool    `json:"adult,omitempty"`
		BackdropPath     *string  `json:"backdrop_path,omitempty"`
		GenreIds         *[]int   `json:"genre_ids,omitempty"`
		Id               *int     `json:"id,omitempty"`
		OriginalLanguage *string  `json:"original_language,omitempty"`
		OriginalTitle    *string  `json:"original_title,omitempty"`
		Overview         *string  `json:"overview,omitempty"`
		Popularity       *float32 `json:"popularity,omitempty"`
		PosterPath       *string  `json:"poster_path,omitempty"`
		ReleaseDate      *string  `json:"release_date,omitempty"`
		Title            *string  `json:"title,omitempty"`
		Video            *bool    `json:"video,omitempty"`
		VoteAverage      *float32 `json:"vote_average,omitempty"`
		VoteCount        *int     `json:"vote_count,omitempty"`
	} `json:"results,omitempty"`
	TotalPages   *int `json:"total_pages,omitempty"`
	TotalResults *int `json:"total_results,omitempty"`
}

func (response SearchMovie200JSONResponse) VisitSearchMovieResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type SearchTvRequestObject struct {
	Params SearchTvParams
}

type SearchTvResponseObject interface {
	VisitSearchTvResponse(w http.ResponseWriter) error
}

type SearchTv200JSONResponse struct {
	Page    *int `json:"page,omitempty"`
	Results *[]struct {
		Adult            *bool     `json:"adult,omitempty"`
		BackdropPath     *string   `json:"backdrop_path,omitempty"`
		FirstAirDate     *string   `json:"first_air_date,omitempty"`
		GenreIds         *[]int    `json:"genre_ids,omitempty"`
		Id               *int      `json:"id,omitempty"`
		Name             *string   `json:"name,omitempty"`
		OriginCountry    *[]string `json:"origin_country,omitempty"`
		OriginalLanguage *string   `json:"original_language,omitempty"`
		OriginalName     *string   `json:"original_name,omitempty"`
		Overview         *string   `json:"overview,omitempty"`
		Popularity       *float32  `json:"popularity,omitempty"`
		PosterPath       *string   `json:"poster_path,omitempty"`
		VoteAverage      *float32  `json:"vote_average,omitempty"`
		VoteCount        *int      `json:"vote_count,omitempty"`
	} `json:"results,omitempty"`
	TotalPages   *int `json:"total_pages,omitempty"`
	TotalResults *int `json:"total_results,omitempty"`
}

func (response SearchTv200JSONResponse) VisitSearchTvResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type TvSeriesDetailsRequestObject struct {
	SeriesId int32 `json:"series_id"`
	Params   TvSeriesDetailsParams
}

type TvSeriesDetailsResponseObject interface {
	VisitTvSeriesDetailsResponse(w http.ResponseWriter) error
}

type TvSeriesDetails200JSONResponse struct {
	Adult        *bool   `json:"adult,omitempty"`
	BackdropPath *string `json:"backdrop_path,omitempty"`
	CreatedBy    *[]struct {
		CreditId    *string `json:"credit_id,omitempty"`
		Gender      *int    `json:"gender,omitempty"`
		Id          *int    `json:"id,omitempty"`
		Name        *string `json:"name,omitempty"`
		ProfilePath *string `json:"profile_path,omitempty"`
	} `json:"created_by,omitempty"`
	EpisodeRunTime *[]int  `json:"episode_run_time,omitempty"`
	FirstAirDate   *string `json:"first_air_date,omitempty"`
	Genres         *[]struct {
		Id   *int    `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	} `json:"genres,omitempty"`
	Homepage         *string   `json:"homepage,omitempty"`
	Id               *int      `json:"id,omitempty"`
	InProduction     *bool     `json:"in_production,omitempty"`
	Languages        *[]string `json:"languages,omitempty"`
	LastAirDate      *string   `json:"last_air_date,omitempty"`
	LastEpisodeToAir *struct {
		AirDate        *string  `json:"air_date,omitempty"`
		EpisodeNumber  *int     `json:"episode_number,omitempty"`
		Id             *int     `json:"id,omitempty"`
		Name           *string  `json:"name,omitempty"`
		Overview       *string  `json:"overview,omitempty"`
		ProductionCode *string  `json:"production_code,omitempty"`
		Runtime        *int     `json:"runtime,omitempty"`
		SeasonNumber   *int     `json:"season_number,omitempty"`
		ShowId         *int     `json:"show_id,omitempty"`
		StillPath      *string  `json:"still_path,omitempty"`
		VoteAverage    *float32 `json:"vote_average,omitempty"`
		VoteCount      *int     `json:"vote_count,omitempty"`
	} `json:"last_episode_to_air,omitempty"`
	Name     *string `json:"name,omitempty"`
	Networks *[]struct {
		Id            *int    `json:"id,omitempty"`
		LogoPath      *string `json:"logo_path,omitempty"`
		Name          *string `json:"name,omitempty"`
		OriginCountry *string `json:"origin_country,omitempty"`
	} `json:"networks,omitempty"`
	NextEpisodeToAir    *interface{} `json:"next_episode_to_air,omitempty"`
	NumberOfEpisodes    *int         `json:"number_of_episodes,omitempty"`
	NumberOfSeasons     *int         `json:"number_of_seasons,omitempty"`
	OriginCountry       *[]string    `json:"origin_country,omitempty"`
	OriginalLanguage    *string      `json:"original_language,omitempty"`
	OriginalName        *string      `json:"original_name,omitempty"`
	Overview            *string      `json:"overview,omitempty"`
	Popularity          *float32     `json:"popularity,omitempty"`
	PosterPath          *string      `json:"poster_path,omitempty"`
	ProductionCompanies *[]struct {
		Id            *int    `json:"id,omitempty"`
		LogoPath      *string `json:"logo_path,omitempty"`
		Name          *string `json:"name,omitempty"`
		OriginCountry *string `json:"origin_country,omitempty"`
	} `json:"production_companies,omitempty"`
	ProductionCountries *[]struct {
		Iso31661 *string `json:"iso_3166_1,omitempty"`
		Name     *string `json:"name,omitempty"`
	} `json:"production_countries,omitempty"`
	Seasons *[]struct {
		AirDate      *string `json:"air_date,omitempty"`
		EpisodeCount *int    `json:"episode_count,omitempty"`
		Id           *int    `json:"id,omitempty"`
		Name         *string `json:"name,omitempty"`
		Overview     *string `json:"overview,omitempty"`
		PosterPath   *string `json:"poster_path,omitempty"`
		SeasonNumber *int    `json:"season_number,omitempty"`
		VoteAverage  *int    `json:"vote_average,omitempty"`
	} `json:"seasons,omitempty"`
	SpokenLanguages *[]struct {
		EnglishName *string `json:"english_name,omitempty"`
		Iso6391     *string `json:"iso_639_1,omitempty"`
		Name        *string `json:"name,omitempty"`
	} `json:"spoken_languages,omitempty"`
	Status      *string  `json:"status,omitempty"`
	Tagline     *string  `json:"tagline,omitempty"`
	Type        *string  `json:"type,omitempty"`
	VoteAverage *float32 `json:"vote_average,omitempty"`
	VoteCount   *int     `json:"vote_count,omitempty"`
}

func (response TvSeriesDetails200JSONResponse) VisitTvSeriesDetailsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Validate Key
	// (GET /3/authentication)
	AuthenticationValidateKey(ctx context.Context, request AuthenticationValidateKeyRequestObject) (AuthenticationValidateKeyResponseObject, error)
	// Find By ID
	// (GET /3/find/{external_id})
	FindById(ctx context.Context, request FindByIdRequestObject) (FindByIdResponseObject, error)
	// Details
	// (GET /3/movie/{movie_id})
	MovieDetails(ctx context.Context, request MovieDetailsRequestObject) (MovieDetailsResponseObject, error)
	// Movie
	// (GET /3/search/movie)
	SearchMovie(ctx context.Context, request SearchMovieRequestObject) (SearchMovieResponseObject, error)
	// TV
	// (GET /3/search/tv)
	SearchTv(ctx context.Context, request SearchTvRequestObject) (SearchTvResponseObject, error)
	// Details
	// (GET /3/tv/{series_id})
	TvSeriesDetails(ctx context.Context, request TvSeriesDetailsRequestObject) (TvSeriesDetailsResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// AuthenticationValidateKey operation middleware
func (sh *strictHandler) AuthenticationValidateKey(ctx *gin.Context) {
	var request AuthenticationValidateKeyRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AuthenticationValidateKey(ctx, request.(AuthenticationValidateKeyRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AuthenticationValidateKey")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(AuthenticationValidateKeyResponseObject); ok {
		if err := validResponse.VisitAuthenticationValidateKeyResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// FindById operation middleware
func (sh *strictHandler) FindById(ctx *gin.Context, externalId string, params FindByIdParams) {
	var request FindByIdRequestObject

	request.ExternalId = externalId
	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.FindById(ctx, request.(FindByIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "FindById")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(FindByIdResponseObject); ok {
		if err := validResponse.VisitFindByIdResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// MovieDetails operation middleware
func (sh *strictHandler) MovieDetails(ctx *gin.Context, movieId int32, params MovieDetailsParams) {
	var request MovieDetailsRequestObject

	request.MovieId = movieId
	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.MovieDetails(ctx, request.(MovieDetailsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "MovieDetails")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(MovieDetailsResponseObject); ok {
		if err := validResponse.VisitMovieDetailsResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// SearchMovie operation middleware
func (sh *strictHandler) SearchMovie(ctx *gin.Context, params SearchMovieParams) {
	var request SearchMovieRequestObject

	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.SearchMovie(ctx, request.(SearchMovieRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "SearchMovie")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(SearchMovieResponseObject); ok {
		if err := validResponse.VisitSearchMovieResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// SearchTv operation middleware
func (sh *strictHandler) SearchTv(ctx *gin.Context, params SearchTvParams) {
	var request SearchTvRequestObject

	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.SearchTv(ctx, request.(SearchTvRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "SearchTv")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(SearchTvResponseObject); ok {
		if err := validResponse.VisitSearchTvResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// TvSeriesDetails operation middleware
func (sh *strictHandler) TvSeriesDetails(ctx *gin.Context, seriesId int32, params TvSeriesDetailsParams) {
	var request TvSeriesDetailsRequestObject

	request.SeriesId = seriesId
	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.TvSeriesDetails(ctx, request.(TvSeriesDetailsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "TvSeriesDetails")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(TvSeriesDetailsResponseObject); ok {
		if err := validResponse.VisitTvSeriesDetailsResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
