package main

import (
	"fmt"
	"math"
)

func main() {
	const a, b float64 = 115, 346

	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	fmt.Println(c) // гіпотенуза

	area := a * b / 2
	fmt.Println(area) // площа
}
