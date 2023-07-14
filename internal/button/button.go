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

type Texter interface {
	DrawMedium(screen *ebiten.Image, text string, x, y int, color color.Color)
}

type Toucher interface {
	IsTouched() (int, int, bool)
}

type TickPlayer interface {
	Play()
	Rewind() error
}

type Button struct {
	x, y         int
	texter       Texter
	label        string
	color        color.Color
	hoverColor   color.Color
	currentColor color.Color
	width        int
	toucher      Toucher
	tickPlayer   TickPlayer
	tickPlayed   bool
}

func New(text Texter, touch Toucher, tickPlayer TickPlayer, label string, color, hoverColor color.Color) *Button {
	button := &Button{
		texter:     text,
		label:      label,
		color:      color,
		hoverColor: hoverColor,
		toucher:    touch,
		tickPlayer: tickPlayer,
	}

	button.width = utf8.RuneCountInString(button.label) * charWidth
	button.currentColor = color

	return button
}

func (b *Button) Update(callback func()) {
	if b.IsTouchHovered() {
		callback()
	}

	isHovered := b.IsHovered()
	if isHovered {
		if !b.tickPlayed {
			b.tickPlayer.Rewind()
			b.tickPlayer.Play()
			b.tickPlayed = true
		}
		b.currentColor = b.hoverColor
	} else {
		b.tickPlayed = false
		b.currentColor = b.color
	}
}

func (b *Button) Draw(screen *ebiten.Image, x, y int) {
	b.x, b.y = x, y
	b.texter.DrawMedium(screen, b.label, x, y, b.currentColor)
}

func (b *Button) IsHovered() bool {
	mx, my := ebiten.CursorPosition()
	return mx > b.x && mx < b.x+b.width && my > b.y && my < b.y+buttonHeight
}

func (b *Button) IsTouchHovered() bool {
	tx, ty, isTouched := b.toucher.IsTouched()
	return isTouched && tx > b.x && tx < b.x+b.width && ty > b.y && ty < b.y+buttonHeight
}
