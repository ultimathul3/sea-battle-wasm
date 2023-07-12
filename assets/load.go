package assets

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Assets struct {
	BackgroundImages []*ebiten.Image
}

func New() *Assets {
	return &Assets{
		BackgroundImages: loadBackgroundImages(),
	}
}

func loadBackgroundImages() []*ebiten.Image {
	var images []*ebiten.Image

	dir, _ := BackgroundImagesDir.ReadDir(BackgroundImagesDirPath)

	for i := range dir {
		data, _ := BackgroundImagesDir.ReadFile(
			fmt.Sprintf("%s/%d.png", BackgroundImagesDirPath, i),
		)
		images = append(images, imageFromBytes(data))
	}

	return images
}

func imageFromBytes(imageBytes []byte) *ebiten.Image {
	img, _, _ := image.Decode(bytes.NewReader(imageBytes))
	return ebiten.NewImageFromImage(img)
}
