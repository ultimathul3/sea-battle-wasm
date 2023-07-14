package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ultimathul3/sea-battle-wasm/assets"
	"github.com/ultimathul3/sea-battle-wasm/internal/background"
	"github.com/ultimathul3/sea-battle-wasm/internal/button"
	"github.com/ultimathul3/sea-battle-wasm/internal/text"
	"github.com/ultimathul3/sea-battle-wasm/internal/touch"
)

type Game struct {
	assets     *assets.Assets
	background *background.Background
	text       *text.Text
	touch      *touch.Touch

	createGameButton *button.Button
	joinGameButton   *button.Button
}

func New() *Game {
	game := &Game{
		assets: assets.New(),
		touch:  touch.New(),
	}

	game.background = background.New(game.assets.BackgroundImages, backgroundAnimationSpeed)
	game.text = text.New(game.assets.LargeFont, game.assets.MediumFont, yLargeFontOffset, yMediumFontOffset)

	game.createGameButton = button.New(9, 250, game.text, game.touch, createGameText, GrayColor, GreenColor)
	game.joinGameButton = button.New(9, 300, game.text, game.touch, joinGameText, GrayColor, GreenColor)

	return game
}

func (g *Game) Update() error {
	g.background.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)
	g.text.DrawLarge(screen, WindowTitle, 9, 9, color.White)

	g.createGameButton.Draw(screen)
	g.joinGameButton.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
