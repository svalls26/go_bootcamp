// This is barely used
package main

import "fmt"


func main() {
	var i int

loop: //Not recommend the use, it is better just to use a for loop
	if i < 3 {
		fmt.Println("looping")
		i++
		goto loop
	}
	fmt.Println("done")
}

