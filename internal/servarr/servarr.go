package servarr

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/rraymondgh/arr-interface/internal/gqlclient"
	"golift.io/starr/prowlarr"
	"golift.io/starr/radarr"
	"golift.io/starr/sonarr"
)

type Content = gqlclient.TorrentContentTorrentContentByIDIdTorrentContentSearchResultItemsTorrentContent
type attribute = gqlclient.TorrentContentTorrentContentByIDIdTorrentContentSearchResultItemsTorrentContentContentAttributesContentAttribute
type ArrIndexer = sonarr.IndexerOutput
type ArrRelease = prowlarr.Search
type ArrClient = prowlarr.DownloadClientOutput

type ServarrSpecific interface {
	GetIndexers(ctx context.Context) ([]*ArrIndexer, error)
	UpdateIndexers(ctx context.Context, indexerIds []int64, enabled bool) error
	SearchRelease(ctx context.Context, content *Content, indexer *ArrIndexer) ([]*ArrRelease, error)
	GrabRelease(ctx context.Context, indexer *ArrIndexer, release *ArrRelease) error
	GetDownloadClients(ctx context.Context) ([]*ArrClient, error)
}

type ServarrInstance struct {
	Arr         ServarrSpecific
	IndexerName string
	Prowlarr    *prowlarr.Prowlarr
	Sonarr      *sonarr.Sonarr
	Radarr      *radarr.Radarr
}

type DownloadClient struct {
	Port     string
	Host     string
	Category string
	Username string
	Password string
}

func GetIndexers(arr ServarrSpecific, ctx context.Context, indexerName string) ([]*ArrIndexer, *ArrIndexer, error) {
	indexers, err := arr.GetIndexers(ctx)
	if err != nil {
		return nil, nil, err
	}

	i := slices.IndexFunc(indexers, func(i *ArrIndexer) bool { return i.Name == indexerName })
	if i == -1 {
		return nil, nil, fmt.Errorf("indexer not found %s", indexerName)
	}

	return indexers, indexers[i], err
}

func GetDownloadClient(arr ServarrSpecific, ctx context.Context, clientId string) (*DownloadClient, error) {
	clients, err := arr.GetDownloadClients(ctx)
	if err != nil {
		return nil, err
	}
	i := slices.IndexFunc(clients, func(c *ArrClient) bool { return strings.EqualFold(c.ImplementationName, clientId) })
	if i == -1 {
		return &DownloadClient{}, fmt.Errorf("failed to locate %v", clientId)
	}
	client := DownloadClient{}
	for _, field := range clients[i].Fields {
		switch field.Name {
		case "port":
			client.Port = fmt.Sprintf("%v", field.Value)
		case "host":
			client.Host = fmt.Sprintf("%v", field.Value)
		case "category", "tvCategory", "movieCategory":
			client.Category = fmt.Sprintf("%v", field.Value)
		case "username":
			client.Username = fmt.Sprintf("%v", field.Value)
		case "password":
			client.Password = fmt.Sprintf("%v", field.Value)
		}
	}
	return &client, nil
}

func New(cfg *Config, content *Content) (*ServarrInstance, error) {
	d := NewSonarr(content, cfg)
	if d == nil {
		d = NewRadarr(content, cfg)
	}
	if d == nil {
		d = NewProwlarr(content, cfg)
	}
	if d == nil {
		msg := fmt.Sprintf("no servarr client accepts (%s) (%s)", content.ContentType, content.InfoHash)
		return &ServarrInstance{}, errors.New(msg)
	}

	return d, nil

}
