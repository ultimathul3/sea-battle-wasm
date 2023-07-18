package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ultimathul3/sea-battle-wasm/internal/config"
	"github.com/ultimathul3/sea-battle-wasm/internal/game"
)

func main() {
	cfg := config.New()

	g := game.New(cfg)

	ebiten.SetWindowSize(game.WindowWidth, game.WindowHeight)
	ebiten.SetWindowTitle(game.WindowTitle)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
