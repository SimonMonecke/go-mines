package game

import "github.com/smonecke/go-mines/convert"

type Mode int

const (
	Easy Mode = iota
	Normal
	Hard
)

type Game struct {
	board       *convert.Board
	mode        Mode
	selectedRow int
	selectedCol int
}

func New(board *convert.Board, mode Mode) *Game {
	return &Game{board: board, mode: mode}
}

func (g *Game) Mode() Mode {
	return g.mode
}

func (g *Game) Width() int {
	return g.board.Width()
}

func (g *Game) Height() int {
	return g.board.Height()
}

func (g *Game) GetSquare(col, row int) *convert.Square {
	return g.board.GetSquare(col, row)
}

func (g *Game) SelectedRow() int {
	return g.selectedRow
}

func (g *Game) SelectedCol() int {
	return g.selectedCol
}

func (g *Game) IsWon() bool {
	return g.board.UncoveredMines() == 0 && g.board.CoveredNonMines() == 0
}

func (g *Game) IsLost() bool {
	return g.board.UncoveredMines() != 0
}

func (g *Game) Uncover() {
	square := g.board.GetSquare(g.selectedCol, g.selectedRow)
	square.Uncover()
	if !square.IsMine() && square.CountOfMinesNeighbour() == 0 {
		g.uncoverSquaresWithoutMineNeighbours()
	}
}

func (g *Game) uncoverSquaresWithoutMineNeighbours() {
	countOfUncoveredSquares := 0
	for {
		countOfUncoveredSquares = 0
		for row := 0; row < g.board.Height(); row++ {
			for col := 0; col < g.board.Width(); col++ {
				square := g.board.GetSquare(col, row)
				if square.IsUncovered() && square.CountOfMinesNeighbour() == 0 {
					countOfUncoveredSquares += g.uncoverAllNeighbours(row, col)
				}
			}
		}
		if countOfUncoveredSquares == 0 {
			break
		}
	}
}

func (g *Game) uncoverAllNeighbours(row, col int) int {
	countOfUncoveredSquares := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) && col+j >= 0 && col+j < g.board.Width() && row+i >= 0 && row+i < g.board.Height() {
				square := g.board.GetSquare(col+j, row+i)
				if !square.IsUncovered() {
					square.Uncover()
					countOfUncoveredSquares++
				}
			}
		}
	}
	return countOfUncoveredSquares
}

func (g *Game) SelectSquareBottom() {
	if g.selectedRow != g.board.Height()-1 {
		g.selectedRow += 1
	}
}

func (g *Game) SelectSquareTop() {
	if g.selectedRow != 0 {
		g.selectedRow -= 1
	}
}

func (g *Game) SelectSquareRight() {
	if g.selectedCol != g.board.Width()-1 {
		g.selectedCol += 1
	}
}

func (g *Game) SelectSquareLeft() {
	if g.selectedCol != 0 {
		g.selectedCol -= 1
	}
}
