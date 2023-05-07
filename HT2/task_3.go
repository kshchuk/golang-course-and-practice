package main

import "fmt"

func main() {
	const num int = 135
	reversed := num/100 + num/10%10*10 + num%10*100

	fmt.Println(reversed)
}
