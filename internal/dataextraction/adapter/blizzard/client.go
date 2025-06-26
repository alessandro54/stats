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
	Locale     string
}

type clientKey struct {
	region string
	locale string
}

var (
	clientCache = make(map[clientKey]*Client)
	mu          sync.Mutex
)

func GetClient(ctx context.Context, region string, locale string) (*Client, error) {
	if region == "" {
		region = "us"
	}
	if locale == "" {
		locale = "en_US"
	}

	key := clientKey{region, locale}

	mu.Lock()
	defer mu.Unlock()

	if client, ok := clientCache[key]; ok {
		return client, nil
	}

	token, err := NewTokenProvider().GetToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get Blizzard token: %w", err)
	}

	client := &Client{
		Region:     region,
		BaseURL:    fmt.Sprintf("https://%s.api.blizzard.com", region),
		Token:      token,
		HTTPClient: &http.Client{},
		Locale:     locale,
	}

	clientCache[key] = client
	return client, nil
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
