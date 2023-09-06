package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please, enter a string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	str := scanner.Text()
	str = strings.ToLower(str)

	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
