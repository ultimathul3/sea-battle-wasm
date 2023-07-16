package field

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
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

func (f *Field) getX(i int) int {
	return i*TileSize + f.offsetX + TileSize
}

func (f *Field) getY(j int) int {
	return j*TileSize + f.offsetY + TileSize
}
