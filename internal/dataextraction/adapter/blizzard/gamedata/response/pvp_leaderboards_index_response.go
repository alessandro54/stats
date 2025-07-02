package response

type PvpLeaderboardsIndexResponse struct {
	Season       pvpLeaderboardsIndexSeasonResponse
	Leaderboards []pvpLeaderboardsIndexLeaderboardResponse
}

type pvpLeaderboardsIndexSeasonResponse struct {
	ID int `json:"id"`
}

type pvpLeaderboardsIndexLeaderboardResponse struct {
	Name string  `json:"name"`
	ID   *string `json:"id"`
}
