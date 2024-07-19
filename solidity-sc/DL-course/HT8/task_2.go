package main

import (
	"fmt"
	"math"
)

// AB^2 = AC^2 + BC^2 - 2 * AC * BC * cos(∠ACB)
func CosTheorem(AC, BC, angleC float64) float64 {
	radians := angleC * (math.Pi / 180)
	return AC*AC + BC*BC - 2*AC*BC*math.Cos(radians)
}

func main() {
	var AC, BC, angleC float64 = 22, 21, 60
	AB_squared := CosTheorem(AC, BC, angleC)
	fmt.Println("AB^2: ", AB_squared)

	fmt.Println("Enter AC: ")
	fmt.Scanln(&AC)
	fmt.Println("Enter BC: ")
	fmt.Scanln(&BC)
	fmt.Println("Enter ∠ACB: ")
	fmt.Scanln(&angleC)

	AB_squared = CosTheorem(AC, BC, angleC)
	fmt.Println("AB^2: ", AB_squared)
}
