package main

import (
	"fmt"
	"os"

	"github.com/HACK3R911/practice/downloader/internal/cli"
)

var (
//outputFile = flag.String("out", "downloads/", "Файл для записи результата")

// verbose    = flag.Bool("v", false, "Подробный вывод")
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

	//if len(os.Args) < 3 {
	//	fmt.Println("Использование: downloader <директория> <url1> [url2...]")
	//	os.Exit(1)
	//}
	//
	//savePath := os.Args[1]
	//urls := os.Args[2:]
	//
	//fmt.Println("Директория для сохранения:", savePath)
	//fmt.Println("URL для скачивания:")
	//for _, url := range urls {
	//	fmt.Println("-", url)
	//}
	//
	//fmt.Println(os.Args[0])
}
