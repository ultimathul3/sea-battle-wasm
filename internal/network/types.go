package network

type GetGames struct {
	Games []string `json:"games"`
}

type GetGamesResponse struct {
	Games []string
	Error error
}

type CreateGame struct {
	HostUuid string `json:"host_uuid"`
}

type CreateGameRequest struct {
	HostNickname string `json:"host_nickname"`
	HostField    string `json:"host_field"`
}

type CreateGameResponse struct {
	HostUuid string
	Error    error
}
