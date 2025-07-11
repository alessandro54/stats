package blizzard

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchPvpLeaderboard(ctx context.Context, seasonId string, bracket string) ([]byte, error) {
	path := fmt.Sprintf("/data/wow/pvp-season/%s/pvp-leaderboard/%s", seasonId, bracket)
	namespace := "dynamic-" + c.Region
	locale := "en_US"

	url := fmt.Sprintf("%s%s?namespace=%s&locale=%s", c.BaseURL, path, namespace, locale)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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
