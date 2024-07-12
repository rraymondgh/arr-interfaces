package servarr

import (
	"encoding/json"

	"golift.io/starr/prowlarr"
	"golift.io/starr/radarr"
	"golift.io/starr/sonarr"
)

func Recast(a, b interface{}) error {
	js, err := json.Marshal(a)
	if err != nil {
		return err
	}
	return json.Unmarshal(js, b)
}

func RecastIndexerSlice[
	S prowlarr.IndexerOutput | radarr.IndexerOutput | ArrIndexer,
	D ArrIndexer](
	src []*S, dest []*D) error {
	for i, idx := range src {
		dest[i] = &D{}
		err := Recast(idx, dest[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func RecastReleaseSlice[
	S sonarr.Release | ArrRelease | radarr.Release,
	D ArrRelease](
	src []*S, dest []*D) error {
	for i, idx := range src {
		dest[i] = &D{}
		err := Recast(idx, dest[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func RecastClientDownloadSlice[
	S ArrClient | sonarr.DownloadClientOutput | radarr.DownloadClientOutput,
	D ArrClient](
	src []*S, dest []*D) error {
	for i, idx := range src {
		dest[i] = &D{}
		err := Recast(idx, dest[i])
		if err != nil {
			return err
		}
	}

	return nil
}
