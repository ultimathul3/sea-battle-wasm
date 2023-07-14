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
