package game

import (
	"log"

	"github.com/ultimathul3/sea-battle-wasm/internal/button"
	"github.com/ultimathul3/sea-battle-wasm/internal/field"
	"github.com/ultimathul3/sea-battle-wasm/internal/network"
)

func (g *Game) Update() error {
	select {
	case <-g.loadChannel:
	default:
		return nil
	}

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
				HostNickname: g.hostNickname,
				HostField:    stringField,
			}, g.createGameResponse)
		})
	}
}

func (g *Game) updateJoinGameState() {
	g.backButton.Update(func() {
		g.state = MenuState
	})

	if r, err := g.updateGetGamesResponse(); err != nil {
		g.handleErrorResponse(err)
	} else if r != nil {
		for _, game := range r.Games {
			g.gameButtons = append(g.gameButtons, button.New(g.text, g.touch, g.assets.ButtonTickPlayer, game, GrayColor, GreenColor))
		}
	}

	for _, btn := range g.gameButtons {
		btn.Update(func() {
			g.state = JoinPlacementState
			g.hostNickname = btn.Label
			go g.network.JoinGame(network.JoinGameRequest{
				HostNickname:     g.hostNickname,
				OpponentNickname: g.opponentNickname,
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

	if r, err := g.updateCreateGameResponse(); err != nil {
		g.handleErrorResponse(err)
	} else if r != nil {
		g.hostUuid = r.HostUuid
		go g.network.Wait(network.WaitRequest{Uuid: g.hostUuid}, g.waitResponse)
	}

	if r, err := g.updateWaitResponse(); err != nil {
		g.handleErrorResponse(err)
	} else if r != nil {
		if r.Status == network.WaitingForOpponentStatus {
			g.state = HostWaitOpponentState
			g.opponentNickname = r.Message
			go g.network.Wait(network.WaitRequest{Uuid: g.hostUuid}, g.waitResponse)
		}
	}
}

func (g *Game) updateJoinPlacementState() {
	g.backButton.Update(func() {
		g.state = MenuState
	})

	g.field.Update()

	if r, err := g.updateJoinGameResponse(); err != nil {
		g.handleErrorResponse(err)
	} else if r != nil {
		g.opponentUuid = r.OpponentUuid
	}

	if g.field.GetNumOfAvailableShips() == 0 {
		g.startButton.Update(func() {
			g.field.SetState(field.PlacementFinishedState)
			stringField := g.field.ConvertFieldRuneMatrixToString()
			go g.network.StartGame(network.StartGameRequest{
				HostNickname:  g.hostNickname,
				OpponentField: stringField,
				OpponentUuid:  g.opponentUuid,
			}, g.startGameResponse)
		})
	}

	if r, err := g.updateStartGameResponse(); err != nil {
		g.handleErrorResponse(err)
	} else if r != nil {
		g.state = OpponentGameStartedState
		g.opponentField.SetState(field.ShootState)
		go g.network.Wait(network.WaitRequest{Uuid: g.opponentUuid}, g.waitResponse)
	}
}

func (g *Game) updateOpponentGameStartedState() {
	g.backButton.Update(func() {
		g.state = MenuState
	})

	g.field.Update()
	g.opponentField.Update()

	x, y, isTouched := g.opponentField.IsEmptyCellTouched()
	if isTouched && g.turn == OpponentTurn && !g.isShot {
		g.lastX = x
		g.lastY = y
		g.isShot = true
		go g.network.Shoot(network.ShootRequest{
			HostNickname: g.hostNickname,
			X:            uint32(x),
			Y:            uint32(y),
			Uuid:         g.opponentUuid,
		}, g.shootResponse)
	}

	if r, err := g.updateShootResponse(); err != nil {
		g.handleErrorResponse(err)
	} else if r != nil {
		if r.Status == network.OpponentMissStatus {
			go g.network.Wait(network.WaitRequest{Uuid: g.opponentUuid}, g.waitResponse)
			g.opponentField.SetMissCell(g.lastX, g.lastY)
			g.turn = HostTurn
		} else if r.Status == network.OpponentHitStatus {
			g.opponentField.SetHitCell(g.lastX, g.lastY)
		}
		g.isShot = false
	}

	if r, err := g.updateWaitResponse(); err != nil {
		g.handleErrorResponse(err)
	} else if r != nil {
		if r.Status == network.HostMissStatus {
			g.field.SetMissCell(int(r.X), int(r.Y))
			g.turn = OpponentTurn
		} else if r.Status == network.HostHitStatus {
			go g.network.Wait(network.WaitRequest{Uuid: g.opponentUuid}, g.waitResponse)
			g.field.SetHitCell(int(r.X), int(r.Y))
		}
	}
}

func (g *Game) updateHostWaitOpponentState() {
	g.backButton.Update(func() {
		g.state = MenuState
	})

	if r, err := g.updateWaitResponse(); err != nil {
		g.handleErrorResponse(err)
	} else if r != nil {
		if r.Status == network.GameStartedStatus {
			g.state = HostGameStartedState
			g.opponentField.SetState(field.ShootState)
		}
	}

	g.field.Update()
	g.opponentField.Update()
}

func (g *Game) updateHostGameStartedState() {
	g.backButton.Update(func() {
		g.state = MenuState
	})

	g.field.Update()
	g.opponentField.Update()

	x, y, isTouched := g.opponentField.IsEmptyCellTouched()
	if isTouched && g.turn == HostTurn && !g.isShot {
		g.lastX = x
		g.lastY = y
		g.isShot = true
		go g.network.Shoot(network.ShootRequest{
			HostNickname: g.hostNickname,
			X:            uint32(x),
			Y:            uint32(y),
			Uuid:         g.hostUuid,
		}, g.shootResponse)
	}

	if r, err := g.updateShootResponse(); err != nil {
		g.handleErrorResponse(err)
	} else if r != nil {
		if r.Status == network.HostMissStatus {
			go g.network.Wait(network.WaitRequest{Uuid: g.hostUuid}, g.waitResponse)
			g.opponentField.SetMissCell(g.lastX, g.lastY)
			g.turn = OpponentTurn
		} else if r.Status == network.HostHitStatus {
			g.opponentField.SetHitCell(g.lastX, g.lastY)
		}
		g.isShot = false
	}

	if r, err := g.updateWaitResponse(); err != nil {
		g.handleErrorResponse(err)
	} else if r != nil {
		if r.Status == network.OpponentMissStatus {
			g.field.SetMissCell(int(r.X), int(r.Y))
			g.turn = HostTurn
		} else if r.Status == network.OpponentHitStatus {
			go g.network.Wait(network.WaitRequest{Uuid: g.hostUuid}, g.waitResponse)
			g.field.SetHitCell(int(r.X), int(r.Y))
		}
	}
}

func (g *Game) handleErrorResponse(err error) {
	log.Println(err)
	g.state = MenuState
}

func (g *Game) updateShootResponse() (*network.ShootResponse, error) {
	select {
	case r := <-g.shootResponse:
		if r.Error != nil {
			return nil, r.Error
		}
		return &r, nil
	default:
		return nil, nil
	}
}

func (g *Game) updateWaitResponse() (*network.WaitResponse, error) {
	select {
	case r := <-g.waitResponse:
		if r.Error != nil {
			return nil, r.Error
		}
		return &r, nil
	default:
		return nil, nil
	}
}

func (g *Game) updateStartGameResponse() (*network.StartGameResponse, error) {
	select {
	case r := <-g.startGameResponse:
		if r.Error != nil {
			return nil, r.Error
		}
		return &r, nil
	default:
		return nil, nil
	}
}

func (g *Game) updateJoinGameResponse() (*network.JoinGameResponse, error) {
	select {
	case r := <-g.joinGameResponse:
		if r.Error != nil {
			return nil, r.Error
		}
		return &r, nil
	default:
		return nil, nil
	}
}

func (g *Game) updateCreateGameResponse() (*network.CreateGameResponse, error) {
	select {
	case r := <-g.createGameResponse:
		if r.Error != nil {
			return nil, r.Error
		}
		return &r, nil
	default:
		return nil, nil
	}
}

func (g *Game) updateGetGamesResponse() (*network.GetGamesResponse, error) {
	select {
	case r := <-g.getGamesResponse:
		if r.Error != nil {
			return nil, r.Error
		}
		return &r, nil
	default:
		return nil, nil
	}
}
