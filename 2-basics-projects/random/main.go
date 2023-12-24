package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	rand.NewSource(time.Now().UnixNano())
	
	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}