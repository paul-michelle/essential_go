// custom type, new instance, pointer

package main

import (
	"fmt"
	"math"
	"os"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) *Point {
	return &Point{x,y}
}
func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

type Square struct {
	Center Point
	Length float64
}

func (anySquare *Square) Move(dx int, dy int) {
	anySquare.Center.Move(dx, dy)
}


func (anySqare *Square) Area() float64 {
	return math.Pow(anySqare.Length, 2)
}

func NewSquare(x int, y int, length float64) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("Invalid length value")
	}
	point := NewPoint(x, y)
	square := &Square{*point, length}
	return square, nil

}

func main() {
	square, err := NewSquare(1, -3, 1.5)
	if err != nil {
		fmt.Printf("Error: failed to create a new square - %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("New square created: %+v\n", square)
	fmt.Printf("New square's area: %0.2f\n", square.Area())
	square.Move(-1, 3)
	fmt.Printf("Square's center moved to point: %+v\n", square.Center)
	fmt.Printf("Square now: %+v\n", square)
}
