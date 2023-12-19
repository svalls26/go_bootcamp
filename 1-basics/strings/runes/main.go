package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "īnanç"

	fmt.Println(
		utf8.RuneCountInString(name))
}