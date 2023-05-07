package main

import (
	"fmt"
	"math"
)

func main() {
	var arr = []int{48, 96, 86, 68, 57, 82, 63, 70}

	slice := arr[len(arr)-4:]
	for i := range slice {
		if slice[i] < 65 {
			slice[i] = int(math.Pow(float64(slice[i]), 2))
		}
	}

	for i := range slice {
		for j := range slice {
			if slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	fmt.Println(slice)
}
