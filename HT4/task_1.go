package main

import "fmt"

func main() {
	const N int = 10

	var bacteria int = 1

	for i := 0; i < N; i++ {
		bacteria *= 2
	}

	fmt.Println(bacteria)
}
