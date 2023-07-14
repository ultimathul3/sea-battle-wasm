package state

type State int64

const (
	Menu State = iota
	CreateGame
	JoinGame
)
