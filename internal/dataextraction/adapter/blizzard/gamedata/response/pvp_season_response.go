package response

type PvpSeasonResponse struct {
	ID                uint   `json:"id"`
	SeasonStartUnixMs int64  `json:"season_start_timestamp"`
	SeasonEndUnixMs   int64  `json:"season_end_timestamp"`
	SeasonName        string `json:"season_name"`
}
