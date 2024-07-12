package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hekmon/transmissionrpc/v3"
	"github.com/rraymondgh/arr-interface/internal/servarr"
)

type transmissionClient struct {
	commonClient
}

func (c *transmissionClient) downloadOne(ctx context.Context, content *Content, clientId string) error {

	d, err := servarr.New(c.config, content)
	if err != nil {
		return err
	}

	client, err := servarr.GetDownloadClient(d.Arr, ctx, clientId)
	if err != nil {
		return err
	}

	endpoint, err := url.Parse(fmt.Sprintf("http://%v%v:%v/transmission/rpc", client.Host, c.config.DownloadClientDomain, client.Port))
	if err != nil {
		return err
	}
	tbt, err := transmissionrpc.New(endpoint, nil)
	if err != nil {
		return err
	}

	settings, err := tbt.SessionArgumentsGetAll(ctx)
	if err != nil {
		return err
	}
	dir := *settings.DownloadDir + "/" + client.Category

	_, err = tbt.TorrentAdd(ctx, transmissionrpc.TorrentAddPayload{
		Filename:    &content.Torrent.MagnetUri,
		DownloadDir: &dir,
	})
	return err

}

func (c *transmissionClient) AddInfoHashes(ctx context.Context, req AddInfoHashesRequest) error {
	return AddInfoHashes(ctx, req, c.downloadOne, c.config)

}
