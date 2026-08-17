package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	DefaultWorkers   = 2
	DefaultChunkSize = 10 * 1024 * 1024 // 10 MiB
)

type Config struct {
	Output    string
	Workers   int
	ChunkSize int64
	URLs      []string
}

func (c *Config) Validate() error {

	if err := validateOutput(c.Output); err != nil {
		return err
	}

	if c.Workers <= 0 {
		return errors.New("workers must be greater than zero")
	}

	if c.Workers > 128 {
		return errors.New("workers must not be greater than 128")
	}

	if c.ChunkSize <= 0 {
		return errors.New("chunk-size must be greater than zero")
	}

	if len(c.URLs) == 0 {
		return errors.New("необходимо указать хотя бы один URL")
	}

	for _, rawUrl := range c.URLs {
		if err := validateURL(rawUrl); err != nil {
			return err
		}
	}

	return nil
}

func validateURL(rawUrl string) error {
	if strings.TrimSpace(rawUrl) == "" {
		return errors.New("url is empty")
	}

	u, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return err
	}

	switch strings.ToLower(u.Scheme) {
	case "http", "https":
	default:
		return fmt.Errorf("unsupported URL scheme %q: use http or https", u.Scheme)
	}

	if u.Host == "" {
		return errors.New("URL must contain a host")
	}

	return nil
}

func validateOutput(output string) error {
	if strings.TrimSpace(output) == "" {
		return errors.New("output is empty")
	}

	if filepath.IsAbs(output) {
		return nil
	}

	cleaned := filepath.Clean(output)

	if cleaned == "." || cleaned == string(os.PathSeparator) {
		return errors.New("output path must point to a file")
	}

	return nil
}
