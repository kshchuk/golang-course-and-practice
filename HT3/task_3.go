package main

func main() {
	// 1) D=15, M=1;
	var D, M int = 20, 9

	switch M {
	case 1:
		if D < 20 {
			println("Козеріг")
		} else {
			println("Водолій")
		}
		break
	case 2:
		if D < 19 {
			println("Водолій")
		} else {
			println("Риби")
		}
		break
	case 3:
		if D < 21 {
			println("Риби")
		} else {
			println("Овен")
		}
		break
	case 4:
		if D < 20 {
			println("Овен")
		} else {
			println("Телець")
		}
		break
	case 5:
		if D < 21 {
			println("Телець")
		} else {
			println("Близнюки")
		}
		break
	case 6:
		if D < 22 {
			println("Близнюки")
		} else {
			println("Рак")
		}
		break
	case 7:
		if D < 23 {
			println("Рак")
		} else {
			println("Лев")
		}
		break
	case 8:
		if D < 23 {
			println("Лев")
		} else {
			println("Діва")
		}
		break
	case 9:
		if D < 23 {
			println("Діва")
		} else {
			println("Терези")
		}
		break
	case 10:
		if D < 23 {
			println("Терези")
		} else {
			println("Скорпіон")
		}
		break
	case 11:
		if D < 23 {
			println("Скорпіон")
		} else {
			println("Стрілець")
		}
		break
	case 12:
		if D < 22 {
			println("Стрілець")
		} else {
			println("Козеріг")
		}
		break
	default:
		return
	}
}
