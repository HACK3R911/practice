package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/HACK3R911/practice/downloader/internal/config"
)

func ParseConfig(args []string) (config.Config, error) {
	flags := flag.NewFlagSet("downloader", flag.ContinueOnError)

	outputValue := flags.String(
		"output",
		"",
		"path where the downloaded file will be saved",
	)

	workersValue := flags.Int(
		"workers",
		config.DefaultWorkers,
		"number of concurrent download workers",
	)

	chunkSizeValue := flags.Int64(
		"chunk-size",
		config.DefaultChunkSize,
		"chunk size in bytes",
	)

	//urlValue := flags.String(
	//	"url",
	//	"",
	//	"url of the file to download",
	//)

	flags.SetOutput(os.Stderr)

	flags.Usage = func() {
		fmt.Fprintln(flags.Output(), "Надёжный загрузчик файлов")
		fmt.Fprintln(flags.Output())
		fmt.Fprintln(flags.Output(), "Использование:")
		fmt.Fprintln(flags.Output())
		fmt.Fprintln(flags.Output(), "  downloader [опции] URL [URL...]")
		fmt.Fprintln(flags.Output())
		fmt.Fprintln(flags.Output(), "Опции:")
		flags.PrintDefaults()
		fmt.Fprintln(flags.Output())
		fmt.Fprintln(flags.Output(), "Примеры:")
		fmt.Fprintln(flags.Output())
		fmt.Fprintln(
			flags.Output(),
			"  downloader https://example.com/file.zip",
		)
		fmt.Fprintln(
			flags.Output(),
			"  downloader --output ./downloads https://example.com/file.zip",
		)
		fmt.Fprintln(
			flags.Output(),
			"  downloader --workers 8 --chunk-size 1048576 \\",
		)
		fmt.Fprintln(
			flags.Output(),
			"    https://example.com/file1.zip \\",
		)
		fmt.Fprintln(
			flags.Output(),
			"    https://example.com/file2.zip",
		)
	}

	if err := flags.Parse(args); err != nil {
		return config.Config{}, err
	}

	cfg := config.Config{
		Output:    *outputValue,
		Workers:   *workersValue,
		ChunkSize: *chunkSizeValue,
		URLs:      flags.Args(),
	}

	if err := cfg.Validate(); err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}
