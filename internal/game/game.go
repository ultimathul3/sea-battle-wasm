package game

import "github.com/hajimehoshi/ebiten/v2"

const (
	WindowWidth  = 800
	WindowHeight = 600
	WindowTitle  = "Sea battle"
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
