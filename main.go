package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// type Square struct {
// 	X      int
// 	Y      int
// 	Length int
// }

// func NewSquare(x int, y int, Length int) (*Square, error) {
// 	if Length <= 0 {
// 		return nil, fmt.Errorf("Length must be greater than 0")
// 	}
// 	s := Square{
// 		X:      x,
// 		Y:      y,
// 		Length: Length,
// 	}
// 	return &s, nil
// }

// func (s *Square) Move(dx int, dy int) {
// 	s.X += dx
// 	s.Y += dy
// }

// func (s *Square) Area() int {
// 	area := s.Length * s.Length
// 	return area
// }

type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')
	// out := make([]byte, len(p))
	for i, x := range p {
		if x >= 'a' && x <= 'z' {
			x -= diff
		}
		p[i] = x
	}
	return c.wtr.Write(p)
}

func main() {
	// fmt.Println("Welcome to my world!")
	// s, err := NewSquare(1, 1, 10)
	// if err != nil {
	// 	log.Fatalf("ERROR: can't create square.")
	// }

	// s.Move(2, 3)
	// fmt.Printf("%+v\n", s)
	// fmt.Println(s.Area())

	c := &Capper{os.Stdout}

	var month time.Month
	var year int
	var day int
	year, month, day = time.Now().Local().Date()

	fmt.Fprintf(c, "My name is Mohit Maurya. Today's day is: %d, month is: %v, year is: %d \n", day, month, year)
}
