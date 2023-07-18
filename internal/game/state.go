package game

type GameState int64

const (
	MenuState GameState = iota
	CreateGameState
	JoinGameState
	GameCreatedState
)
