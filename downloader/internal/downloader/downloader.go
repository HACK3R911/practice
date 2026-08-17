package downloader

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/HACK3R911/practice/downloader/internal/httpclient"
)

type Downloader struct {
	client *httpclient.HTTPClient
}

func NewDownloader(client *httpclient.HTTPClient) *Downloader {
	return &Downloader{
		client: client,
	}
}

func (d *Downloader) Download(ctx context.Context, url, outputPath string) error {
	resp, err := d.client.Get(ctx, url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("download file: %s", err)
	}

	return nil
}
