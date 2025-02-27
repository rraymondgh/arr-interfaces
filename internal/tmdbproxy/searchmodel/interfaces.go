package searchmodel

import (
	"errors"
	"time"

	"github.com/meilisearch/meilisearch-go"
)

type TmdbStatus = int

const (
	StatusUndefined TmdbStatus = iota
	StatusNotFound
	StatusFound
)

var ErrNoDocuments = errors.New("no documents in result")

type SearchResponse struct {
	*meilisearch.SearchResponse
	error
	all bool
}

type Tmdb struct {
	ID        int
	MediaType string
}

type TmdbSearch struct {
	ID        int
	MediaType string
	Query     string
	Year      int
}

type TmdbExternalId struct {
	ID             int
	MediaType      string
	ExternalId     string
	ExternalSource string
}

type UrlLog struct {
	TmdbId           *int64     `json:"tmdb_id,omitempty"`
	TmdbStatus       TmdbStatus `json:"tmdb_status"`
	LastTmdbAttempt  *time.Time `json:"last_tmdb_attempt,omitempty"`
	Counter          int
	Path             string // path (relative paths may omit leading slash)
	RawQuery         string // encoded query values, without '?'
	ApiKey           string `form:"api_key"`
	Query            string `form:"query"`
	QueryChanged     bool   `json:"query_changed"`
	Year             *int   `form:"year"`
	FirstAirDateYear *int   `form:"first_air_date_year"`
}

type Similarity struct {
	Score   float64 `json:"score"`
	Measure struct {
		OSADamerauLevenshtein float32 `json:"osadameraulevenshtein"`
		Lcs                   float32 `json:"lcs"`
		Cosine                float32 `json:"cosine"`
		Jaccard               float32 `json:"jaccard"`
		SorensenDice          float32 `json:"sorensendice"`
		Qgram                 float32 `json:"qgram"`
	}
	Summary struct {
		Mean   float32 `json:"mean"`
		Median float32 `json:"median"`
		Min    float32 `json:"min"`
		Max    float32 `json:"max"`
	}
	Distance struct {
		Levenshtein int
	}
}
