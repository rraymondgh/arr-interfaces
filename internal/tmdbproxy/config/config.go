package config

type flagDef struct {
	Required bool
	Calendar string
	All      string
}

type Arr struct {
	ApiKey        string
	Url           string
	GenerateFlags flagDef
}

type Config struct {
	BaseUrl       string
	ApiKey        string
	LogCacheTTL   int
	CompressCache bool
	Sonarr        Arr
	Radarr        Arr
	Meilisearch   meiliconfig
	Tmdb          tmdbconfig
}

type meiliconfig struct {
	Synonyms  map[string][]string
	StopWords []string
}

type tmdbconfig struct {
	FetchMissing   bool
	MinRequests    int
	BackoffMinutes int
}

func NewDefaultConfig() Config {
	syns := make(map[string][]string, 0)
	syns["au"] = []string{"australian"}
	syns["uk"] = []string{"gb"}
	return Config{
		BaseUrl:       "https://api.themoviedb.org/3",
		ApiKey:        defaultTmdbApiKey,
		LogCacheTTL:   120,
		CompressCache: true,
		Sonarr: Arr{
			Url:           "http://localhost:8989",
			GenerateFlags: flagDef{Required: true, Calendar: "sonarr_active", All: "sonarr"}},
		Radarr: Arr{
			Url:           "http://localhost:7878",
			GenerateFlags: flagDef{Required: true, All: "radarr"}},
		Meilisearch: meiliconfig{
			Synonyms:  syns,
			StopWords: []string{"and", "-"},
		},
		Tmdb: tmdbconfig{
			FetchMissing:   true,
			MinRequests:    20,
			BackoffMinutes: 10,
		},
	}
}

const (
	defaultTmdbApiKey = "9c6689fa83ae6814fbfb200d70bba3a8"
)
