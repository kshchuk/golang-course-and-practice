package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name, food, locomotion, noise string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

type Bird struct {
	name, food, locomotion, noise string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

type Snake struct {
	name, food, locomotion, noise string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func ProccessCommand(command string, animal Animal) {
	switch command {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Invalid command")
	}

}

func main() {
	var animal, command, name string
	var animals []Animal

	for {
		fmt.Print("> ")
		fmt.Scan(&command, &name, &animal)

		switch command {
		case "newanimal":
			switch animal {
			case "cow":
				animals = append(animals, Cow{name, "grass", "walk", "moo"})
			case "bird":
				animals = append(animals, Bird{name, "worms", "fly", "peep"})
			case "snake":
				animals = append(animals, Snake{name, "mice", "slither", "hsss"})
			default:
				fmt.Println("Invalid animal type")
			}
			fmt.Println("Created it!")
		case "query":
			for _, a := range animals {
				if an, e := a.(Cow); e {
					if an.name == name {
						ProccessCommand(animal, an)
						break
					}
				} else if an, e := a.(Bird); e {
					if an.name == name {
						ProccessCommand(animal, an)
						break
					}
				} else if an, e := a.(Snake); e {
					if an.name == name {
						ProccessCommand(animal, an)
						break
					}
				}
			}
		default:
			fmt.Println("Invalid command")
		}
	}
}
