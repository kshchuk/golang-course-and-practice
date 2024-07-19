package main

func main() {
	const n, k int = 2, 6

	var part_1, part_2 float64 = 1 / float64(n), 1 / float64(k)

	common := part_1 + part_2
	time := 1 / common
	print(time)
}
