package text

import (
	"image/color"
	"unicode/utf8"

	"github.com/hajimehoshi/ebiten/v2"
	etext "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Text struct {
	largeFont                   font.Face
	mediumFont                  font.Face
	yLargeFontOffset            int
	yMediumFontOffset           int
	yMediumFontCharWidth        int
	yMediumFontSizeBetweenChars int
}

func New(largeFont, mediumFont font.Face, yLargeFontOffset, yMediumFontOffset, yMediumFontCharWidth, yMediumFontSizeBetweenChars int) *Text {
	return &Text{
		largeFont:                   largeFont,
		mediumFont:                  mediumFont,
		yLargeFontOffset:            yLargeFontOffset,
		yMediumFontOffset:           yMediumFontOffset,
		yMediumFontCharWidth:        yMediumFontCharWidth,
		yMediumFontSizeBetweenChars: yMediumFontSizeBetweenChars,
	}
}

func (t *Text) DrawLarge(screen *ebiten.Image, text string, x, y int, color color.Color) {
	etext.Draw(screen, text, t.largeFont, x, y+t.yLargeFontOffset, color)
}

func (t *Text) DrawMedium(screen *ebiten.Image, text string, x, y int, color color.Color) {
	etext.Draw(screen, text, t.mediumFont, x, y+t.yMediumFontOffset, color)
}

func (t *Text) DrawMediumInCenter(screen *ebiten.Image, text string, y int, color color.Color) int {
	charCount := utf8.RuneCountInString(text)
	x := (screen.Bounds().Max.X - charCount*(t.yMediumFontCharWidth+t.yMediumFontSizeBetweenChars-1)) / 2
	etext.Draw(screen, text, t.mediumFont, x, y+t.yMediumFontOffset, color)
	return x
}
