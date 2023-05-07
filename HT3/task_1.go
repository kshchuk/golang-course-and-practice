package main

func main() {
	// 1) 200 і 60
	var height, weight int = 200, 60

	if normal_weight := height - 100; weight < normal_weight {
		println("треба поправитися")
	} else if weight > normal_weight {
		println("треба схуднути")
	} else {
		println("норма")
	}
}
