package game

import (
	"os"
	"runtime"

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
	joinGameResponse   chan network.JoinGameResponse
	startGameResponse  chan network.StartGameResponse
	waitResponse       chan network.WaitResponse
	shootResponse      chan network.ShootResponse
	loadChannel        chan struct{}

	hostNickname     string
	opponentNickname string
	hostUuid         string
	opponentUuid     string

	turn         Turn
	lastX, lastY int
	isShot       bool
}

func New(cfg *config.Config) *Game {
	g := &Game{
		touch:               utils.NewTouch(),
		state:               MenuState,
		cfg:                 cfg,
		gameButtonsPageSize: 4,
		loadChannel:         make(chan struct{}),
	}

	go g.load()

	return g
}

func (g *Game) load() {
	runtime.Gosched()

	g.assets = assets.New()
	g.background = background.New(g.assets.BackgroundImages)
	g.text = text.New(g.assets.LargeFont, g.assets.MediumFont, yLargeFontOffset, yMediumFontOffset, yMediumFontCharWidth, yMediumFontSizeBetweenChars)
	g.network = network.New(g.cfg.HttpServer.Host)

	g.createGameButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, createGameText, GrayColor, GreenColor)
	g.joinGameButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, joinGameText, GrayColor, GreenColor)
	g.backButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, backButtonText, GrayColor, DarkGreenColor)
	g.leftArrowButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, leftArrowButtonText, LightGrayColor, DarkGreenColor)
	g.rightArrowButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, rightArrowButtonText, LightGrayColor, DarkGreenColor)
	g.updateButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, updateButtonText, LightGrayColor, DarkGreenColor)
	g.startButton = button.New(g.text, g.touch, g.assets.ButtonTickPlayer, startButtonText, LightGrayColor, DarkGreenColor)

	g.assets.ThemePlayer.Play()

	close(g.loadChannel)
}

func (g *Game) resetGame() {
	g.field = field.New(
		38, 129,
		g.assets, TransparentColor, g.text, g.touch,
		field.PlacementState,
	)

	g.opponentField = field.New(
		418, 129,
		g.assets,
		TransparentColor, g.text, g.touch,
		field.CurtainState,
	)

	if os.Args[0] != "js" {
		g.hostNickname = g.cfg.DevelopmentNickname
		g.opponentNickname = g.cfg.DevelopmentNickname
	} else {
		g.hostNickname = os.Args[1]
		g.opponentNickname = os.Args[1]
	}

	if os.Getenv("DEVELOPMENT") == "1" {
		if g.state == CreateGameState {
			g.field.SetFieldMatrix([][]rune{
				[]rune("~~~~~~~~~~~~"),
				[]rune("~          ~"),
				[]rune("~          ~"),
				[]rune("xxx     xxxx"),
				[]rune("x1x     x@◄x"),
				[]rune("xxxxxxxxxxxx"),
				[]rune("x1xx$←←◄x  ~"),
				[]rune("xxxxxxxxxxx~"),
				[]rune("x1x#←◄xx@◄x~"),
				[]rune("xxxxxxxxxxx~"),
				[]rune("x1x#←◄xx@◄x~"),
				[]rune("xxxxxxxxxxx~"),
			})
		} else if g.state == JoinGameState {
			g.field.SetFieldMatrix([][]rune{
				[]rune("~~~~~~~~~~~~"),
				[]rune("~          ~"),
				[]rune("~   xxx    ~"),
				[]rune("~xxxx2x    ~"),
				[]rune("~x1xx▲x    ~"),
				[]rune("xxxxxxx xxx~"),
				[]rune("x1x   xxx2x~"),
				[]rune("xxxxxxx4x▲x~"),
				[]rune("x1x3x3x↑xxx~"),
				[]rune("xxx↑x↑x↑x2x~"),
				[]rune("x1x▲x▲x▲x▲x~"),
				[]rune("xxxxxxxxxxx~"),
			})
		}
	}

	g.getGamesResponse = make(chan network.GetGamesResponse)
	g.createGameResponse = make(chan network.CreateGameResponse)
	g.joinGameResponse = make(chan network.JoinGameResponse)
	g.startGameResponse = make(chan network.StartGameResponse)
	g.waitResponse = make(chan network.WaitResponse)
	g.shootResponse = make(chan network.ShootResponse)

	g.gameButtons = nil
	g.gameButtonsOffset = 0

	g.turn = HostTurn
	g.isShot = false
}
