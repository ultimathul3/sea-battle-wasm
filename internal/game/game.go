package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ultimathul3/sea-battle-wasm/assets"
	"github.com/ultimathul3/sea-battle-wasm/internal/background"
	"github.com/ultimathul3/sea-battle-wasm/internal/button"
	"github.com/ultimathul3/sea-battle-wasm/internal/config"
	"github.com/ultimathul3/sea-battle-wasm/internal/state"
	"github.com/ultimathul3/sea-battle-wasm/internal/text"
	"github.com/ultimathul3/sea-battle-wasm/internal/touch"
)

type Game struct {
	assets     *assets.Assets
	background *background.Background
	text       *text.Text
	touch      *touch.Touch
	state      state.State
	cfg        *config.Config

	createGameButton *button.Button
	joinGameButton   *button.Button
	backButton       *button.Button
}

func New(cfg *config.Config) *Game {
	game := &Game{
		assets: assets.New(),
		touch:  touch.New(),
		state:  state.Menu,
		cfg:    cfg,
	}

	game.background = background.New(game.assets.BackgroundImages, backgroundAnimationSpeed)
	game.text = text.New(game.assets.LargeFont, game.assets.MediumFont, yLargeFontOffset, yMediumFontOffset)

	game.createGameButton = button.New(game.text, game.touch, game.assets.ButtonTickPlayer, createGameText, GrayColor, GreenColor)
	game.joinGameButton = button.New(game.text, game.touch, game.assets.ButtonTickPlayer, joinGameText, GrayColor, GreenColor)
	game.backButton = button.New(game.text, game.touch, game.assets.ButtonTickPlayer, backButtonText, GrayColor, DarkGreenColor)

	return game
}

func (g *Game) Update() error {
	g.background.Update()

	switch g.state {
	case state.Menu:
		g.createGameButton.Update(func() {
			g.state = state.CreateGame
		})
		g.joinGameButton.Update(func() {
			g.state = state.JoinGame
		})

	case state.CreateGame:
		g.backButton.Update(func() {
			g.state = state.Menu
		})

	case state.JoinGame:
		g.backButton.Update(func() {
			g.state = state.Menu
		})
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)
	switch g.state {
	case state.Menu:
		g.text.DrawLarge(screen, WindowTitle, 9, 9, color.White)
		g.createGameButton.Draw(screen, 9, 250)
		g.joinGameButton.Draw(screen, 9, 300)

	case state.CreateGame:
		g.backButton.Draw(screen, 9, 9)

	case state.JoinGame:
		g.backButton.Draw(screen, 9, 9)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
