package httpclient

import (
	"context"
	"fmt"
	"net/http"
)

type HTTPClient struct {
	httpclient *http.Client
}

func NewHTTPClient(httpclient *http.Client) *HTTPClient {
	return &HTTPClient{
		httpclient: httpclient,
	}
}

func (hc *HTTPClient) Get(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create GET request: %w", err)
	}

	resp, err := hc.httpclient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("exec GET request: %w", err)
	}

	return resp, nil
}
