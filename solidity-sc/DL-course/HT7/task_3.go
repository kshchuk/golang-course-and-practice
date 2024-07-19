package main

import "fmt"

func CountDigits(number uint) uint {
	var count uint = 0
	for number != 0 {
		count++
		number /= 10
	}
	return count
}

func main() {
	var N1, N2, N3 uint = 321, 1243, 12543
	C := CountDigits(N1)
	fmt.Println("CountDigits(321): ", C)
	C = CountDigits(N2)
	fmt.Println("CountDigits(1243): ", C)
	C = CountDigits(N3)
	fmt.Println("CountDigits(12543): ", C)

	fmt.Println("Enter the Number: ")
	fmt.Scanln(&N1)

	C = CountDigits(N1)
	fmt.Println("RingS: ", C)
}
