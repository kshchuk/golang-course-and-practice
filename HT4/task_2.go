package main

import (
	"fmt"
)

func main() {
	const N int = 1200000
	var sum int = 0
	for i := 1; i <= N; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
