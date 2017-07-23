package convert

import (
	"fmt"
	"strings"
)

type Board struct {
	squareMap [][]Square
}

func (b *Board) String() string {
	rows := []string{}

	for row := range b.squareMap {
		cols := []string{}
		for col := range b.squareMap[row] {
			cols = append(cols, fmt.Sprintf("%s", b.squareMap[row][col]))
		}
		rows = append(rows, strings.Join(cols, " "))
	}
	return strings.Join(rows, "\n")
}

func (b *Board) Width() int {
	return len(b.squareMap[0])
}

func (b *Board) Height() int {
	return len(b.squareMap)
}

func (b *Board) UncoveredMines() int {
	uncoveredMines := 0
	for row := range b.squareMap {
		for col := range b.squareMap[row] {
			if b.squareMap[row][col].isMine && b.squareMap[row][col].isUncovered {
				uncoveredMines++
			}
		}
	}
	return uncoveredMines
}

func (b *Board) CoveredNonMines() int {
	coveredNonMines := 0
	for row := range b.squareMap {
		for col := range b.squareMap[row] {
			if !b.squareMap[row][col].isMine && !b.squareMap[row][col].isUncovered {
				coveredNonMines++
			}
		}
	}
	return coveredNonMines
}

func (b *Board) GetSquare(x, y int) *Square {
	if y < 0 || x < 0 || y >= len(b.squareMap) || x >= len(b.squareMap[0]) {
		panic("illegal coordinates")
	}
	return &b.squareMap[y][x]
}
