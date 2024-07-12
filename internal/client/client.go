package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/rraymondgh/arr-interface/internal/gql/gqlmodel/gen"
	"github.com/rraymondgh/arr-interface/internal/gqlclient"
	"github.com/rraymondgh/arr-interface/internal/protocol"
	"github.com/rraymondgh/arr-interface/internal/servarr"
)

type AddInfoHashesRequest struct {
	ClientID   string
	InfoHashes []protocol.ID
}

type clientWorker interface {
	AddInfoHashes(ctx context.Context, req AddInfoHashesRequest) error
	downloadOne(ctx context.Context, content *Content, clientId string) error
}

type Content = gqlclient.TorrentContentTorrentContentByIDIdTorrentContentSearchResultItemsTorrentContent

type commonClient struct {
	config *servarr.Config
	Client clientWorker
}

func New(clientID *gen.ClientID, cfg *servarr.Config) (*commonClient, error) {
	cc := commonClient{
		config: cfg,
	}

	switch *clientID {
	case gen.ClientIDServarr:
		cc.Client = &servarrClient{commonClient: cc}
	case gen.ClientIDTransmission:
		cc.Client = &transmissionClient{commonClient: cc}
	case gen.ClientIDQBittorrent:
		cc.Client = &qBitClient{commonClient: cc}
	default:
		return nil, fmt.Errorf("not implemented %s", clientID)
	}
	return &cc, nil

}

type downloadOne func(ctx context.Context, content *Content, clientId string) error

func AddInfoHashes(ctx context.Context, req AddInfoHashesRequest, worker downloadOne, cfg *servarr.Config) error {
	for _, infoHash := range req.InfoHashes {
		cr, err := gqlclient.TorrentContent(ctx, graphql.NewClient(fmt.Sprintf("%v/graphql", cfg.BitmagnetUrl), http.DefaultClient), infoHash)
		if err != nil {
			return err
		} else if len(cr.TorrentContentByID.Id.Items) != 1 {
			return fmt.Errorf("too many content results (%d) for download", len(cr.TorrentContentByID.Id.Items))
		}

		err = worker(ctx, &cr.TorrentContentByID.Id.Items[0], req.ClientID)
		if err != nil {
			return err
		}

	}
	return nil
}
