package main

import "fmt"

func SumRange(A, B int) int {
	var sum int = 0
	for i := A; i <= B; i++ {
		sum += i
	}
	return sum
}

func main() {
	var A, B int = 10, 20
	C := SumRange(A, B)
	fmt.Println("SumRange(10, 20): ", C)

	fmt.Println("Enter the first variable: ")
	fmt.Scanln(&A)
	fmt.Println("Enter the second variable: ")
	fmt.Scanln(&B)

	C = SumRange(A, B)
	fmt.Println("SumRange: ", C)
}
