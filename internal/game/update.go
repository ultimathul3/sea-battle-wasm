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
	case JoinPlacementState:
		g.updateJoinPlacementState()
	case OpponentGameStartedState:
		g.updateOpponentGameStartedState()
	case HostWaitOpponentState:
		g.updateHostWaitOpponentState()
	case HostGameStartedState:
		g.updateHostGameStartedState()
	}

	return nil
}

func (g *Game) updateMenuState() {
	g.createGameButton.Update(func() {
		g.state = CreateGameState
		g.resetGame()
	})

	g.joinGameButton.Update(func() {
		g.state = JoinGameState
		g.resetGame()
		g.gameButtons = nil
		go g.network.GetGames(g.getGamesResponse)
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
			g.field.SetState(field.PlacementFinishedState)
			stringField := g.field.ConvertFieldRuneMatrixToString()
			go g.network.CreateGame(network.CreateGameRequest{
				HostNickname: g.nickname,
				HostField:    stringField,
			}, g.createGameResponse)
		})
	}
}

func (g *Game) updateJoinGameState() {
	g.backButton.Update(func() {
		g.gameButtonsOffset = 0
		g.gameButtons = nil
		g.state = MenuState
	})

	if err := g.updateGetGamesResponse(); err != nil {
		log.Println(err)
		g.state = MenuState
		return
	}

	for _, btn := range g.gameButtons {
		btn.Update(func() {
			g.state = JoinPlacementState
			g.opponentNickname = btn.Label
			go g.network.JoinGame(network.JoinGameRequest{
				HostNickname:     g.nickname,
				OpponentNickname: btn.Label,
			}, g.joinGameResponse)
		})
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

	if err := g.updateCreateGameResponse(); err != nil {
		log.Println(err)
		g.state = MenuState
		return
	}

	if err := g.updateWaitResponse(); err != nil {
		log.Println(err)
		g.state = MenuState
		return
	}
}

func (g *Game) updateJoinPlacementState() {
	g.backButton.Update(func() {
		g.state = MenuState
	})

	g.field.Update()

	if err := g.updateJoinGameResponse(); err != nil {
		log.Println(err)
		g.state = MenuState
		return
	}

	if err := g.updateStartGameResponse(); err != nil {
		log.Println(err)
		g.state = MenuState
		return
	}

	if g.field.GetNumOfAvailableShips() == 0 {
		g.startButton.Update(func() {
			g.field.SetState(field.PlacementFinishedState)
			stringField := g.field.ConvertFieldRuneMatrixToString()
			go g.network.StartGame(network.StartGameRequest{
				HostNickname:  g.nickname,
				OpponentField: stringField,
				OpponentUuid:  g.opponentUuid,
			}, g.startGameResponse)
		})
	}
}

func (g *Game) updateOpponentGameStartedState() {
	g.backButton.Update(func() {
		g.state = MenuState
	})
}

func (g *Game) updateHostWaitOpponentState() {
	g.backButton.Update(func() {
		g.state = MenuState
	})

	if err := g.updateWaitResponse(); err != nil {
		log.Println(err)
		g.state = MenuState
		return
	}
}

func (g *Game) updateHostGameStartedState() {
	g.backButton.Update(func() {
		g.state = MenuState
	})
}

func (g *Game) updateWaitResponse() error {
	select {
	case r := <-g.waitResponse:
		if r.Error != nil {
			return r.Error
		}

		switch g.state {
		case GameCreatedState:
			if r.Status == network.WaitingForOpponentStatus {
				g.state = HostWaitOpponentState
				g.opponentNickname = r.Message
				go g.network.Wait(network.WaitRequest{
					Uuid: g.hostUuid,
				}, g.waitResponse)
			}
		case HostWaitOpponentState:
			if r.Status == network.GameStartedStatus {
				g.state = HostGameStartedState
			}
		}
	default:
		return nil
	}

	return nil
}

func (g *Game) updateStartGameResponse() error {
	select {
	case r := <-g.startGameResponse:
		if r.Error != nil {
			return r.Error
		}
		g.state = OpponentGameStartedState
	default:
		return nil
	}

	return nil
}

func (g *Game) updateJoinGameResponse() error {
	select {
	case r := <-g.joinGameResponse:
		if r.Error != nil {
			return r.Error
		}
		g.opponentUuid = r.OpponentUuid
	default:
		return nil
	}

	return nil
}

func (g *Game) updateCreateGameResponse() error {
	select {
	case r := <-g.createGameResponse:
		if r.Error != nil {
			return r.Error
		}
		g.hostUuid = r.HostUuid
		go g.network.Wait(network.WaitRequest{
			Uuid: r.HostUuid,
		}, g.waitResponse)
	default:
		return nil
	}

	return nil
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
