package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Touch struct {
	touchIDs []ebiten.TouchID
}

func NewTouch() *Touch {
	return &Touch{}
}

func (t *Touch) IsTouched() (int, int, bool) {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		return x, y, true
	}

	touchIDs := inpututil.AppendJustPressedTouchIDs(t.touchIDs[:0])
	for _, id := range touchIDs {
		x, y := ebiten.TouchPosition(id)
		return x, y, true
	}

	return -1, -1, false
}
