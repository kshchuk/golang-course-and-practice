package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func TriangleArea(p1, p2, p3 Point) float64 {
	area := math.Abs((p1.X*(p2.Y-p3.Y) + p2.X*(p3.Y-p1.Y) + p3.X*(p1.Y-p2.Y)) / 2)
	return area
}

func main() {
	point1 := Point{2, -4}
	point2 := Point{-5, -6}
	point3 := Point{1, 3}

	area := TriangleArea(point1, point2, point3)

	fmt.Println("Площа трикутника:", area)
}
