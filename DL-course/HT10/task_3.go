package main

import "fmt"

func main() {

	for n := 0; n <= 3; n++ {

		if n%2 == 0 {

			continue

		}

		fmt.Println(n)

	}

}
