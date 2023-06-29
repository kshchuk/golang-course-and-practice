package main

import (
	"fmt"
	"math"
)

// y=x^3+x^4
func MyFunction(X float64) float64 {
	Y := math.Pow(X, 3) + math.Pow(X, 4)
	return Y
}

func main() {
	var X float64 = 2
	Y := MyFunction(X)
	fmt.Println("(y=x^3+x^4) (x = 2): ", Y)

	fmt.Println("Enter X: ")
	fmt.Scanln(&X)

	Y = MyFunction(X)
	fmt.Println("y=x^3+x^4: ", Y)
}
