package network

type GetGames struct {
	Games []string `json:"games"`
}

type GetGamesResponse struct {
	Games []string
	Error error
}
