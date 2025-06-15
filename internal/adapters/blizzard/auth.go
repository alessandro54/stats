package blizzard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/alessandro54/stats/internal/domain"
)

const tokenFilePath = "data/token.json"

type BlizzardTokenProvider struct {
	clientID     string
	clientSecret string
	token        *domain.Token
	mu           sync.Mutex
	client       *http.Client
}

func NewBlizzardTokenProvider() *BlizzardTokenProvider {
	clientID := os.Getenv("BLIZZARD_CLIENT_ID")
	clientSecret := os.Getenv("BLIZZARD_CLIENT_SECRET")

	provider := &BlizzardTokenProvider{
		clientID:     clientID,
		clientSecret: clientSecret,
		client:       &http.Client{Timeout: 10 * time.Second},
	}

	if tokenData, err := os.ReadFile(tokenFilePath); err == nil {
		var token domain.Token
		if err := json.Unmarshal(tokenData, &token); err == nil && !token.IsExpired() {
			provider.token = &token
		}
	}

	return provider
}

func (b *BlizzardTokenProvider) fetchNewToken() (*domain.Token, error) {
	req, err := http.NewRequest("POST", "https://oauth.battle.net/token", nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(b.clientID, b.clientSecret)
	q := req.URL.Query()
	q.Add("grant_type", "client_credentials")
	req.URL.RawQuery = q.Encode()

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token request failed: %s", resp.Status)
	}

	var tok domain.Token
	if err := json.NewDecoder(resp.Body).Decode(&tok); err != nil {
		return nil, err
	}

	tok.FetchedAt = time.Now()

	// Save token to disk
	_ = os.MkdirAll("data", os.ModePerm)
	if tokenBytes, err := json.MarshalIndent(tok, "", "  "); err == nil {
		_ = os.WriteFile(tokenFilePath, tokenBytes, 0644)
	}

	tok.FetchedAt = time.Now()
	return &tok, nil
}

func (b *BlizzardTokenProvider) GetToken() (string, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.token != nil && !b.token.IsExpired() {
		return b.token.AccessToken, nil
	}

	token, err := b.fetchNewToken()
	if err != nil {
		return "", err
	}

	b.token = token
	return token.AccessToken, nil
}
