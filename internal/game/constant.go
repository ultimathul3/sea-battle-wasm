package game

import "image/color"

const (
	WindowWidth  = 800
	WindowHeight = 600
	WindowTitle  = "Морской бой"
)

const (
	createGameText = "Создать игру"
	joinGameText   = "Присоединиться"
)

var (
	GrayColor  = color.Gray{100}
	GreenColor = color.RGBA{0, 200, 0, 255}
)

const (
	backgroundAnimationSpeed = 4
)

const (
	yLargeFontOffset  = 63
	yMediumFontOffset = 28
)