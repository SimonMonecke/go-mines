package convert

import "fmt"

type Square struct {
	isMine                bool
	isUncovered           bool
	countOfMinesNeighbour int
}

func (s *Square) String() string {
	if s.isMine {
		return "X"
	}
	return fmt.Sprintf("%d", s.countOfMinesNeighbour)
}

func (s *Square) IsMine() bool {
	return s.isMine
}

func (s *Square) IsUncovered() bool {
	return s.isUncovered
}

func (s *Square) CountOfMinesNeighbour() int {
	return s.countOfMinesNeighbour
}

func (s *Square) Uncover() bool {
	s.isUncovered = true
	return !s.isMine
}
