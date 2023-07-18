package game

import (
	"os"

	"github.com/ultimathul3/sea-battle-wasm/assets"
	"github.com/ultimathul3/sea-battle-wasm/internal/background"
	"github.com/ultimathul3/sea-battle-wasm/internal/button"
	"github.com/ultimathul3/sea-battle-wasm/internal/config"
	"github.com/ultimathul3/sea-battle-wasm/internal/field"
	"github.com/ultimathul3/sea-battle-wasm/internal/network"
	"github.com/ultimathul3/sea-battle-wasm/internal/text"
	"github.com/ultimathul3/sea-battle-wasm/pkg/utils"
)

type Game struct {
	assets     *assets.Assets
	background *background.Background
	text       *text.Text
	touch      *utils.Touch
	state      GameState
	cfg        *config.Config
	network    *network.Network

	field         *field.Field
	opponentField *field.Field

	createGameButton *button.Button
	joinGameButton   *button.Button
	backButton       *button.Button
	leftArrowButton  *button.Button
	rightArrowButton *button.Button
	updateButton     *button.Button
	startButton      *button.Button

	gameButtons         []*button.Button
	gameButtonsOffset   int
	gameButtonsPageSize int

	getGamesResponse   chan network.GetGamesResponse
	createGameResponse chan network.CreateGameResponse

	nickname string
}

func New(cfg *config.Config) *Game {
	g := &Game{
		assets:              assets.New(),
		touch:               utils.NewTouch(),
		state:               MenuState,
		cfg:                 cfg,
		gameButtonsPageSize: 4,
		getGamesResponse:    make(chan network.GetGamesResponse),
	}

	g.background = background.New(g.assets.BackgroundImages, backgroundAnimationSpeed)
	g.text = text.New(g.assets.LargeFont, g.assets.MediumFont, yLargeFontOffset, yMediumFontOffset)
	g.network = network.New(g.cfg.HttpServer.Host, g.cfg.HttpServer.Port)

	g.field = field.New(
		38, 129,
		g.assets,
		TransparentColor, g.text, g.touch,
		field.PlacementState,
	)

	g.opponentField = field.New(
		418, 129,
		g.assets,
		TransparentColor, g.text, g.touch,
		field.CurtainState,
	)

	if len(os.Args) < 2 {
		g.nickname = g.cfg.DevelopmentNickname
	} else {
		g.nickname = os.Args[1]
	}

	g.createGameButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, createGameText, GrayColor, GreenColor)
	g.joinGameButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, joinGameText, GrayColor, GreenColor)
	g.backButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, backButtonText, GrayColor, DarkGreenColor)
	g.leftArrowButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, leftArrowButtonText, LightGrayColor, DarkGreenColor)
	g.rightArrowButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, rightArrowButtonText, LightGrayColor, DarkGreenColor)
	g.updateButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, updateButtonText, LightGrayColor, DarkGreenColor)
	g.startButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, startButtonText, LightGrayColor, DarkGreenColor)

	return g
}
