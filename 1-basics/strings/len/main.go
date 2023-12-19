package main

import "fmt"

func main() {
	name := "carl"

	//len returns the number of bytes in the string, NOT the number of chars

	fmt.Println(
		len(name))

	name = "īnanç"
	fmt.Println(
		len(name))
}