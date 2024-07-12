package servarr

type UrlKey struct {
	Url    string
	ApiKey string
}

type ClientCredentials struct {
	Password string
}

type Config struct {
	BitmagnetUrl         string
	IndexerName          string
	OnlySearchBitmagnet  bool
	DownloadClientDomain string
	Sonarr               UrlKey
	Radarr               UrlKey
	Prowlarr             UrlKey
	Qbittorrent          ClientCredentials
}

func NewDefaultConfig() Config {
	return Config{
		BitmagnetUrl:         "http://localhost:3333",
		IndexerName:          "bitmagnet",
		OnlySearchBitmagnet:  false,
		DownloadClientDomain: "",
		Sonarr:               UrlKey{Url: "http://localhost:8989", ApiKey: "private"},
		Radarr:               UrlKey{Url: "http://localhost:7878", ApiKey: "private"},
		Prowlarr:             UrlKey{Url: "http://localhost:9696", ApiKey: "private"},
		Qbittorrent:          ClientCredentials{Password: "default"},
	}
}
