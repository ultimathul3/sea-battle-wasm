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
