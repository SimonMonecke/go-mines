package generate

import (
	"math/rand"
	"time"
)

const (
	noMine = "."
	mine   = "X"
)

func EasyStringMap() [][]string {
	return generateStringMap(8, 8, 10)
}

func NormalStringMap() [][]string {
	return generateStringMap(16, 16, 40)
}

func HardStringMap() [][]string {
	return generateStringMap(16, 30, 99)
}

func generateStringMap(width, height, mines int) [][]string {
	stringMap := initStringMapWithoutMines(width, height)
	return placeMinesRandomly(stringMap, mines)
}

func initStringMapWithoutMines(width, height int) [][]string {
	stringMap := make([][]string, height)
	for i := range stringMap {
		stringMap[i] = make([]string, width)
		for j := range stringMap[i] {
			stringMap[i][j] = noMine
		}
	}
	return stringMap
}

func placeMinesRandomly(stringMap [][]string, mines int) [][]string {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	height := len(stringMap)
	width := len(stringMap[0])

	for placedMines := 0; placedMines < mines; {
		row := randGen.Intn(height)
		col := randGen.Intn(width)
		if stringMap[row][col] == noMine {
			stringMap[row][col] = mine
			placedMines++
		}
	}
	return stringMap
}
