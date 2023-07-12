package assets

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Assets struct {
	BackgroundImage *ebiten.Image
}

func New() *Assets {
	return &Assets{
		BackgroundImage: imageFromBytes(BackgroundImage),
	}
}

func imageFromBytes(imageBytes []byte) *ebiten.Image {
	img, _, _ := image.Decode(bytes.NewReader(imageBytes))
	return ebiten.NewImageFromImage(img)
}
