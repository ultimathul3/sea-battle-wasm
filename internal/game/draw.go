package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ultimathul3/sea-battle-wasm/internal/button"
	"github.com/ultimathul3/sea-battle-wasm/internal/state"
)

func (g *Game) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)
	switch g.state {
	case state.Menu:
		g.text.DrawLarge(screen, WindowTitle, 9, 9, color.White)
		g.createGameButton.Draw(screen, 9, 250)
		g.joinGameButton.Draw(screen, 9, 300)

	case state.CreateGame:
		g.backButton.Draw(screen, 9, 9)

	case state.JoinGame:
		g.backButton.Draw(screen, 9, 9)
		from := g.gameButtonsOffset
		to := g.gameButtonsOffset + g.gameButtonsPageSize
		if to > len(g.gameButtons) {
			to = len(g.gameButtons)
		}
		for i, btn := range g.gameButtons[from:to] {
			btn.Draw(screen, 350, 200+i*button.ButtonHeight)
		}
		if len(g.gameButtons) > g.gameButtonsPageSize {
			if g.gameButtonsOffset != 0 {
				g.leftArrowButton.Draw(screen, 380, 150)
			}
			if g.gameButtonsOffset != len(g.gameButtons)-g.gameButtonsPageSize {
				g.rightArrowButton.Draw(screen, 420, 150)
			}
		}
		g.updateButton.Draw(screen, 340, 100)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
