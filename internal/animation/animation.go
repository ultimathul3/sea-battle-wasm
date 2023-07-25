package animation

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ultimathul3/sea-battle-wasm/pkg/utils"
)

type Animation struct {
	images []*ebiten.Image
	x, y   int
	offset int
	Played bool
}

func New(images []*ebiten.Image) *Animation {
	return &Animation{
		images: images,
	}
}

func (a *Animation) Play(x, y int) {
	a.Played = true
	a.x, a.y = x, y
}

func (a *Animation) Update() {
	if !a.Played {
		return
	}

	a.offset++
	if a.offset > len(a.images)-1 {
		a.Played = false
		a.offset = 0
	}
}

func (a *Animation) Draw(screen *ebiten.Image) {
	if a.Played {
		utils.DrawInCoords(screen, a.images[a.offset], a.x, a.y)
	}
}
