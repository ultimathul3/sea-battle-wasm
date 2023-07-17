package utils

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawInCoords(screen *ebiten.Image, image *ebiten.Image, x, y int) {
	var op ebiten.DrawImageOptions
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(image, &op)
}

func DrawInCoordsWithColor(screen *ebiten.Image, image *ebiten.Image, x, y int, color color.Color) {
	var op ebiten.DrawImageOptions
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(color)
	screen.DrawImage(image, &op)
}

func DrawInCoordsWithColorAndRotate(screen *ebiten.Image, image *ebiten.Image, x, y int, color color.Color, angle float64) {
	var op ebiten.DrawImageOptions
	op.GeoM.Rotate(angle)
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(color)
	screen.DrawImage(image, &op)
}

func DrawInCoordsWithRotate(screen *ebiten.Image, image *ebiten.Image, x, y int, angle float64) {
	var op ebiten.DrawImageOptions
	op.GeoM.Rotate(angle)
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(image, &op)
}
