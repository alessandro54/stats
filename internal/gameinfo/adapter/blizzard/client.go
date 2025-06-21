package blizzard

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

type Client struct {
	Region     string
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

var (
	clientInstance *Client
	once           sync.Once
)

func GetClient() *Client {
	once.Do(func() {
		token, err := NewTokenProvider().GetToken(context.Background())
		if err != nil {
			panic(fmt.Sprintf("failed to get Blizzard token: %v", err))
		}

		clientInstance = &Client{
			Region:     "us",
			BaseURL:    fmt.Sprintf("https://%s.api.blizzard.com", "us"),
			Token:      token,
			HTTPClient: &http.Client{},
		}
	})
	return clientInstance
}

func (c *Client) Get(ctx context.Context, path string, queryParams map[string]string) ([]byte, error) {
	baseURL := fmt.Sprintf("%s%s", c.BaseURL, path)

	q := url.Values{}
	for key, values := range queryParams {
		q.Add(key, values)
	}
	fullURL := fmt.Sprintf("%s?%s", baseURL, q.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	return body, nil
}
