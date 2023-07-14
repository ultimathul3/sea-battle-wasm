package game

import (
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
