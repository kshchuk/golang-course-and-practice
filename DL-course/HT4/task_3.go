package main

import (
	"fmt"
)

func main() {
	const N int = 1200000
	var n int = 0 // counter

	for i := 1; i <= N; i++ {
		if i%2 == 0 && i%5 == 0 {
			n += 1
		}
	}

	percentage := float64(n) / float64(N) * 100

	fmt.Println("Divisible by 2 and 5: ", percentage, "%")
}
