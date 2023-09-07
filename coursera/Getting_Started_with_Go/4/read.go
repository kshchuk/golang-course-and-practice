package main

import (
	"fmt"
	"os"
)

const maxLen = 20

type Name struct {
	fname string
	lname string
}

func main() {
	var filename string
	fmt.Println("Enter the name of the file:")
	fmt.Scanln(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var names []Name
	var name Name
	for {
		_, err := fmt.Fscanln(file, &name.fname, &name.lname)
		if err != nil {
			break
		}
		if len(name.fname) > maxLen {
			name.fname = name.fname[:maxLen]
		}
		if len(name.lname) > maxLen {
			name.lname = name.lname[:maxLen]
		}

		names = append(names, name)
	}

	for _, name := range names {
		fmt.Println(name.fname, name.lname)
	}
}
