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
	})

	g.joinGameButton.Update(func() {
		g.state = JoinGameState
		g.resetGame()
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

func (g *Game) updateOpponentGameStartedState() {
	g.updateGameStartedState(OpponentTurn)
}

func (g *Game) updateHostGameStartedState() {
	g.updateGameStartedState(HostTurn)
}

func (g *Game) updateGameStartedState(turn Turn) {
	g.backButton.Update(func() {
		g.state = MenuState
	})

	g.field.Update()
	g.opponentField.Update()

	var uuid string
	var missStatus, hitStatus network.GameStatus
	var oppositeMissStatus, oppositeHitStatus network.GameStatus
	var oppositeTurn Turn

	if turn == HostTurn {
		uuid = g.hostUuid
		missStatus = network.HostMissStatus
		hitStatus = network.HostHitStatus
		oppositeTurn = OpponentTurn
		oppositeMissStatus = network.OpponentMissStatus
		oppositeHitStatus = network.OpponentHitStatus
	} else {
		uuid = g.opponentUuid
		missStatus = network.OpponentMissStatus
		hitStatus = network.OpponentHitStatus
		oppositeTurn = HostTurn
		oppositeMissStatus = network.HostMissStatus
		oppositeHitStatus = network.HostHitStatus
	}

	x, y, isTouched := g.opponentField.IsEmptyCellTouched()
	if isTouched && g.turn == turn && !g.isShot {
		g.lastX = x
		g.lastY = y
		g.isShot = true
		go g.network.Shoot(network.ShootRequest{
			HostNickname: g.hostNickname,
			X:            uint32(x),
			Y:            uint32(y),
			Uuid:         uuid,
		}, g.shootResponse)
	}

	if r, err := g.updateShootResponse(); err != nil {
		g.handleErrorResponse(err)
	} else if r != nil {
		if r.Status == missStatus {
			go g.network.Wait(network.WaitRequest{Uuid: uuid}, g.waitResponse)
			g.opponentField.SetMissCell(g.lastX, g.lastY)
			g.turn = oppositeTurn
		} else if r.Status == hitStatus {
			if r.DestroyedShip != "" {
				g.destroyShip(g.opponentField, r.DestroyedShip, int(r.X), int(r.Y))
			}
			g.opponentField.SetHitCell(g.lastX, g.lastY)
		}
		g.isShot = false
	}

	if r, err := g.updateWaitResponse(); err != nil {
		g.handleErrorResponse(err)
	} else if r != nil {
		if r.Status == oppositeMissStatus {
			g.field.SetMissCell(int(r.X), int(r.Y))
			g.turn = turn
		} else if r.Status == oppositeHitStatus {
			go g.network.Wait(network.WaitRequest{Uuid: uuid}, g.waitResponse)
			if r.DestroyedShip != "" {
				g.destroyShip(g.field, r.DestroyedShip, int(r.DestroyedX), int(r.DestroyedY))
			}
			g.field.SetHitCell(int(r.X), int(r.Y))
		}
	}
}

func (g *Game) destroyShip(f *field.Field, ship network.Ship, x, y int) {
	switch ship {
	case network.SingleDeckShip:
		f.DestroyShip(field.SingleDeckShip, x, y)
	case network.DoubleDeckShipDown:
		f.DestroyShip(field.DoubleDeckShipDown, x, y)
	case network.ThreeDeckShipDown:
		f.DestroyShip(field.ThreeDeckShipDown, x, y)
	case network.FourDeckShipDown:
		f.DestroyShip(field.FourDeckShipDown, x, y)
	case network.DoubleDeckShipRight:
		f.DestroyShip(field.DoubleDeckShipRight, x, y)
	case network.ThreeDeckShipRight:
		f.DestroyShip(field.ThreeDeckShipRight, x, y)
	case network.FourDeckShipRight:
		f.DestroyShip(field.FourDeckShipRight, x, y)
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
