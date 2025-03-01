// Package tmdbproxyoapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package tmdbproxyoapi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

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
