package field

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ultimathul3/sea-battle-wasm/pkg/utils"
)

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

type Texter interface {
	DrawMedium(screen *ebiten.Image, text string, x, y int, color color.Color)
}

type Toucher interface {
	IsTouched() (int, int, bool)
}

type Field struct {
	offsetX, offsetY int

	pickFrameOffsetX, pickFrameOffsetY int

	singleDeckShipImage *ebiten.Image
	doubleDeckShipImage *ebiten.Image
	threeDeckShipImage  *ebiten.Image
	fourDeckShipImage   *ebiten.Image

	singleDeckShipPickImage *ebiten.Image
	doubleDeckShipPickImage *ebiten.Image
	threeDeckShipPickImage  *ebiten.Image
	fourDeckShipPickImage   *ebiten.Image
	pickFrameImage          *ebiten.Image

	fieldImage  *ebiten.Image
	selectImage *ebiten.Image

	transparentColor color.Color

	text  Texter
	touch Toucher

	selectedShip SelectedShip

	availableSingleDeckShips int
	availableDoubleDeckShips int
	availableThreeDeckShips  int
	availableFourDeckShips   int
}

func New(
	offsetX, offsetY int,
	singleDeckShipImage, doubleDeckShipImage, threeDeckShipImage, fourDeckShipImage *ebiten.Image,
	singleDeckShipPickImage, doubleDeckShipPickImage, threeDeckShipPickImage, fourDeckShipPickImage, pickFrameImage *ebiten.Image,
	fieldImage *ebiten.Image, selectImage *ebiten.Image, transparentColor color.Color, text Texter, touch Toucher,
) *Field {
	return &Field{
		offsetX: offsetX,
		offsetY: offsetY,

		pickFrameOffsetX: offsetX,
		pickFrameOffsetY: offsetY + 350,

		singleDeckShipImage: singleDeckShipImage,
		doubleDeckShipImage: doubleDeckShipImage,
		threeDeckShipImage:  threeDeckShipImage,
		fourDeckShipImage:   fourDeckShipImage,

		singleDeckShipPickImage: singleDeckShipPickImage,
		doubleDeckShipPickImage: doubleDeckShipPickImage,
		threeDeckShipPickImage:  threeDeckShipPickImage,
		fourDeckShipPickImage:   fourDeckShipPickImage,
		pickFrameImage:          pickFrameImage,

		fieldImage:  fieldImage,
		selectImage: selectImage,

		transparentColor: transparentColor,

		text:  text,
		touch: touch,

		selectedShip: SingleDeckShipSelected,

		availableSingleDeckShips: 4,
		availableDoubleDeckShips: 3,
		availableThreeDeckShips:  2,
		availableFourDeckShips:   1,
	}
}

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

func (f *Field) getX(i int) int {
	return i*TileSize + f.offsetX + TileSize
}

func (f *Field) getY(j int) int {
	return j*TileSize + f.offsetY + TileSize
}
