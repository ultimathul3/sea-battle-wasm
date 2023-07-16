package field

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ultimathul3/sea-battle-wasm/pkg/utils"
)

func (f *Field) Draw(screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()

	utils.DrawInCoordsWithColor(screen, f.fieldImage, f.offsetX, f.offsetY, f.transparentColor)

	i := (mx - f.offsetX - TileSize) / TileSize
	j := (my - f.offsetY - TileSize) / TileSize

	if mx >= f.offsetX+32 && mx < f.offsetX+341 && my >= f.offsetY+32 && my < f.offsetY+341 {
		utils.DrawInCoordsWithColor(screen, f.selectImage, f.getX(i), f.getY(j), f.transparentColor)
	}

	utils.DrawInCoordsWithColor(screen, f.pickFrameImage, f.pickFrameOffsetX, f.pickFrameOffsetY, f.transparentColor)

	switch f.selectedShip {
	case SingleDeckShipSelected:
		utils.DrawInCoordsWithColor(screen, f.singleDeckShipPickImage, f.pickFrameOffsetX+212, f.pickFrameOffsetY+62, f.transparentColor)
	case DoubleDeckShipSelected:
		utils.DrawInCoordsWithColor(screen, f.doubleDeckShipPickImage, f.pickFrameOffsetX+212, f.pickFrameOffsetY+14, f.transparentColor)
	case ThreeDeckShipSelected:
		utils.DrawInCoordsWithColor(screen, f.threeDeckShipPickImage, f.pickFrameOffsetX+14, f.pickFrameOffsetY+62, f.transparentColor)
	case FourDeckShipSelected:
		utils.DrawInCoordsWithColor(screen, f.fourDeckShipPickImage, f.pickFrameOffsetX+14, f.pickFrameOffsetY+14, f.transparentColor)
	}

	f.text.DrawMedium(screen, fmt.Sprintf("x%d", f.availableSingleDeckShips), f.pickFrameOffsetX+212+64+8, f.pickFrameOffsetY+62, f.transparentColor)
	f.text.DrawMedium(screen, fmt.Sprintf("x%d", f.availableDoubleDeckShips), f.pickFrameOffsetX+212+64+8, f.pickFrameOffsetY+14, f.transparentColor)
	f.text.DrawMedium(screen, fmt.Sprintf("x%d", f.availableThreeDeckShips), f.pickFrameOffsetX+14+128+8, f.pickFrameOffsetY+62, f.transparentColor)
	f.text.DrawMedium(screen, fmt.Sprintf("x%d", f.availableFourDeckShips), f.pickFrameOffsetX+14+128+8, f.pickFrameOffsetY+14, f.transparentColor)

}
