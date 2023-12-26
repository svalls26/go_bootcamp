package main

import "fmt"


func main() {
	var (
		blue = [...]int{6, 9, 3}
		red = [3]int{6, 9, 3}
	)	

	fmt.Printf("Blue bookcase : %v\n", blue)
	fmt.Printf("Red bookcase : %v\n", red)

	fmt.Println("Are they equal?", blue == red)
}