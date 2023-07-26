package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ultimathul3/sea-battle-wasm/internal/button"
)

func (g *Game) Draw(screen *ebiten.Image) {
	select {
	case <-g.loadChannel:
	default:
		ebitenutil.DebugPrint(screen, LoadText)
		return
	}

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
	case HostWonState:
		g.drawHostWonState(screen)
	case OpponentWonState:
		g.drawOpponentWonState(screen)
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
		btn.DrawInCenter(screen, 200+i*button.ButtonHeight)
	}

	if len(g.gameButtons) > g.gameButtonsPageSize {
		if g.gameButtonsOffset != 0 {
			g.leftArrowButton.Draw(screen, 380, 150)
		}
		if g.gameButtonsOffset != len(g.gameButtons)-g.gameButtonsPageSize {
			g.rightArrowButton.Draw(screen, 420, 150)
		}
	}

	g.updateButton.DrawInCenter(screen, 100)
}

func (g *Game) drawGameCreatedState(screen *ebiten.Image) {
	g.backButton.Draw(screen, 9, 9)

	g.field.Draw(screen)
	g.opponentField.Draw(screen)

	g.text.DrawMediumInCenter(screen, PlayerWaitingText, 9, color.White)
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
		g.text.DrawMediumInCenter(screen, YourTurnText, 9, color.White)
	} else {
		g.text.DrawMediumInCenter(screen, fmt.Sprintf(PlayerTurnTextFmt, g.hostNickname), 9, color.White)
	}
}

func (g *Game) drawHostWaitOpponentState(screen *ebiten.Image) {
	g.backButton.Draw(screen, 9, 9)

	g.field.Draw(screen)
	g.opponentField.Draw(screen)

	g.text.DrawMediumInCenter(screen,
		fmt.Sprintf(PlayerWaitingTextFmt, g.opponentNickname),
		9, color.White,
	)
}

func (g *Game) drawHostGameStartedState(screen *ebiten.Image) {
	g.backButton.Draw(screen, 9, 9)

	g.field.Draw(screen)
	g.opponentField.Draw(screen)

	if g.turn == HostTurn {
		g.text.DrawMediumInCenter(screen, YourTurnText, 9, color.White)
	} else {
		g.text.DrawMediumInCenter(screen, fmt.Sprintf(PlayerTurnTextFmt, g.opponentNickname), 9, color.White)
	}
}

func (g *Game) drawHostWonState(screen *ebiten.Image) {
	g.backButton.Draw(screen, 9, 9)

	g.field.Draw(screen)
	g.opponentField.Draw(screen)

	g.text.DrawMediumInCenter(screen, fmt.Sprintf(PlayerWonTextFmt, g.hostNickname), 9, color.White)
}

func (g *Game) drawOpponentWonState(screen *ebiten.Image) {
	g.backButton.Draw(screen, 9, 9)

	g.field.Draw(screen)
	g.opponentField.Draw(screen)

	g.text.DrawMediumInCenter(screen, fmt.Sprintf(PlayerWonTextFmt, g.opponentNickname), 9, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
