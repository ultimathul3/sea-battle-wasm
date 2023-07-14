package button

import (
	"image/color"
	"unicode/utf8"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	charWidth    = 20
	buttonHeight = 42
)

type Text interface {
	DrawMedium(screen *ebiten.Image, text string, x, y int, color color.Color)
}

type Touch interface {
	IsTouched() (int, int, bool)
}

type Button struct {
	x, y       int
	text       Text
	label      string
	color      color.Color
	hoverColor color.Color
	width      int
	touch      Touch
}

func New(x, y int, text Text, touch Touch, label string, color, hoverColor color.Color) *Button {
	button := &Button{
		x:          x,
		y:          y,
		text:       text,
		label:      label,
		color:      color,
		hoverColor: hoverColor,
		touch:      touch,
	}

	button.width = utf8.RuneCountInString(button.label) * charWidth

	return button
}

func (b *Button) Update(callback func()) {
	_, _, isTouched := b.touch.IsTouched()
	if b.IsHovered() && isTouched {
		callback()
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	if b.IsHovered() {
		b.text.DrawMedium(screen, b.label, b.x, b.y, b.hoverColor)
	} else {
		b.text.DrawMedium(screen, b.label, b.x, b.y, b.color)
	}
}

func (b *Button) IsHovered() bool {
	mx, my := ebiten.CursorPosition()
	return mx > b.x && mx < b.x+b.width && my > b.y && my < b.y+buttonHeight
}
