package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Swap(slice []int, index int) {
	temp := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = temp
}

func BubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func main() {
	fmt.Println("Enter a sequence of integers separated by space (up to 10):")
	slice := make([]int, 0, 10)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	numbers := strings.Fields(scanner.Text())
	for _, number := range numbers {
		integer, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		slice = append(slice, integer)
	}

	BubbleSort(slice)

	fmt.Println("Sorted slice:")
	for _, value := range slice {
		fmt.Print(value, " ")
	}
	fmt.Println()
}
