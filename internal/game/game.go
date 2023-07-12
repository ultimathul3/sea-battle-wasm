package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ultimathul3/sea-battle-wasm/assets"
)

const (
	WindowWidth  = 800
	WindowHeight = 600
	WindowTitle  = "Sea battle"
)

type Game struct {
	Assets *assets.Assets
}

func New() *Game {
	return &Game{
		Assets: assets.New(),
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Assets.BackgroundImage, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
