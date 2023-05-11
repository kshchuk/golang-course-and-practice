package main

import (
	"fmt"
	"strings"
)

func main() {
	// Заміна С++ на Go
	s1 := "Я люблю програмування на мові C++"
	s2 := strings.Replace(s1, "C++", "Go", -1)
	fmt.Println(s2)

	// Всі літери великі
	s3 := strings.ToUpper(s1)
	fmt.Println(s3)

	// Виведення фрази 10 разів
	for i := 0; i < 10; i++ {
		fmt.Println(s1)
	}
}
