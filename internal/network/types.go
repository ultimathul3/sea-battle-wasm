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
	Status        string `json:"status"`
	X             uint32 `json:"x"`
	Y             uint32 `json:"y"`
	DestroyedShip string `json:"destroyed_ship"`
	DestroyedX    uint32 `json:"destroyed_x"`
	DestroyedY    uint32 `json:"destroyed_y"`
	Message       string `json:"message"`
}

type WaitRequest struct {
	Uuid string `json:"uuid"`
}

type WaitResponse struct {
	Status        GameStatus
	X             uint32
	Y             uint32
	DestroyedShip Ship
	DestroyedX    uint32
	DestroyedY    uint32
	Message       string
	Error         error
}

type GameStatus string

const (
	UnknownStatus            GameStatus = "UNKNOWN"
	GameCreatedStatus        GameStatus = "GAME_CREATED"
	WaitingForOpponentStatus GameStatus = "WAITING_FOR_OPPONENT"
	GameStartedStatus        GameStatus = "GAME_STARTED"
	HostHitStatus            GameStatus = "HOST_HIT"
	HostMissStatus           GameStatus = "HOST_MISS"
	OpponentHitStatus        GameStatus = "OPPONENT_HIT"
	OpponentMissStatus       GameStatus = "OPPONENT_MISS"
	HostWonStatus            GameStatus = "HOST_WON"
	OpponentWonStatus        GameStatus = "OPPONENT_WON"
)

type Ship string

const (
	UnknownShip         Ship = "UNKNOWN"
	SingleDeckShip      Ship = "SINGLE_DECK"
	DoubleDeckShipDown  Ship = "DOUBLE_DECK_DOWN"
	ThreeDeckShipDown   Ship = "THREE_DECK_DOWN"
	FourDeckShipDown    Ship = "FOUR_DECK_DOWN"
	DoubleDeckShipRight Ship = "DOUBLE_DECK_RIGHT"
	ThreeDeckShipRight  Ship = "THREE_DECK_RIGHT"
	FourDeckShipRight   Ship = "FOUR_DECK_RIGHT"
)

type Shoot struct {
	Status        string `json:"status"`
	DestroyedShip string `json:"destroyed_ship"`
	X             uint32 `json:"x"`
	Y             uint32 `json:"y"`
}

type ShootRequest struct {
	HostNickname string `json:"host_nickname"`
	X            uint32 `json:"x"`
	Y            uint32 `json:"y"`
	Uuid         string `json:"uuid"`
}

type ShootResponse struct {
	Status        GameStatus
	DestroyedShip Ship
	X             uint32
	Y             uint32
	Error         error
}
