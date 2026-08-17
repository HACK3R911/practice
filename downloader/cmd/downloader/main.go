package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/HACK3R911/practice/downloader/internal/cli"
	"github.com/HACK3R911/practice/downloader/internal/downloader"
	"github.com/HACK3R911/practice/downloader/internal/httpclient"
)

func main() {
	cfg, err := cli.ParseConfig(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n\n", err)

		os.Exit(1)
	}

	fmt.Println("download configuration:")
	fmt.Printf("  output:     %s\n", cfg.Output)
	output := cfg.Output
	fmt.Printf("  workers:    %d\n", cfg.Workers)
	fmt.Printf("  chunk size: %d bytes\n", cfg.ChunkSize)
	fmt.Printf("  urls:       %s\n", len(cfg.URLs))
	fmt.Println()

	client := httpclient.NewHTTPClient(&http.Client{})
	fileDownloader := downloader.NewDownloader(client)

	ctx := context.Background()

	for _, url := range cfg.URLs {
		fmt.Printf("Скачивание: %s\n", url)

		if err := fileDownloader.Download(ctx, url, output); err != nil {
			fmt.Fprintf(
				os.Stderr,
				"ошибка загрузки %q: %v\n",
				url,
				err,
			)
			os.Exit(1)
		}
		fmt.Printf("Готово: %s\n\n", output)
	}
}
