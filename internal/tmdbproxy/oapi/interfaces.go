package tmdbproxyoapi

import "time"

// oapi strict generates some incorrect and nested structures
// pastes them here for simpler use and correct

type AuthenticationValidateKey200JSONResponse struct {
	StatusCode    int    `json:"status_code,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`
	Success       bool   `json:"success,omitempty"`
}

type intArray []int

type rankingScoreDetails struct {
	Words struct {
		Order            int     `json:"order"`
		MatchingWords    int     `json:"matchingWords"`
		MaxMatchingWords int     `json:"maxMatchingWords"`
		Score            float64 `json:"score"`
	} `json:"words"`
	Typo struct {
		Order        int     `json:"order"`
		TypoCount    int     `json:"typoCount"`
		MaxTypoCount int     `json:"maxTypoCount"`
		Score        float64 `json:"score"`
	} `json:"typo"`
	Proximity struct {
		Order int     `json:"order"`
		Score float64 `json:"score"`
	} `json:"proximity"`
	Attribute struct {
		Order                      int     `json:"order"`
		AttributeRankingOrderScore float64 `json:"attributeRankingOrderScore"`
		QueryWordDistanceScore     float64 `json:"queryWordDistanceScore"`
		Score                      float64 `json:"score"`
	} `json:"attribute"`
	Exactness struct {
		Order            int     `json:"order"`
		MatchType        string  `json:"matchType"`
		MatchingWords    *int    `json:"matchingWords,omitempty"`
		MaxMatchingWords *int    `json:"maxMatchingWords,omitempty"`
		Score            float64 `json:"score"`
	} `json:"exactness"`
}

type FindBy struct {
	Adult               bool                `json:"adult"`
	BackdropPath        *string             `json:"backdrop_path,omitempty"`
	GenreIds            intArray            `json:"genre_ids,omitempty"`
	Id                  int64               `json:"id,omitempty"`
	MediaType           string              `json:"media_type,omitempty"`
	OriginalLanguage    *string             `json:"original_language,omitempty"`
	OriginalTitle       *string             `json:"original_title,omitempty"`
	Overview            *string             `json:"overview,omitempty"`
	Popularity          *float64            `json:"popularity,omitempty"`
	PosterPath          *string             `json:"poster_path,omitempty"`
	ReleaseDate         *string             `json:"release_date,omitempty"`
	Title               *string             `json:"title,omitempty"`
	Video               *bool               `json:"video,omitempty"`
	VoteAverage         *float64            `json:"vote_average,omitempty"`
	VoteCount           *int                `json:"vote_count,omitempty"`
	Name                *string             `json:"name,omitempty"`
	OriginalName        *string             `json:"original_name,omitempty"`
	FirstAirDate        *string             `json:"first_air_date,omitempty"`
	RankingScore        float64             `json:"_rankingScore"`
	RankingScoreDetails rankingScoreDetails `json:"_rankingScoreDetails"`
}

type FindById200JSONResponse struct {
	MovieResults     []*FindBy      `json:"movie_results"`
	PersonResults    *[]interface{} `json:"person_results,omitempty"`
	TvEpisodeResults *[]interface{} `json:"tv_episode_results,omitempty"`
	TvResults        []*FindBy      `json:"tv_results"`
	TvSeasonResults  *[]interface{} `json:"tv_season_results,omitempty"`
}

type idName struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type idNameArray []idName

type MovieDetails struct {
	Adult               *bool        `json:"adult,omitempty"`
	BackdropPath        *string      `json:"backdrop_path,omitempty"`
	BelongsToCollection *idName      `json:"belongs_to_collection,omitempty"`
	Budget              *int         `json:"budget,omitempty"`
	Genres              idNameArray  `json:"genres,omitempty"`
	Homepage            *string      `json:"homepage,omitempty"`
	Id                  int          `json:"id,omitempty"`
	ImdbId              *string      `json:"imdb_id,omitempty"`
	OriginalLanguage    *string      `json:"original_language,omitempty"`
	OriginalTitle       *string      `json:"original_title,omitempty"`
	Overview            *string      `json:"overview,omitempty"`
	Popularity          *float64     `json:"popularity,omitempty"`
	PosterPath          *string      `json:"poster_path,omitempty"`
	ProductionCompanies *idNameArray `json:"production_companies,omitempty"`
	ReleaseDate         *string      `json:"release_date,omitempty"`
	Revenue             *int         `json:"revenue,omitempty"`
	Runtime             *int         `json:"runtime,omitempty"`
	Status              *string      `json:"status,omitempty"`
	Tagline             *string      `json:"tagline,omitempty"`
	Title               *string      `json:"title,omitempty"`
	Video               *bool        `json:"video,omitempty"`
	VoteAverage         *float64     `json:"vote_average,omitempty"`
	VoteCount           *int         `json:"vote_count,omitempty"`
	ProductionCountries *[]struct {
		Iso31661 *string `json:"iso_3166_1,omitempty"`
		Name     *string `json:"name,omitempty"`
	} `json:"production_countries,omitempty"`
	SpokenLanguages *[]struct {
		EnglishName *string `json:"english_name,omitempty"`
		Iso6391     *string `json:"iso_639_1,omitempty"`
		Name        *string `json:"name,omitempty"`
	} `json:"spoken_languages,omitempty"`
	ExternalIDs ExternalIdsResponse `json:"external_ids,omitempty"`

	// needed by meili / proxy
	MediaType  *string    `json:"media_type,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	PrimaryKey string     `json:"pk"`
}

type Search200JSONResponse struct {
	Page         int       `json:"page,omitempty"`
	Results      []*FindBy `json:"results,omitempty"`
	TotalPages   int       `json:"total_pages,omitempty"`
	TotalResults int       `json:"total_results,omitempty"`
}

type SearchTv200JSONResponse struct {
	Page         int       `json:"page,omitempty"`
	Results      []*FindBy `json:"results,omitempty"`
	TotalPages   int       `json:"total_pages,omitempty"`
	TotalResults int       `json:"total_results,omitempty"`
}

type TvSeriesDetails struct {
	Adult        *bool   `json:"adult,omitempty"`
	BackdropPath *string `json:"backdrop_path,omitempty"`
	// CreatedBy    *[]struct {
	// 	CreditId    *string `json:"credit_id,omitempty"`
	// 	Gender      *int    `json:"gender,omitempty"`
	// 	Id          *int    `json:"id,omitempty"`
	// 	Name        *string `json:"name,omitempty"`
	// 	ProfilePath *string `json:"profile_path,omitempty"`
	// } `json:"created_by,omitempty"`
	EpisodeRunTime *intArray   `json:"episode_run_time,omitempty"`
	FirstAirDate   *string     `json:"first_air_date,omitempty"`
	Genres         idNameArray `json:"genres,omitempty"`
	Homepage       *string     `json:"homepage,omitempty"`
	Id             int         `json:"id,omitempty"`
	InProduction   *bool       `json:"in_production,omitempty"`
	Languages      *[]string   `json:"languages,omitempty"`
	LastAirDate    *string     `json:"last_air_date,omitempty"`
	// LastEpisodeToAir *episode `json:"last_episode_to_air,omitempty"`
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
	Popularity          *float64     `json:"popularity,omitempty"`
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
	// Seasons seriesSeasonArray `json:"seasons,omitempty"`
	Seasons *[]struct {
		AirDate      *string  `json:"air_date,omitempty"`
		EpisodeCount *int     `json:"episode_count,omitempty"`
		Id           *int     `json:"id,omitempty"`
		Name         *string  `json:"name,omitempty"`
		Overview     *string  `json:"overview,omitempty"`
		PosterPath   *string  `json:"poster_path,omitempty"`
		SeasonNumber *int     `json:"season_number,omitempty"`
		VoteAverage  *float64 `json:"vote_average,omitempty"`
	} `json:"seasons,omitempty"`
	SpokenLanguages *[]struct {
		EnglishName *string `json:"english_name,omitempty"`
		Iso6391     *string `json:"iso_639_1,omitempty"`
		Name        *string `json:"name,omitempty"`
	} `json:"spoken_languages,omitempty"`
	Status      *string             `json:"status,omitempty"`
	Tagline     *string             `json:"tagline,omitempty"`
	Type        *string             `json:"type,omitempty"`
	VoteAverage *float64            `json:"vote_average,omitempty"`
	VoteCount   *int                `json:"vote_count,omitempty"`
	ExternalIDs ExternalIdsResponse `json:"external_ids,omitempty"`

	// needed by meili / proxy
	MediaType  *string    `json:"media_type,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	PrimaryKey string     `json:"pk"`
}

type ExternalIdsResponse struct {
	FacebookId  *string `json:"facebook_id,omitempty"`
	FreebaseId  *string `json:"freebase_id,omitempty"`
	FreebaseMid *string `json:"freebase_mid,omitempty"`
	Id          *int    `json:"id,omitempty"`
	ImdbId      *string `json:"imdb_id,omitempty"`
	InstagramId *string `json:"instagram_id,omitempty"`
	TvdbId      *int    `json:"tvdb_id,omitempty"`
	TvrageId    *int    `json:"tvrage_id,omitempty"`
	TwitterId   *string `json:"twitter_id,omitempty"`
	WikidataId  *string `json:"wikidata_id,omitempty"`
	// needed by meili / proxy
	MediaType  *string    `json:"media_type,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	PrimaryKey string     `json:"pk"`
}
