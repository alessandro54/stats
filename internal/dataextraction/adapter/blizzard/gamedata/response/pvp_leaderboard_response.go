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
	Rank                  int                   `json:"rank"`
	Rating                int                   `json:"rating"`
	SeasonMatchStatistics SeasonMatchStatistics `json:"season_match_statistics"`
	Tier                  Tier                  `json:"tier"`
}

type Character struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Realm Realm  `json:"realm"`
}

type Realm struct {
	Key  Link   `json:"key"`
	ID   int    `json:"id"`
	Slug string `json:"slug"`
}

type Link struct {
	Href string `json:"href"`
}

type Faction struct {
	Type string `json:"type"`
}

type SeasonMatchStatistics struct {
	Played int `json:"played"`
	Won    int `json:"won"`
	Lost   int `json:"lost"`
}

type Tier struct {
	Key Link `json:"key"`
	ID  int  `json:"id"`
}
