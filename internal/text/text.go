package text

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	etext "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Text struct {
	largeFont         font.Face
	mediumFont        font.Face
	yLargeFontOffset  int
	yMediumFontOffset int
}

func New(largeFont, mediumFont font.Face, yLargeFontOffset, yMediumFontOffset int) *Text {
	return &Text{
		largeFont:         largeFont,
		mediumFont:        mediumFont,
		yLargeFontOffset:  yLargeFontOffset,
		yMediumFontOffset: yMediumFontOffset,
	}
}

func (t *Text) DrawLarge(screen *ebiten.Image, text string, x, y int, color color.Color) {
	etext.Draw(screen, text, t.largeFont, x, y+t.yLargeFontOffset, color)
}

func (t *Text) DrawMedium(screen *ebiten.Image, text string, x, y int, color color.Color) {
	etext.Draw(screen, text, t.mediumFont, x, y+t.yMediumFontOffset, color)
}
