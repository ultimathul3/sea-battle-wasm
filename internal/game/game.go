package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ultimathul3/sea-battle-wasm/assets"
	"github.com/ultimathul3/sea-battle-wasm/internal/background"
)

const (
	WindowWidth  = 800
	WindowHeight = 600
	WindowTitle  = "Sea battle"
)

const (
	backgroundAnimationSpeed = 4
)

type Game struct {
	assets     *assets.Assets
	background *background.Background
}

func New() *Game {
	game := &Game{
		assets: assets.New(),
	}

	game.background = background.New(game.assets.BackgroundImages, backgroundAnimationSpeed)

	return game
}

func (g *Game) Update() error {
	g.background.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
