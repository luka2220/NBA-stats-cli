package models

// Root JSON reposone fields
type Standings struct {
	Errors   []int                 `json:"errors"`
	Response [][]StandingsResponse `json:"response"`
}

// Standings JSON repsonse fields
type StandingsResponse struct {
	Position int        `json:"position"`
	Group    Conference `json:"group"`
	Team     Team       `json:"team"`
	Games    Games      `json:"games"`
}

// Conference JSON repsonse fields
type Conference struct {
	Name string `json:"name"`
}

// Team JSON repsonse fields
type Team struct {
	Name string `json:"name"`
}

// Games JSON response fields
type Games struct {
	Played int               `json:"played"`
	Win    WinLossPercentage `json:"win"`
	Lose   WinLossPercentage `json:"lose"`
}

// Win Loss and Percentage stats JSON response fields
type WinLossPercentage struct {
	Total      int    `json:"total"`
	Percentage string `json:"percentage"`
}
