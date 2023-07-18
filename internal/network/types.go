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

type JoinGame struct {
	OpponentUuid string `json:"opponent_uuid"`
}

type JoinGameRequest struct {
	HostNickname     string `json:"host_nickname"`
	OpponentNickname string `json:"opponent_nickname"`
}

type JoinGameResponse struct {
	OpponentUuid string
	Error        error
}

type StartGame struct {
}

type StartGameRequest struct {
	HostNickname  string `json:"host_nickname"`
	OpponentField string `json:"opponent_field"`
	OpponentUuid  string `json:"opponent_uuid"`
}

type StartGameResponse struct {
	Error error
}

type Wait struct {
	Status  string `json:"status"`
	X       uint32 `json:"x"`
	Y       uint32 `json:"y"`
	Message string `json:"message"`
}

type WaitRequest struct {
	Uuid string `json:"uuid"`
}

type WaitResponse struct {
	Status  WaitStatus
	X       uint32
	Y       uint32
	Message string
	Error   error
}

type WaitStatus string

const (
	GameCreatedStatus        WaitStatus = "GAME_CREATED"
	WaitingForOpponentStatus WaitStatus = "WAITING_FOR_OPPONENT"
	GameStartedStatus        WaitStatus = "GAME_STARTED"
	HostHitStatus            WaitStatus = "HOST_HIT"
	HostMissStatus           WaitStatus = "HOST_MISS"
	OpponentHitStatus        WaitStatus = "OPPONENT_HIT"
	OpponentMissStatus       WaitStatus = "OPPONENT_MISS"
	HostWonStatus            WaitStatus = "HOST_WON"
	OpponentWonStatus        WaitStatus = "OPPONENT_WON"
)
