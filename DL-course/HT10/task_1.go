package main

import "fmt"

type Employee struct {
	LastName string
	Salary   int
}

func main() {
	employees := []Employee{
		{"Сміт", 10000},
		{"Джонс", 12000},
		{"Доу", 14000},
		{"Смітт", 16000},
		{"Джо", 20000},
	}

	totalSalary := 0
	for _, emp := range employees {
		totalSalary += emp.Salary
	}

	fmt.Println("Сума окладів:", totalSalary)
}
