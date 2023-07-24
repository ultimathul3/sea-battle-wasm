package field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (f *Field) Update() {
	tx, ty, isTouched := f.touch.IsTouched()

	mx, my := ebiten.CursorPosition()
	if isTouched {
		mx, my = tx, ty
	}

	f.i = (my-f.offsetY-TileSize)/TileSize + 1
	f.j = (mx-f.offsetX-TileSize)/TileSize + 1

	switch f.state {
	case PlacementState:
		f.updatePlacementState(tx, ty, mx, my, isTouched)
	}
}

func (f *Field) IsEmptyCellTouched() (int, int, bool) {
	tx, ty, isTouched := f.touch.IsTouched()

	if isTouched && f.isFieldHover(tx, ty) && f.fieldMatrix[f.i][f.j] == EmptyCell {
		return f.j - 1, f.i - 1, true
	}

	return -1, -1, false
}

func (f *Field) SetMissCell(x, y int) {
	f.fieldMatrix[y+1][x+1] = MissCell
}

func (f *Field) SetHitCell(x, y int) {
	f.hitMask[y][x] = true
}

func (f *Field) DestroyShip(ship Ship, x, y int) {
	switch ship {
	case SingleDeckShip:
		f.fieldMatrix[y+1][x+1] = SingleDeckShipCell
		f.fillLeftCells(y+1, x+1, MissCell)
		f.fillRightCells(y+1, x+1, MissCell)
		f.fillTopAndBottomCells(y+1, x+1, MissCell)
	case DoubleDeckShipDown:
		f.cleanDown(y+1, x+1, 2)
		f.placeShipDown(y+1, x+1, MissCell, 2)
	case ThreeDeckShipDown:
		f.cleanDown(y+1, x+1, 3)
		f.placeShipDown(y+1, x+1, MissCell, 3)
	case FourDeckShipDown:
		f.cleanDown(y+1, x+1, 4)
		f.placeShipDown(y+1, x+1, MissCell, 4)
	case DoubleDeckShipRight:
		f.cleanRight(y+1, x+1, 2)
		f.placeShipRight(y+1, x+1, MissCell, 2)
	case ThreeDeckShipRight:
		f.cleanRight(y+1, x+1, 3)
		f.placeShipRight(y+1, x+1, MissCell, 3)
	case FourDeckShipRight:
		f.cleanRight(y+1, x+1, 4)
		f.placeShipRight(y+1, x+1, MissCell, 4)
	}
}

func (f *Field) cleanRight(i, j, count int) {
	for l := 0; l < count; l++ {
		f.fieldMatrix[i][j+l] = EmptyCell
	}
}

func (f *Field) cleanDown(i, j, count int) {
	for l := 0; l < count; l++ {
		f.fieldMatrix[i+l][j] = EmptyCell
	}
}

func (f *Field) updatePlacementState(tx, ty, mx, my int, isTouched bool) {
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

		if f.isFieldHover(mx, my) {
			if f.fieldMatrix[f.i][f.j] == EmptyCell {
				f.placeShip()
			} else {
				f.removeShip()
			}
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

func (f *Field) removeShip() {
	tmpI, tmpJ := f.i, f.j

	switch f.fieldMatrix[f.i][f.j] {
	case ShipLeftCell, ShipLeftEndCell:
		f.j--
		for f.fieldMatrix[f.i][f.j] == ShipLeftCell {
			f.j--
		}
	case ShipUpCell, ShipUpEndCell:
		f.i--
		for f.fieldMatrix[f.i][f.j] == ShipUpCell {
			f.i--
		}
	}

	switch f.fieldMatrix[f.i][f.j] {
	case SingleDeckShipCell:
		f.availableSingleDeckShips++
		f.fieldMatrix[f.i][f.j] = EmptyCell
		f.fillLeftCells(f.i, f.j, EmptyCell)
		f.fillRightCells(f.i, f.j, EmptyCell)
		f.fillTopAndBottomCells(f.i, f.j, EmptyCell)
	case DoubleDeckShipRightCell:
		f.availableDoubleDeckShips++
		f.removeShipRight()
	case ThreeDeckShipRightCell:
		f.availableThreeDeckShips++
		f.removeShipRight()
	case FourDeckShipRightCell:
		f.availableFourDeckShips++
		f.removeShipRight()
	case DoubleDeckShipDownCell:
		f.availableDoubleDeckShips++
		f.removeShipDown()
	case ThreeDeckShipDownCell:
		f.availableThreeDeckShips++
		f.removeShipDown()
	case FourDeckShipDownCell:
		f.availableFourDeckShips++
		f.removeShipDown()
	default:
		return
	}

	for i := 0; i < FieldDimension+2; i++ {
		f.fieldMatrix[0][i] = FrameCell
		f.fieldMatrix[i][0] = FrameCell
		f.fieldMatrix[i][FieldDimension+1] = FrameCell
		f.fieldMatrix[FieldDimension+1][i] = FrameCell
	}

	for i := 1; i < FieldDimension+1; i++ {
		for j := 1; j < FieldDimension+1; j++ {
			switch f.fieldMatrix[i][j] {
			case SingleDeckShipCell:
				f.fillLeftCells(i, j, OccupiedCell)
				f.fillTopAndBottomCells(i, j, OccupiedCell)
				f.fillRightCells(i, j, OccupiedCell)
			case DoubleDeckShipRightCell, ThreeDeckShipRightCell, FourDeckShipRightCell:
				f.fillLeftCells(i, j, OccupiedCell)
				f.fillTopAndBottomCells(i, j, OccupiedCell)
			case DoubleDeckShipDownCell, ThreeDeckShipDownCell, FourDeckShipDownCell:
				f.fillTopCells(i, j, OccupiedCell)
				f.fillLeftAndRightCells(i, j, OccupiedCell)
			case ShipLeftCell:
				f.fillTopAndBottomCells(i, j, OccupiedCell)
			case ShipUpCell:
				f.fillLeftAndRightCells(i, j, OccupiedCell)
			case ShipLeftEndCell:
				f.fillTopAndBottomCells(i, j, OccupiedCell)
				f.fillRightCells(i, j, OccupiedCell)
			case ShipUpEndCell:
				f.fillLeftAndRightCells(i, j, OccupiedCell)
				f.fillBottomCells(i, j, OccupiedCell)
			}
		}
	}

	f.i, f.j = tmpI, tmpJ
}

func (f *Field) removeShipRight() {
	f.fieldMatrix[f.i][f.j] = EmptyCell
	f.fillLeftCells(f.i, f.j, EmptyCell)
	for f.fieldMatrix[f.i][f.j] != ShipLeftEndCell {
		f.fillTopAndBottomCells(f.i, f.j, EmptyCell)
		f.fieldMatrix[f.i][f.j] = EmptyCell
		f.j++
	}
	f.fieldMatrix[f.i][f.j] = EmptyCell
	f.fillTopAndBottomCells(f.i, f.j, EmptyCell)
	f.fillRightCells(f.i, f.j, EmptyCell)
}

func (f *Field) removeShipDown() {
	f.fieldMatrix[f.i][f.j] = EmptyCell
	f.fillTopCells(f.i, f.j, EmptyCell)
	for f.fieldMatrix[f.i][f.j] != ShipUpEndCell {
		f.fillLeftAndRightCells(f.i, f.j, EmptyCell)
		f.fieldMatrix[f.i][f.j] = EmptyCell
		f.i++
	}
	f.fieldMatrix[f.i][f.j] = EmptyCell
	f.fillLeftAndRightCells(f.i, f.j, EmptyCell)
	f.fillBottomCells(f.i, f.j, EmptyCell)
}

func (f *Field) placeShip() {
	switch f.selectedShip {
	case SingleDeckShipSelected:
		if f.availableSingleDeckShips == 0 {
			return
		}
		f.fieldMatrix[f.i][f.j] = SingleDeckShipCell
		f.fillLeftCells(f.i, f.j, OccupiedCell)
		f.fillRightCells(f.i, f.j, OccupiedCell)
		f.fillTopAndBottomCells(f.i, f.j, OccupiedCell)
		f.availableSingleDeckShips--
	case DoubleDeckShipSelected:
		if f.availableDoubleDeckShips == 0 {
			return
		}
		if f.placeDirection == RightDirection {
			f.placeShipRight(f.i, f.j, OccupiedCell, 2)
		} else {
			f.placeShipDown(f.i, f.j, OccupiedCell, 2)
		}
	case ThreeDeckShipSelected:
		if f.availableThreeDeckShips == 0 {
			return
		}
		if f.placeDirection == RightDirection {
			f.placeShipRight(f.i, f.j, OccupiedCell, 3)
		} else {
			f.placeShipDown(f.i, f.j, OccupiedCell, 3)
		}
	case FourDeckShipSelected:
		if f.availableFourDeckShips == 0 {
			return
		}
		if f.placeDirection == RightDirection {
			f.placeShipRight(f.i, f.j, OccupiedCell, 4)
		} else {
			f.placeShipDown(f.i, f.j, OccupiedCell, 4)
		}
	}
}

func (f *Field) placeShipRight(i, j int, cellAround rune, shipDeck int) {
	for l := 1; l < shipDeck; l++ {
		if f.fieldMatrix[i][j+l] != EmptyCell {
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

	f.fieldMatrix[i][j] = ship
	f.fillLeftCells(i, j, cellAround)

	for l := 0; l < shipDeck; l++ {
		f.fillTopAndBottomCells(i, j+l, cellAround)
		if l == shipDeck-1 {
			f.fieldMatrix[i][j+l] = ShipLeftEndCell
		} else if l > 0 {
			f.fieldMatrix[i][j+l] = ShipLeftCell
		}
	}

	f.fillRightCells(i, j+shipDeck-1, cellAround)
}

func (f *Field) placeShipDown(i, j int, cellAround rune, shipDeck int) {
	for l := 1; l < shipDeck; l++ {
		if f.fieldMatrix[i+l][j] != EmptyCell {
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

	f.fieldMatrix[i][j] = ship
	f.fillTopCells(i, j, cellAround)

	for l := 0; l < shipDeck; l++ {
		f.fillLeftAndRightCells(i+l, j, cellAround)
		if l == shipDeck-1 {
			f.fieldMatrix[i+l][j] = ShipUpEndCell
		} else if l > 0 {
			f.fieldMatrix[i+l][j] = ShipUpCell
		}
	}

	f.fillBottomCells(i+shipDeck-1, j, cellAround)
}
