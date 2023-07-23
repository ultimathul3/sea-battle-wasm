package field

const (
	FieldDimension = 10
	TileSize       = 31
)

type SelectedShip uint8

const (
	SingleDeckShipSelected SelectedShip = iota
	DoubleDeckShipSelected
	ThreeDeckShipSelected
	FourDeckShipSelected
)

type PlaceDirection uint8

const (
	RightDirection PlaceDirection = iota
	DownDirection
)

const (
	EmptyCell               = rune(' ')
	OccupiedCell            = rune('x')
	SingleDeckShipCell      = rune('1')
	DoubleDeckShipDownCell  = rune('2')
	ThreeDeckShipDownCell   = rune('3')
	FourDeckShipDownCell    = rune('4')
	DoubleDeckShipRightCell = rune('@')
	ThreeDeckShipRightCell  = rune('#')
	FourDeckShipRightCell   = rune('$')
	ShipLeftCell            = rune('←')
	ShipLeftEndCell         = rune('◄')
	ShipUpCell              = rune('↑')
	ShipUpEndCell           = rune('▲')
	HitCell                 = rune('o')
	MissCell                = rune('.')
	FrameCell               = rune('~')
)

type Ship uint8

const (
	UnknownShip Ship = iota
	SingleDeckShip
	DoubleDeckShipDown
	ThreeDeckShipDown
	FourDeckShipDown
	DoubleDeckShipRight
	ThreeDeckShipRight
	FourDeckShipRight
)
