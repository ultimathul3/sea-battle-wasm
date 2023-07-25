package field

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ultimathul3/sea-battle-wasm/assets"
	"github.com/ultimathul3/sea-battle-wasm/internal/animation"
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
	hitMask     [][]bool
	i, j        int

	state FieldState

	explosionAnimation *animation.Animation
}

func (f *Field) GetFieldMatrix() [][]rune {
	return f.fieldMatrix
}

func (f *Field) SetFieldMatrix(matrix [][]rune) {
	f.fieldMatrix = matrix
	f.availableSingleDeckShips = 0
	f.availableDoubleDeckShips = 0
	f.availableThreeDeckShips = 0
	f.availableFourDeckShips = 0
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

	f.explosionAnimation = animation.New(
		f.assets.ExplosionImages,
	)

	return f
}

func (f *Field) GetNumOfAvailableShips() int {
	return f.availableSingleDeckShips +
		f.availableDoubleDeckShips +
		f.availableThreeDeckShips +
		f.availableFourDeckShips
}

func (f *Field) SetState(state FieldState) {
	f.state = state
}

func (f *Field) ConvertFieldRuneMatrixToString() string {
	var field string

	for i := 1; i < FieldDimension+1; i++ {
		field += string(f.fieldMatrix[i][1 : FieldDimension+1])
	}

	return field
}

func (f *Field) initFieldMatrix() {
	f.fieldMatrix = make([][]rune, 0, FieldDimension+2)

	f.hitMask = make([][]bool, FieldDimension)
	for i := 0; i < FieldDimension; i++ {
		f.hitMask[i] = make([]bool, FieldDimension)
	}

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

func (f *Field) fillTopAndBottomCells(i, j int, cell rune) {
	f.fieldMatrix[i-1][j] = cell
	f.fieldMatrix[i+1][j] = cell
}

func (f *Field) fillLeftAndRightCells(i, j int, cell rune) {
	f.fieldMatrix[i][j-1] = cell
	f.fieldMatrix[i][j+1] = cell
}

func (f *Field) fillLeftCells(i, j int, cell rune) {
	f.fieldMatrix[i-1][j-1] = cell
	f.fieldMatrix[i][j-1] = cell
	f.fieldMatrix[i+1][j-1] = cell
}

func (f *Field) fillRightCells(i, j int, cell rune) {
	f.fieldMatrix[i-1][j+1] = cell
	f.fieldMatrix[i][j+1] = cell
	f.fieldMatrix[i+1][j+1] = cell
}

func (f *Field) fillTopCells(i, j int, cell rune) {
	f.fieldMatrix[i-1][j-1] = cell
	f.fieldMatrix[i-1][j] = cell
	f.fieldMatrix[i-1][j+1] = cell
}

func (f *Field) fillBottomCells(i, j int, cell rune) {
	f.fieldMatrix[i+1][j-1] = cell
	f.fieldMatrix[i+1][j] = cell
	f.fieldMatrix[i+1][j+1] = cell
}

func (f *Field) isFieldHover(x, y int) bool {
	return x >= f.offsetX+TileSize+1 && x < f.offsetX+341 && y >= f.offsetY+TileSize+1 && y < f.offsetY+341
}
