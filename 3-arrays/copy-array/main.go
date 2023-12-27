package main

import "fmt"


func main() {
	prev := [3]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
	}	

	var books [4]string

	for i, b := range prev {
		books[i] = b + " 2nd Ed."
	}
	books[3] = "Awesomness"

	fmt.Printf("last year:\n%#v\n", prev)
	fmt.Printf("\nthis year:\n%#v\n", books)
}