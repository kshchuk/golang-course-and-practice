package main

import (
	"fmt"
)

func main() {
	var f float64
	fmt.Println("Please, enter a floating point number:")
	fmt.Scan(&f)
	fmt.Println("Truncated number:", int64(f))
}
