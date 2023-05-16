package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func main() {
	point1 := Point{1, 2}
	point2 := Point{1, 4}
	point3 := Point{7, 4}

	width := math.Abs(point2.Y - point1.Y)
	height := math.Abs(point3.X - point1.X)

	area := width * height

	fmt.Println("Площа прямокутника:", area)
}
