package field

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ultimathul3/sea-battle-wasm/pkg/utils"
)

const (
	FieldDimension = 10
	TileSize       = 31
)

type Field struct {
	offsetX, offsetY int

	singleDeckShipImage *ebiten.Image
	doubleDeckShipImage *ebiten.Image
	threeDeckShipImage  *ebiten.Image
	fourDeckShipImage   *ebiten.Image

	fieldImage  *ebiten.Image
	selectImage *ebiten.Image

	transparentColor color.Color
}

func New(
	offsetX, offsetY int,
	singleDeckShipImage, doubleDeckShipImage, threeDeckShipImage, fourDeckShipImage *ebiten.Image,
	fieldImage *ebiten.Image, selectImage *ebiten.Image, transparentColor color.Color,
) *Field {
	return &Field{
		offsetX: offsetX,
		offsetY: offsetY,

		singleDeckShipImage: singleDeckShipImage,
		doubleDeckShipImage: doubleDeckShipImage,
		threeDeckShipImage:  threeDeckShipImage,
		fourDeckShipImage:   fourDeckShipImage,

		fieldImage:  fieldImage,
		selectImage: selectImage,

		transparentColor: transparentColor,
	}
}

func (f *Field) Update() {
}

func (f *Field) Draw(screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()

	utils.DrawInCoordsWithColor(screen, f.fieldImage, f.offsetX, f.offsetY, f.transparentColor)

	for i := 0; i < FieldDimension; i++ {
		for j := 0; j < FieldDimension; j++ {
			if mx >= f.getX(i) && mx < f.getX(i)+TileSize && my >= f.getY(j) && my < f.getY(j)+TileSize {
				utils.DrawInCoordsWithColor(screen, f.selectImage, f.getX(i), f.getY(j), f.transparentColor)
			}
		}
	}
}

func (f *Field) getX(i int) int {
	return i*TileSize + f.offsetX + TileSize
}

func (f *Field) getY(j int) int {
	return j*TileSize + f.offsetY + TileSize
}
