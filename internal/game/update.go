package game

import (
	"log"

	"github.com/ultimathul3/sea-battle-wasm/internal/button"
	"github.com/ultimathul3/sea-battle-wasm/internal/field"
	"github.com/ultimathul3/sea-battle-wasm/internal/network"
)

func (g *Game) Update() error {
	g.background.Update()

	switch g.state {
	case MenuState:
		g.updateMenuState()
	case CreateGameState:
		g.updateCreateGameState()
	case JoinGameState:
		g.updateJoinGameState()
	case GameCreatedState:
		g.updateGameCreatedState()
	}

	return nil
}

func (g *Game) updateMenuState() {
	g.createGameButton.Update(func() {
		g.state = CreateGameState
		g.field = field.New(
			38, 129,
			g.assets, TransparentColor, g.text, g.touch,
			field.PlacementState,
		)
	})

	g.joinGameButton.Update(func() {
		g.gameButtons = nil
		go g.network.GetGames(g.getGamesResponse)
		g.state = JoinGameState
	})
}

func (g *Game) updateCreateGameState() {
	g.backButton.Update(func() {
		g.state = MenuState
	})

	g.field.Update()

	if g.field.GetNumOfAvailableShips() == 0 {
		g.startButton.Update(func() {
			g.state = GameCreatedState
			field := g.field.ConvertFieldRuneMatrixToString()
			go g.network.CreateGame(network.CreateGameRequest{
				HostNickname: g.nickname,
				HostField:    field,
			}, g.createGameResponse)
		})
	}
}

func (g *Game) updateJoinGameState() {
	g.backButton.Update(func() {
		g.state = MenuState
	})

	if err := g.updateGetGamesResponse(); err != nil {
		log.Println(err)
		g.state = MenuState
		return
	}

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
		g.gameButtons = nil
		go g.network.GetGames(g.getGamesResponse)
	})
}

func (g *Game) updateGameCreatedState() {
	g.backButton.Update(func() {
		g.state = MenuState
	})

	g.opponentField.Update()
}

func (g *Game) updateGetGamesResponse() error {
	select {
	case r := <-g.getGamesResponse:
		if r.Error != nil {
			return r.Error
		}
		for _, game := range r.Games {
			g.gameButtons = append(g.gameButtons, button.New(g.text, g.touch, g.assets.ButtonTickPlayer, game, GrayColor, GreenColor))
		}
	default:
		return nil
	}

	return nil
}
