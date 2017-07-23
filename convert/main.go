package convert

func FromStringMapToBoard(stringMap [][]string) Board {
	squareMap := make([][]Square, len(stringMap))
	for i := range stringMap {
		squareRow := make([]Square, len(stringMap[i]))
		for j := range stringMap[i] {
			if stringMap[i][j] == "X" {
				squareRow[j] = Square{isMine: true, countOfMinesNeighbour: calculateNeighbours(i, j, stringMap), isUncovered: false}
			} else {
				squareRow[j] = Square{isMine: false, countOfMinesNeighbour: calculateNeighbours(i, j, stringMap), isUncovered: false}
			}
		}
		squareMap[i] = squareRow
	}
	return Board{squareMap: squareMap}
}

func calculateNeighbours(i int, j int, stringMap [][]string) int {
	countOfMinesNeighbour := 0
	for ii := -1; ii <= 1; ii++ {
		for jj := -1; jj <= 1; jj++ {
			if isMine(i, ii, j, jj, stringMap) {
				countOfMinesNeighbour++
			}
		}
	}
	return countOfMinesNeighbour
}

func isMine(i int, ii int, j int, jj int, stringMap [][]string) bool {
	return !(i+ii < 0 || j+jj < 0 || i+ii > len(stringMap)-1 || j+jj > len(stringMap[0])-1 || (ii == 0 && jj == 0) || stringMap[i+ii][j+jj] != "X")
}
