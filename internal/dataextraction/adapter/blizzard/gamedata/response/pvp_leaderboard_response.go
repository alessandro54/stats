package response

type PvpLeaderboardResponse struct {
	Entries []PvpEntry `json:"entries"`
	Season  Season     `json:"season"`
	Name    string
}

type Season struct {
	ID uint `json:"id"`
}

type PvpEntry struct {
	Character             Character             `json:"character"`
	Faction               Faction               `json:"faction"`
	Rank                  uint                  `json:"rank"`
	Rating                uint                  `json:"rating"`
	SeasonMatchStatistics SeasonMatchStatistics `json:"season_match_statistics"`
	Tier                  Tier                  `json:"tier"`
}

type Character struct {
	Name  string `json:"name"`
	ID    uint   `json:"id"`
	Realm Realm  `json:"realm"`
}

type Realm struct {
	Key  Link   `json:"key"`
	ID   uint   `json:"id"`
	Slug string `json:"slug"`
}

type Link struct {
	Href string `json:"href"`
}

type Faction struct {
	Type string `json:"type"`
}

type SeasonMatchStatistics struct {
	Played uint `json:"played"`
	Won    uint `json:"won"`
	Lost   uint `json:"lost"`
}

type Tier struct {
	Key Link `json:"key"`
	ID  uint `json:"id"`
}
