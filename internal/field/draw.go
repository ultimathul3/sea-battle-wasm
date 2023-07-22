package field

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ultimathul3/sea-battle-wasm/pkg/utils"
)

func (f *Field) Draw(screen *ebiten.Image) {
	utils.DrawInCoordsWithColor(screen, f.assets.FieldImage, f.offsetX, f.offsetY, f.transparentColor)

	for i := 1; i < FieldDimension+1; i++ {
		for j := 1; j < FieldDimension+1; j++ {
			switch f.fieldMatrix[i][j] {
			case SingleDeckShipCell:
				utils.DrawInCoordsWithColor(screen, f.assets.SingleDeckShipImage, f.getX(j), f.getY(i), f.transparentColor)
			case DoubleDeckShipRightCell:
				utils.DrawInCoordsWithColorAndRotate(screen, f.assets.DoubleDeckShipImage, f.getX(j), f.getY(i)+32, f.transparentColor, -math.Pi/2)
			case DoubleDeckShipDownCell:
				utils.DrawInCoordsWithColorAndRotate(screen, f.assets.DoubleDeckShipImage, f.getX(j)+32, f.getY(i)+63, f.transparentColor, math.Pi)
			case ThreeDeckShipRightCell:
				utils.DrawInCoordsWithColorAndRotate(screen, f.assets.ThreeDeckShipImage, f.getX(j), f.getY(i)+32, f.transparentColor, -math.Pi/2)
			case ThreeDeckShipDownCell:
				utils.DrawInCoordsWithColorAndRotate(screen, f.assets.ThreeDeckShipImage, f.getX(j)+32, f.getY(i)+94, f.transparentColor, math.Pi)
			case FourDeckShipRightCell:
				utils.DrawInCoordsWithColorAndRotate(screen, f.assets.FourDeckShipImage, f.getX(j), f.getY(i)+32, f.transparentColor, -math.Pi/2)
			case FourDeckShipDownCell:
				utils.DrawInCoordsWithColorAndRotate(screen, f.assets.FourDeckShipImage, f.getX(j)+32, f.getY(i)+125, f.transparentColor, math.Pi)
			case MissCell:
				utils.DrawInCoordsWithColor(screen, f.assets.MissImage, f.getX(j)+1, f.getY(i)+1, f.transparentColor)
			}

			if f.hitMask[i-1][j-1] {
				utils.DrawInCoordsWithColor(screen, f.assets.HitImage, f.getX(j)+1, f.getY(i)+1, f.transparentColor)
			}
		}
	}

	switch f.state {
	case PlacementState:
		f.drawPlacementState(screen)
	case CurtainState:
		f.drawCurtainState(screen)
	case ShootState:
		f.drawShootState(screen)
	}
}

func (f *Field) drawCurtainState(screen *ebiten.Image) {
	utils.DrawInCoordsWithColor(screen, f.assets.CurtainImage, f.offsetX, f.offsetY, f.transparentColor)
}

func (f *Field) drawPlacementState(screen *ebiten.Image) {
	switch f.placeDirection {
	case RightDirection:
		utils.DrawInCoordsWithColorAndRotate(screen, f.assets.ArrowImage, f.offsetX+30, f.offsetY+25, f.transparentColor, math.Pi)
	case DownDirection:
		utils.DrawInCoordsWithColorAndRotate(screen, f.assets.ArrowImage, f.offsetX+8, f.offsetY+30, f.transparentColor, 3*math.Pi/2)
	}

	mx, my := ebiten.CursorPosition()
	if f.isFieldHover(mx, my) {
		utils.DrawInCoordsWithColor(screen, f.assets.SelectImage, f.getX(f.j), f.getY(f.i), f.transparentColor)
	}

	utils.DrawInCoordsWithColor(screen, f.assets.PickFrameImage, f.pickFrameOffsetX, f.pickFrameOffsetY, f.transparentColor)

	switch f.selectedShip {
	case SingleDeckShipSelected:
		utils.DrawInCoordsWithColor(screen, f.assets.SingleDeckShipPickImage, f.pickFrameOffsetX+212, f.pickFrameOffsetY+62, f.transparentColor)
	case DoubleDeckShipSelected:
		utils.DrawInCoordsWithColor(screen, f.assets.DoubleDeckShipPickImage, f.pickFrameOffsetX+212, f.pickFrameOffsetY+14, f.transparentColor)
	case ThreeDeckShipSelected:
		utils.DrawInCoordsWithColor(screen, f.assets.ThreeDeckShipPickImage, f.pickFrameOffsetX+14, f.pickFrameOffsetY+62, f.transparentColor)
	case FourDeckShipSelected:
		utils.DrawInCoordsWithColor(screen, f.assets.FourDeckShipPickImage, f.pickFrameOffsetX+14, f.pickFrameOffsetY+14, f.transparentColor)
	}

	f.text.DrawMedium(screen, fmt.Sprintf("x%d", f.availableSingleDeckShips), f.pickFrameOffsetX+212+64+8, f.pickFrameOffsetY+62, f.transparentColor)
	f.text.DrawMedium(screen, fmt.Sprintf("x%d", f.availableDoubleDeckShips), f.pickFrameOffsetX+212+64+8, f.pickFrameOffsetY+14, f.transparentColor)
	f.text.DrawMedium(screen, fmt.Sprintf("x%d", f.availableThreeDeckShips), f.pickFrameOffsetX+14+128+8, f.pickFrameOffsetY+62, f.transparentColor)
	f.text.DrawMedium(screen, fmt.Sprintf("x%d", f.availableFourDeckShips), f.pickFrameOffsetX+14+128+8, f.pickFrameOffsetY+14, f.transparentColor)
}

func (f *Field) drawShootState(screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()
	if f.isFieldHover(mx, my) {
		utils.DrawInCoordsWithColor(screen, f.assets.SelectImage, f.getX(f.j), f.getY(f.i), f.transparentColor)
	}
}
