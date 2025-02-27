package meiliclient

type Config struct {
	Uri       string
	MasterKey string
}

func NewDefaultConfig() Config {
	return Config{
		Uri:       "http://localhost:7700",
		MasterKey: "invalid",
	}
}
