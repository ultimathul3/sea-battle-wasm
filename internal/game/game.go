package game

import (
	"github.com/ultimathul3/sea-battle-wasm/assets"
	"github.com/ultimathul3/sea-battle-wasm/internal/background"
	"github.com/ultimathul3/sea-battle-wasm/internal/button"
	"github.com/ultimathul3/sea-battle-wasm/internal/config"
	"github.com/ultimathul3/sea-battle-wasm/internal/field"
	"github.com/ultimathul3/sea-battle-wasm/internal/network"
	"github.com/ultimathul3/sea-battle-wasm/internal/state"
	"github.com/ultimathul3/sea-battle-wasm/internal/text"
	"github.com/ultimathul3/sea-battle-wasm/pkg/utils"
)

type Game struct {
	assets     *assets.Assets
	background *background.Background
	text       *text.Text
	touch      *utils.Touch
	state      state.State
	cfg        *config.Config
	network    *network.Network
	field      *field.Field

	createGameButton *button.Button
	joinGameButton   *button.Button
	backButton       *button.Button
	leftArrowButton  *button.Button
	rightArrowButton *button.Button
	updateButton     *button.Button

	gameButtons         []*button.Button
	gameButtonsOffset   int
	gameButtonsPageSize int

	getGamesResponse chan network.GetGamesResponse
}

func New(cfg *config.Config) *Game {
	g := &Game{
		assets:              assets.New(),
		touch:               utils.NewTouch(),
		state:               state.Menu,
		cfg:                 cfg,
		gameButtonsPageSize: 4,
		getGamesResponse:    make(chan network.GetGamesResponse),
	}

	g.background = background.New(g.assets.BackgroundImages, backgroundAnimationSpeed)
	g.text = text.New(g.assets.LargeFont, g.assets.MediumFont, yLargeFontOffset, yMediumFontOffset)
	g.network = network.New(g.cfg.HttpServer.Host, g.cfg.HttpServer.Port)
	g.field = field.New(
		38, 129,
		g.assets.SingleDeckShipImage, g.assets.DoubleDeckShipImage, g.assets.ThreeDeckShipImage, g.assets.FourDeckShipImage,
		g.assets.FieldImage, g.assets.SelectImage,
		TransparentColor,
	)

	g.createGameButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, createGameText, GrayColor, GreenColor)
	g.joinGameButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, joinGameText, GrayColor, GreenColor)
	g.backButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, backButtonText, GrayColor, DarkGreenColor)
	g.leftArrowButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, leftArrowButtonText, LightGrayColor, DarkGreenColor)
	g.rightArrowButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, rightArrowButtonText, LightGrayColor, DarkGreenColor)
	g.updateButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, updateButtonText, LightGrayColor, DarkGreenColor)

	return g
}
