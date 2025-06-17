package blizzard

import (
	"context"
	"fmt"
	"net/http"
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
