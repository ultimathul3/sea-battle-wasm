package field

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ultimathul3/sea-battle-wasm/assets"
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

	assets *assets.Assets

	transparentColor color.Color

	text  Texter
	touch Toucher

	selectedShip   SelectedShip
	placeDirection PlaceDirection

	availableSingleDeckShips int
	availableDoubleDeckShips int
	availableThreeDeckShips  int
	availableFourDeckShips   int

	fieldMatrix [][]rune
	i, j        int

	state FieldState
}

func New(
	offsetX, offsetY int,
	assets *assets.Assets, transparentColor color.Color, text Texter, touch Toucher,
	state FieldState,
) *Field {
	f := &Field{
		offsetX: offsetX,
		offsetY: offsetY,

		pickFrameOffsetX: offsetX,
		pickFrameOffsetY: offsetY + 350,

		assets: assets,

		transparentColor: transparentColor,

		text:  text,
		touch: touch,

		selectedShip:   SingleDeckShipSelected,
		placeDirection: RightDirection,

		availableSingleDeckShips: 4,
		availableDoubleDeckShips: 3,
		availableThreeDeckShips:  2,
		availableFourDeckShips:   1,

		state: state,
	}

	f.initFieldMatrix()

	return f
}

func (f *Field) initFieldMatrix() {
	f.fieldMatrix = make([][]rune, 0, FieldDimension+2)

	frame := make([]rune, 0, FieldDimension+2)
	frame2 := make([]rune, 0, FieldDimension+2)
	for i := 0; i < FieldDimension+2; i++ {
		frame = append(frame, FrameCell)
		frame2 = append(frame2, FrameCell)
	}

	f.fieldMatrix = append(f.fieldMatrix, frame)
	for i := 1; i < FieldDimension+1; i++ {
		f.fieldMatrix = append(f.fieldMatrix, []rune{FrameCell})
		for j := 0; j < FieldDimension; j++ {
			f.fieldMatrix[i] = append(f.fieldMatrix[i], EmptyCell)
		}
		f.fieldMatrix[i] = append(f.fieldMatrix[i], FrameCell)
	}
	f.fieldMatrix = append(f.fieldMatrix, frame2)
}

func (f *Field) getX(j int) int {
	return (j-1)*TileSize + f.offsetX + TileSize
}

func (f *Field) getY(i int) int {
	return (i-1)*TileSize + f.offsetY + TileSize
}
