package field

type FieldState int

const (
	PlacementState FieldState = iota
	PlacementFinishedState
	CurtainState
	ShootState
)
