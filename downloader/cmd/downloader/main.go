package main

import (
	"fmt"
	"os"

	"github.com/HACK3R911/practice/downloader/internal/cli"
)

func main() {
	cfg, err := cli.ParseConfig(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n\n", err)

		os.Exit(1)
	}

	fmt.Println("download configuration:")
	fmt.Printf("  output:     %s\n", cfg.Output)
	fmt.Printf("  workers:    %d\n", cfg.Workers)
	fmt.Printf("  chunk size: %d bytes\n", cfg.ChunkSize)
	fmt.Printf("  urls:       %s\n", len(cfg.URLs))
	fmt.Println()
	fmt.Println("download is not implemented yet")
}
