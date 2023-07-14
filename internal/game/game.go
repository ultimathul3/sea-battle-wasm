package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ultimathul3/sea-battle-wasm/assets"
	"github.com/ultimathul3/sea-battle-wasm/internal/background"
	"github.com/ultimathul3/sea-battle-wasm/internal/button"
	"github.com/ultimathul3/sea-battle-wasm/internal/config"
	"github.com/ultimathul3/sea-battle-wasm/internal/network"
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
	network    *network.Network

	createGameButton *button.Button
	joinGameButton   *button.Button
	backButton       *button.Button
	leftArrowButton  *button.Button
	rightArrowButton *button.Button
	updateButton     *button.Button

	gameButtons         []*button.Button
	gameButtonsOffset   int
	gameButtonsPageSize int
}

func New(cfg *config.Config) *Game {
	game := &Game{
		assets:              assets.New(),
		touch:               touch.New(),
		state:               state.Menu,
		cfg:                 cfg,
		gameButtonsPageSize: 4,
	}

	game.background = background.New(game.assets.BackgroundImages, backgroundAnimationSpeed)
	game.text = text.New(game.assets.LargeFont, game.assets.MediumFont, yLargeFontOffset, yMediumFontOffset)
	game.network = network.New(game.cfg.HttpServer.Host, game.cfg.HttpServer.Port)

	game.createGameButton = button.New(game.text, game.touch, game.assets.ButtonTickPlayer, createGameText, GrayColor, GreenColor)
	game.joinGameButton = button.New(game.text, game.touch, game.assets.ButtonTickPlayer, joinGameText, GrayColor, GreenColor)
	game.backButton = button.New(game.text, game.touch, game.assets.ButtonTickPlayer, backButtonText, GrayColor, DarkGreenColor)
	game.leftArrowButton = button.New(game.text, game.touch, game.assets.ButtonTickPlayer, leftArrowButtonText, LightGrayColor, DarkGreenColor)
	game.rightArrowButton = button.New(game.text, game.touch, game.assets.ButtonTickPlayer, rightArrowButtonText, LightGrayColor, DarkGreenColor)
	game.updateButton = button.New(game.text, game.touch, game.assets.ButtonTickPlayer, updateButtonText, LightGrayColor, DarkGreenColor)

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
			if err := g.updateGameButtons(); err != nil {
				g.state = state.Menu
				return
			}
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
		for _, btn := range g.gameButtons {
			btn.Update(nil)
		}
		if g.gameButtonsOffset != 0 {
			g.leftArrowButton.Update(func() {
				g.gameButtonsOffset--
			})
		}
		if g.gameButtonsOffset != len(g.gameButtons)-g.gameButtonsPageSize {
			g.rightArrowButton.Update(func() {
				g.gameButtonsOffset++
			})
		}
		g.updateButton.Update(func() {
			g.gameButtonsOffset = 0
			if err := g.updateGameButtons(); err != nil {
				g.state = state.Menu
			}
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
		from := g.gameButtonsOffset
		to := g.gameButtonsOffset + g.gameButtonsPageSize
		if to > len(g.gameButtons) {
			to = len(g.gameButtons)
		}
		for i, btn := range g.gameButtons[from:to] {
			btn.Draw(screen, 350, 200+i*button.ButtonHeight)
		}
		if len(g.gameButtons) > g.gameButtonsPageSize {
			if g.gameButtonsOffset != 0 {
				g.leftArrowButton.Draw(screen, 380, 150)
			}
			if g.gameButtonsOffset != len(g.gameButtons)-g.gameButtonsPageSize {
				g.rightArrowButton.Draw(screen, 420, 150)
			}
		}
		g.updateButton.Draw(screen, 340, 100)
	}
}

func (g *Game) updateGameButtons() error {
	g.gameButtons = nil

	games, err := g.network.GetGames()
	if err != nil {
		return err
	}

	for _, game := range games {
		g.gameButtons = append(g.gameButtons, button.New(g.text, g.touch, g.assets.ButtonTickPlayer, game, GrayColor, GreenColor))
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
