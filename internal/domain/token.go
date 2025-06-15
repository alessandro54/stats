package domain

import "time"

type Token struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   int       `json:"expires_in"`
	Sub         string    `json:"sub"`
	FetchedAt   time.Time `json:"fetched_at"`
}

func (t *Token) IsExpired() bool {
	return time.Since(t.FetchedAt) > time.Duration(t.ExpiresIn-60)*time.Second
}
