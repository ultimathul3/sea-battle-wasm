package field

func (f *Field) Update() {
	tx, ty, isTouched := f.touch.IsTouched()

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
	}
}
