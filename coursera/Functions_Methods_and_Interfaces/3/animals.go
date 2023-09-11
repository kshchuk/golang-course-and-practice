package main

import "fmt"

type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) Eat() {
	println(a.food)
}

func (a *Animal) Move() {
	println(a.locomotion)
}

func (a *Animal) Speak() {
	println(a.noise)
}

func NewAnimal(animalType string) *Animal {
	switch animalType {
	case "cow":
		return &Animal{"grass", "walk", "moo"}
	case "bird":
		return &Animal{"worms", "fly", "peep"}
	case "snake":
		return &Animal{"mice", "slither", "hsss"}
	default:
		return nil
	}
}

func (a *Animal) DoAction(action string) {
	switch action {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("Please, enter a valid action and try again.")
	}
}

func main() {
	var animalType, action string
	println("Enter animal name and action separated by space (e.g. cow eat):")
	for {
		fmt.Print("> ")
		_, err := fmt.Scan(&animalType, &action)
		if err != nil {
			fmt.Println("Please, enter a valid value and try again.")
			continue
		}
		animal := NewAnimal(animalType)
		if animal == nil {
			fmt.Println("Please, enter a valid animal name and try again.")
			continue
		}
		animal.DoAction(action)
	}
}
