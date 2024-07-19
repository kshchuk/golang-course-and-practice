package main

func main() {
	var arr = []int{4, 3, 1, 5, 3, 14, 3, 4, 7, 10, 11, 12, 4, 3, 3, 56, 3, 12, 122, 3, 2, 3}
	var differentValues int = 0

	var sameValueFound bool
	for i, v := range arr {
		sameValueFound = false
		for j := 0; j < i; j++ {
			if v == arr[j] {
				sameValueFound = true
				break
			}
		}
		if !sameValueFound {
			differentValues++
		}
	}
	println(differentValues)
}
