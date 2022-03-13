// interfaces

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Square struct {
	Side float64
}

func (square *Square) Area() float64 {
	return math.Pow(square.Side, 2)
}

type Circle struct {
	Radius float64
}

func (circle *Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radius, 2)
}

func getTotalArea(shapes []Shape) float64 {
	areaAccumulator := 0.0

	for _, shape := range shapes {
		areaAccumulator += shape.Area()
	}

	return math.Round(areaAccumulator*1000) / 1000
}

func main() {
	square := Square{4}
	fmt.Printf("%+v\n", square)

	circle := Circle{2.5}
	fmt.Printf("%+v\n", circle)

	shapes := []Shape{&square, &circle}
	totalArea := getTotalArea(shapes)

	fmt.Println(totalArea)
}
