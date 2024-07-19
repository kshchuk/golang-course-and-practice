package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	var name, address string
	fmt.Println("Please, enter your name:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	name = scanner.Text()

	fmt.Println("Please, enter your address:")

	scanner.Scan()
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	address = scanner.Text()

	m := map[string]string{
		"name":    name,
		"address": address,
	}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(b))
}
