package main

import (
	"fmt"
	"log"
)

type Square struct {
	X      int
	Y      int
	Length int
}

func NewSquare(x int, y int, Length int) (*Square, error) {
	if Length <= 0 {
		return nil, fmt.Errorf("Length must be greater than 0")
	}
	s := Square{
		X:      x,
		Y:      y,
		Length: Length,
	}
	return &s, nil
}

func (s *Square) Move(dx int, dy int) {
	s.X += dx
	s.Y += dy
}

func (s *Square) Area() int {
	area := s.Length * s.Length
	return area
}

func main() {
	// fmt.Println("Welcome to my world!")
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("ERROR: can't create square.")
	}

	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}
