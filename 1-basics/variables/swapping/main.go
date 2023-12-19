package main

import "fmt"

func main() {
	var (
		speed = 100
		prevSpeed = 50
	)

	fmt.Println(speed, prevSpeed)
	speed, prevSpeed = prevSpeed, speed
	fmt.Println(speed, prevSpeed)


}