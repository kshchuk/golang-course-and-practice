package main

import "fmt"

func main() {
	var arr = []int{48, 96, 86, 68, 57, 82, 63, 70}

	slice := arr[:4]

	fmt.Println(slice)
}
