package button

import (
	"image/color"
	"unicode/utf8"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ButtonHeight = 42
	charWidth    = 20
)

type Texter interface {
	DrawMedium(screen *ebiten.Image, text string, x, y int, color color.Color)
	DrawMediumInCenter(screen *ebiten.Image, text string, y int, color color.Color) int
}

type Toucher interface {
	IsTouched() (int, int, bool)
}

type TickPlayer interface {
	Play()
	Rewind() error
}

type Button struct {
	Label string

	x, y         int
	texter       Texter
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
		Label:      label,
		color:      color,
		hoverColor: hoverColor,
		toucher:    touch,
		tickPlayer: tickPlayer,
	}

	button.width = utf8.RuneCountInString(button.Label) * charWidth
	button.currentColor = color

	return button
}

func (b *Button) Update(callback func()) {
	if b.IsTouchHovered() && callback != nil {
		callback()
	}

	isHovered := b.IsHovered()
	if isHovered {
		if !b.tickPlayed {
			if err := b.tickPlayer.Rewind(); err != nil {
				return
			}
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
	b.texter.DrawMedium(screen, b.Label, x, y, b.currentColor)
}

func (b *Button) DrawInCenter(screen *ebiten.Image, y int) {
	b.y = y
	b.x = b.texter.DrawMediumInCenter(screen, b.Label, y, b.currentColor)
}

func (b *Button) IsHovered() bool {
	mx, my := ebiten.CursorPosition()
	return mx > b.x && mx < b.x+b.width && my > b.y && my < b.y+ButtonHeight
}

func (b *Button) IsTouchHovered() bool {
	tx, ty, isTouched := b.toucher.IsTouched()
	return isTouched && tx > b.x && tx < b.x+b.width && ty > b.y && ty < b.y+ButtonHeight
}
