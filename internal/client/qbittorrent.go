package client

import (
	"context"
	"fmt"

	"github.com/autobrr/go-qbittorrent"
	"github.com/rraymondgh/arr-interface/internal/servarr"
)

type qBitClient struct {
	commonClient
}

func (c *qBitClient) downloadOne(ctx context.Context, content *Content, clientId string) error {
	d, err := servarr.New(c.config, content)
	if err != nil {
		return err
	}

	client, err := servarr.GetDownloadClient(d.Arr, ctx, clientId)
	if err != nil {
		return err
	}

	qb := qbittorrent.NewClient(qbittorrent.Config{
		Host:     fmt.Sprintf("http://%v%v:%v/", client.Host, c.config.DownloadClientDomain, client.Port),
		Username: client.Username,
		Password: c.config.Qbittorrent.Password,
	})

	err = qb.LoginCtx(ctx)
	if err != nil {
		return err
	}

	pref, err := qb.GetAppPreferencesCtx(ctx)
	if err != nil {
		return err
	}

	err = qb.AddTorrentFromUrlCtx(
		ctx,
		content.Torrent.MagnetUri,
		map[string]string{
			"savepath": fmt.Sprintf("%v/%v", pref.SavePath, client.Category),
			"category": client.Category,
		},
	)

	return err

}

func (c *qBitClient) AddInfoHashes(ctx context.Context, req AddInfoHashesRequest) error {
	return AddInfoHashes(ctx, req, c.downloadOne, c.config)

}
