package main

import "fmt"

func GenDisplaceFn(acc, vel, disp float64) func(float64) float64 {
	return func(time float64) float64 {
		return 0.5*acc*time*time + vel*time + disp
	}
}

func GetUserParams() (float64, float64, float64) {
	var acc, vel, disp float64
	fmt.Println("Enter acceleration:")
	fmt.Scan(&acc)
	fmt.Println("Enter initial velocity:")
	fmt.Scan(&vel)
	fmt.Println("Enter initial displacement:")
	fmt.Scan(&disp)
	return acc, vel, disp
}

func GetUserTime() float64 {
	var time float64
	fmt.Println("Enter time:")
	fmt.Scan(&time)
	return time
}

func main() {
	acc, vel, disp := GetUserParams()
	fn := GenDisplaceFn(acc, vel, disp)

	time := GetUserTime()
	fmt.Println("Displacement after", time, "seconds:", fn(time))
}
