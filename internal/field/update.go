package field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (f *Field) Update() {
	tx, ty, isTouched := f.touch.IsTouched()

	mx, my := ebiten.CursorPosition()
	f.i = (my-f.offsetY-TileSize)/TileSize + 1
	f.j = (mx-f.offsetX-TileSize)/TileSize + 1

	if isTouched {
		if tx >= f.pickFrameOffsetX+212 && tx <= f.pickFrameOffsetX+212+32 && ty >= f.pickFrameOffsetY+62 && ty <= f.pickFrameOffsetY+62+32 {
			f.selectedShip = SingleDeckShipSelected
		} else if tx >= f.pickFrameOffsetX+212 && tx <= f.pickFrameOffsetX+212+64 && ty >= f.pickFrameOffsetY+14 && ty <= f.pickFrameOffsetY+14+32 {
			f.selectedShip = DoubleDeckShipSelected
		} else if tx >= f.pickFrameOffsetX+14 && tx <= f.pickFrameOffsetX+14+96 && ty >= f.pickFrameOffsetY+62 && ty <= f.pickFrameOffsetY+62+32 {
			f.selectedShip = ThreeDeckShipSelected
		} else if tx >= f.pickFrameOffsetX+14 && tx <= f.pickFrameOffsetX+14+128 && ty >= f.pickFrameOffsetY+14 && ty <= f.pickFrameOffsetY+14+32 {
			f.selectedShip = FourDeckShipSelected
		}

		if f.isFieldHover() && f.FieldMatrix[f.i][f.j] == EmptyCell {
			f.placeShip()
		}
	}

	if (isTouched && mx >= f.offsetX && mx < f.offsetX+32 && my >= f.offsetY && my < f.offsetY+32) ||
		inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		if f.placeDirection == RightDirection {
			f.placeDirection = DownDirection
		} else {
			f.placeDirection = RightDirection
		}
	}
}

func (f *Field) placeShip() {
	switch f.selectedShip {
	case SingleDeckShipSelected:
		if f.availableSingleDeckShips == 0 {
			return
		}
		f.FieldMatrix[f.i][f.j] = SingleDeckShipCell
		f.fillLeftCells(f.i, f.j, OccupiedCell)
		f.fillRightCells(f.i, f.j, OccupiedCell)
		f.fillTopAndBottomCells(f.i, f.j, OccupiedCell)
		f.availableSingleDeckShips--
	case DoubleDeckShipSelected:
		if f.availableDoubleDeckShips == 0 {
			return
		}
		if f.placeDirection == RightDirection {
			f.placeShipRight(2)
		} else {
			f.placeShipDown(2)
		}
	case ThreeDeckShipSelected:
		if f.availableThreeDeckShips == 0 {
			return
		}
		if f.placeDirection == RightDirection {
			f.placeShipRight(3)
		} else {
			f.placeShipDown(3)
		}
	case FourDeckShipSelected:
		if f.availableFourDeckShips == 0 {
			return
		}
		if f.placeDirection == RightDirection {
			f.placeShipRight(4)
		} else {
			f.placeShipDown(4)
		}
	}
}

func (f *Field) placeShipRight(shipDeck int) {
	for i := 1; i < shipDeck; i++ {
		if f.FieldMatrix[f.i][f.j+i] != EmptyCell {
			return
		}
	}

	var ship rune
	switch shipDeck {
	case 2:
		ship = DoubleDeckShipRightCell
		f.availableDoubleDeckShips--
	case 3:
		ship = ThreeDeckShipRightCell
		f.availableThreeDeckShips--
	case 4:
		ship = FourDeckShipRightCell
		f.availableFourDeckShips--
	default:
		return
	}

	f.FieldMatrix[f.i][f.j] = ship
	f.fillLeftCells(f.i, f.j, OccupiedCell)

	for i := 0; i < shipDeck; i++ {
		f.fillTopAndBottomCells(f.i, f.j+i, OccupiedCell)
		if i == shipDeck-1 {
			f.FieldMatrix[f.i][f.j+i] = ShipLeftEndCell
		} else if i > 0 {
			f.FieldMatrix[f.i][f.j+i] = ShipLeftCell
		}
	}

	f.fillRightCells(f.i, f.j+shipDeck-1, OccupiedCell)
}

func (f *Field) placeShipDown(shipDeck int) {
	for i := 1; i < shipDeck; i++ {
		if f.FieldMatrix[f.i+i][f.j] != EmptyCell {
			return
		}
	}

	var ship rune
	switch shipDeck {
	case 2:
		ship = DoubleDeckShipDownCell
		f.availableDoubleDeckShips--
	case 3:
		ship = ThreeDeckShipDownCell
		f.availableThreeDeckShips--
	case 4:
		ship = FourDeckShipDownCell
		f.availableFourDeckShips--
	default:
		return
	}

	f.FieldMatrix[f.i][f.j] = ship
	f.fillTopCells(f.i, f.j, OccupiedCell)

	for i := 0; i < shipDeck; i++ {
		f.fillLeftAndRightCells(f.i+i, f.j, OccupiedCell)
		if i == shipDeck-1 {
			f.FieldMatrix[f.i+i][f.j] = ShipUpEndCell
		} else if i > 0 {
			f.FieldMatrix[f.i+i][f.j] = ShipUpCell
		}
	}

	f.fillBottomCells(f.i+shipDeck-1, f.j, OccupiedCell)
}

func (f *Field) fillTopAndBottomCells(i, j int, cell rune) {
	f.FieldMatrix[i-1][j] = cell
	f.FieldMatrix[i+1][j] = cell
}

func (f *Field) fillLeftAndRightCells(i, j int, cell rune) {
	f.FieldMatrix[i][j-1] = cell
	f.FieldMatrix[i][j+1] = cell
}

func (f *Field) fillLeftCells(i, j int, cell rune) {
	f.FieldMatrix[i-1][j-1] = cell
	f.FieldMatrix[i][j-1] = cell
	f.FieldMatrix[i+1][j-1] = cell
}

func (f *Field) fillRightCells(i, j int, cell rune) {
	f.FieldMatrix[i-1][j+1] = cell
	f.FieldMatrix[i][j+1] = cell
	f.FieldMatrix[i+1][j+1] = cell
}

func (f *Field) fillTopCells(i, j int, cell rune) {
	f.FieldMatrix[i-1][j-1] = cell
	f.FieldMatrix[i-1][j] = cell
	f.FieldMatrix[i-1][j+1] = cell
}

func (f *Field) fillBottomCells(i, j int, cell rune) {
	f.FieldMatrix[i+1][j-1] = cell
	f.FieldMatrix[i+1][j] = cell
	f.FieldMatrix[i+1][j+1] = cell
}

func (f *Field) isFieldHover() bool {
	mx, my := ebiten.CursorPosition()
	return mx >= f.offsetX+32 && mx < f.offsetX+341 && my >= f.offsetY+32 && my < f.offsetY+341
}
