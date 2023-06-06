package main

import "fmt"

type Square struct {
	x      int
	y      int
	length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	return &Square{x: x, y: y, length: length}, nil
}

func (c *Square) Move(dx int, dy int) {
	c.x = c.x + dx
	c.y = c.y + dy
}

func (c *Square) Area() int {
	return c.length ^ 2
}

func main() {

	s, _ := NewSquare(3, 4, 5)
	fmt.Println(s)
	s.Move(10, 10)
	fmt.Println(s)
	fmt.Println(s.Area())

}
