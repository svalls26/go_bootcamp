package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {
	max, err := strconv.Atoi(os.Args[1]) 
	if err != nil {
		fmt.Printf("%d is not permitted: %v\n", max, err)
		return
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)		
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}

		fmt.Println()
	}
}