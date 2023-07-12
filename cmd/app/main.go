package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ultimathul3/sea-battle-wasm/internal/game"
)

func main() {
	g := &game.Game{}

	ebiten.SetWindowSize(game.WindowWidth, game.WindowHeight)
	ebiten.SetWindowTitle(game.WindowTitle)
	ebiten.RunGame(g)
}
