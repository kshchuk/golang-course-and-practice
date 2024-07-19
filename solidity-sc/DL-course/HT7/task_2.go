package main

import "fmt"

func RingS(R1, R2 float32) float32 {
	const pi = 3.14
	return pi * (R2*R2 - R1*R1)
}

func main() {
	var R1, R2 float32 = 5, 10
	S := RingS(R1, R2)
	fmt.Println("RingS(5, 10): ", S)

	fmt.Println("Enter the first Radius: ")
	fmt.Scanln(&R1)
	fmt.Println("Enter the second Radius: ")
	fmt.Scanln(&R2)

	S = RingS(R1, R2)
	fmt.Println("RingS: ", S)
}
