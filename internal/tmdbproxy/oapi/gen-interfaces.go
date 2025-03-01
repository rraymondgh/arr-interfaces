// Package tmdbproxyoapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package tmdbproxyoapi

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
	TvrageId    FindByIdParamsExternalSource = "tvrage_id"
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
