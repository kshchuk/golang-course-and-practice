package main

func main() {
	var arr = []int{4, 3, 1, 5, 3, 14, 3, 4, 7, 10, 11, 12, 4, 3, 3, 56, 3, 12, 122, 3, 2, 3}
	var max int = arr[0]

	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}

	println(max)
}
