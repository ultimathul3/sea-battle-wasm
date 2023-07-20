package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ultimathul3/sea-battle-wasm/internal/button"
)

func (g *Game) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)

	switch g.state {
	case MenuState:
		g.drawMenuState(screen)
	case CreateGameState:
		g.drawCreateGameState(screen)
	case JoinGameState:
		g.drawJoinGameState(screen)
	case GameCreatedState:
		g.drawGameCreatedState(screen)
	case JoinPlacementState:
		g.drawJoinPlacementState(screen)
	case OpponentGameStartedState:
		g.drawOpponentGameStartedState(screen)
	case HostWaitOpponentState:
		g.drawHostWaitOpponentState(screen)
	case HostGameStartedState:
		g.drawHostGameStartedState(screen)
	}
}

func (g *Game) drawMenuState(screen *ebiten.Image) {
	g.text.DrawLarge(screen, WindowTitle, 9, 9, color.White)
	g.createGameButton.Draw(screen, 9, 250)
	g.joinGameButton.Draw(screen, 9, 300)
}

func (g *Game) drawCreateGameState(screen *ebiten.Image) {
	g.backButton.Draw(screen, 9, 9)

	g.field.Draw(screen)
	g.opponentField.Draw(screen)

	if g.field.GetNumOfAvailableShips() == 0 {
		g.startButton.Draw(screen, 529, 519)
	}
}

func (g *Game) drawJoinGameState(screen *ebiten.Image) {
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

	g.updateButton.Draw(screen, 320, 100)
}

func (g *Game) drawGameCreatedState(screen *ebiten.Image) {
	g.backButton.Draw(screen, 9, 9)

	g.field.Draw(screen)
	g.opponentField.Draw(screen)

	g.text.DrawMedium(screen, "Ожидание игрока...", 240, 9, color.White)
}

func (g *Game) drawJoinPlacementState(screen *ebiten.Image) {
	g.backButton.Draw(screen, 9, 9)

	g.field.Draw(screen)
	g.opponentField.Draw(screen)

	if g.field.GetNumOfAvailableShips() == 0 {
		g.startButton.Draw(screen, 529, 519)
	}
}

func (g *Game) drawOpponentGameStartedState(screen *ebiten.Image) {
	g.backButton.Draw(screen, 9, 9)

	g.field.Draw(screen)
	g.opponentField.Draw(screen)

	if g.turn == OpponentTurn {
		g.text.DrawMedium(screen, "Ваш ход", 333, 9, color.White)
	} else {
		g.text.DrawMedium(screen,
			fmt.Sprintf("Ход игрока %s", g.nickname),
			220, 9, color.White,
		)
	}
}

func (g *Game) drawHostWaitOpponentState(screen *ebiten.Image) {
	g.backButton.Draw(screen, 9, 9)

	g.field.Draw(screen)
	g.opponentField.Draw(screen)

	g.text.DrawMedium(screen,
		fmt.Sprintf("Ожидание игрока %s...", g.opponentNickname),
		160, 9, color.White,
	)
}

func (g *Game) drawHostGameStartedState(screen *ebiten.Image) {
	g.backButton.Draw(screen, 9, 9)

	g.field.Draw(screen)
	g.opponentField.Draw(screen)

	if g.turn == HostTurn {
		g.text.DrawMedium(screen, "Ваш ход", 333, 9, color.White)
	} else {
		g.text.DrawMedium(screen,
			fmt.Sprintf("Ход игрока %s", g.nickname),
			220, 9, color.White,
		)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
