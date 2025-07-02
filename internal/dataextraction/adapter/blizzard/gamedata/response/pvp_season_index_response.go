package response

type PvpSeasonIndexResponse struct {
	Seasons       []pvpSeasonIndexSeason      `json:"seasons"`
	CurrentSeason pvpSeasonIndexCurrentSeason `json:"current_season"`
}

type pvpSeasonIndexSeason struct {
	ID int `json:"id"`
}

type pvpSeasonIndexCurrentSeason struct {
	ID int `json:"id"`
}
