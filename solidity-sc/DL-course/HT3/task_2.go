package main

func main() {
	// 1) 12
	var num int = 134

	if num%2 == 0 {
		if num < 10 {
			println("парне однозначне число")
		} else if num < 100 {
			println("парне двозначне число")
		} else if num < 1000 {
			println("парне тризначне число")
		}
	} else {
		if num < 10 {
			println("непарне однозначне число")
		} else if num < 100 {
			println("непарне двозначне число")
		} else if num < 1000 {
			println("непарне тризначне число")
		}
	}
}
