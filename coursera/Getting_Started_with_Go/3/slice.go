package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	const initialSliceSize = 3
	var input string
	var err error
	slice := make([]int, initialSliceSize)

	for i := 0; ; i++ {
		fmt.Println("Please, enter an integer:")
		fmt.Scan(&input)
		if input == "X" {
			break
		}
		slice[i], err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please, enter an integer:")
			i--
			continue
		}
		if i == len(slice)-1 {
			slice = append(slice, make([]int, initialSliceSize)...)
		}

		sort.Ints(slice[:i+1])
		fmt.Println(slice[:i+1])
	}
}
