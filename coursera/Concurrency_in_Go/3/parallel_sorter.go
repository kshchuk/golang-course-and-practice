package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func QuickSort(arr []int, ch chan []int) {
	if len(arr) <= 1 {
		ch <- arr
		return
	}

	pivot := arr[0]
	left, right := []int{}, []int{}

	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	chLeft := make(chan []int)
	chRight := make(chan []int)

	go QuickSort(left, chLeft)
	go QuickSort(right, chRight)

	ch <- append(append(<-chLeft, pivot), <-chRight...)
}

func merge(arr1, arr2 []int) []int {
	fmt.Println("Merging:", arr1, arr2)

	arr := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr[k] = arr1[i]
			i++
		} else {
			arr[k] = arr2[j]
			j++
		}
		k++
	}

	for i < len(arr1) {
		arr[k] = arr1[i]
		i++
		k++
	}

	for j < len(arr2) {
		arr[k] = arr2[j]
		j++
		k++
	}

	return arr
}

func ParallelQuickSort(arr []int, parts int) []int {
	ch := make(chan []int)
	size := (len(arr) + 1) / parts

	for i := 0; i < parts; i++ {
		start := i * size
		end := (i + 1) * size

		if end > len(arr) {
			end = len(arr)
		}

		go func() {
			fmt.Println("Sorting:", arr[start:end])
			QuickSort(arr[start:end], ch)
		}()
	}

	sortedArr := []int{}
	for i := 0; i < parts; i++ {
		sortedArr = merge(sortedArr, <-ch)
	}

	return sortedArr
}

func main() {
	fmt.Println("Please enter a sequence of integers separated by space: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	arr := []int{}
	for _, v := range strings.Split(scanner.Text(), " ") {
		n, _ := strconv.Atoi(v)
		arr = append(arr, n)
	}

	fmt.Println("Unsorted Array:", arr)

	fmt.Println("Sorted Array:", ParallelQuickSort(arr, 4))
}
