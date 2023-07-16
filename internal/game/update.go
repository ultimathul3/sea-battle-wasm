package game

import (
	"github.com/ultimathul3/sea-battle-wasm/internal/button"
	"github.com/ultimathul3/sea-battle-wasm/internal/state"
)

func (g *Game) Update() error {
	g.background.Update()

	switch g.state {
	case state.Menu:
		g.createGameButton.Update(func() {
			g.state = state.CreateGame
		})
		g.joinGameButton.Update(func() {
			g.gameButtons = nil
			go g.network.GetGames(g.getGamesResponse)
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
		if err := g.updateGameButtons(); err != nil {
			g.state = state.Menu
			break
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

	return nil
}

func (g *Game) updateGameButtons() error {
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
