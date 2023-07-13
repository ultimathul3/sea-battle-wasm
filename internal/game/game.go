package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ultimathul3/sea-battle-wasm/assets"
	"github.com/ultimathul3/sea-battle-wasm/internal/background"
	"github.com/ultimathul3/sea-battle-wasm/internal/text"
)

type Game struct {
	assets     *assets.Assets
	background *background.Background
	text       *text.Text
}

func New() *Game {
	game := &Game{
		assets: assets.New(),
	}

	game.background = background.New(game.assets.BackgroundImages, backgroundAnimationSpeed)
	game.text = text.New(game.assets.LargeFont, game.assets.MediumFont, yLargeFontOffset, yMediumFontOffset)

	return game
}

func (g *Game) Update() error {
	g.background.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)
	g.text.DrawLarge(screen, WindowTitle, 9, 9, color.White)
	g.text.DrawMedium(screen, createGameText, 9, 250, GrayColor)
	g.text.DrawMedium(screen, joinGameText, 9, 300, GrayColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
