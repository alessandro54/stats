package response

type PvpSeasonIndexResponse struct {
	Seasons       []PvpSeasonIndexSeason      `json:"seasons"`
	CurrentSeason PvpSeasonIndexCurrentSeason `json:"current_season"`
}

type PvpSeasonIndexSeason struct {
	ID int `json:"id"`
}

type PvpSeasonIndexCurrentSeason struct {
	ID int `json:"id"`
}
