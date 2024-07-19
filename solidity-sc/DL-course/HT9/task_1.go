package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Створення директорії Test
	err := os.Mkdir("Test", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Запис у файл test.txt
	content := []byte("Я люблю програмування\nМова програмування Golang")
	err = ioutil.WriteFile("Test/test.txt", content, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File created!")
}
