package blizzard

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

type TokenProvider struct {
	clientID     string
	clientSecret string
	token        string
	cachePath    string

	expiresAt  time.Time
	mu         sync.Mutex
	httpClient *http.Client
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"` // in seconds
	TokenType   string `json:"token_type"` // usually "bearer"
}

type tokenFileData struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

func NewTokenProvider() *TokenProvider {
	return &TokenProvider{
		clientID:     os.Getenv("BLIZZARD_CLIENT_ID"),
		clientSecret: os.Getenv("BLIZZARD_CLIENT_SECRET"),
		cachePath:    "./data/.blizzard_token.json",
		httpClient:   &http.Client{Timeout: 10 * time.Second},
	}
}

func (p *TokenProvider) GetToken(ctx context.Context) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if time.Now().Before(p.expiresAt) && p.token != "" {
		return p.token, nil
	}

	// Fetch a new token
	form := url.Values{}
	form.Set("grant_type", "client_credentials")

	req, err := http.NewRequestWithContext(ctx, "POST", "https://us.battle.net/oauth/token", strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("build request failed: %w", err)
	}

	req.SetBasicAuth(p.clientID, p.clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("token request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("token request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var tr tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return "", fmt.Errorf("failed to decode token response: %w", err)
	}

	p.token = tr.AccessToken
	p.expiresAt = time.Now().Add(time.Duration(tr.ExpiresIn-60) * time.Second)

	_ = p.saveToFile()

	return p.token, nil
}

func (p *TokenProvider) loadFromFile() error {
	data, err := os.ReadFile(p.cachePath)
	if err != nil {
		return err
	}
	var tf tokenFileData

	if err := json.Unmarshal(data, &tf); err != nil {
		return err
	}

	if time.Now().Before(tf.ExpiresAt) {
		p.token = tf.AccessToken
		p.expiresAt = tf.ExpiresAt
	}
	return nil
}

func (p *TokenProvider) saveToFile() error {
	tf := tokenFileData{
		AccessToken: p.token,
		ExpiresAt:   p.expiresAt,
	}
	data, err := json.MarshalIndent(tf, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(p.cachePath, data, 0600)
}
