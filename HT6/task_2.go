package main

import "fmt"

func main() {
	var arr = []int{48, 96, 86, 68, 57, 82, 63, 70}

	slice := arr[:5]

	for i := range slice {
		slice[i] += 10
	}
	fmt.Println(slice)
}
