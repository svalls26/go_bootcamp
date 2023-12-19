package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l := len(msg)

	exclamation := strings.Repeat("!", l)
	s := exclamation + msg + exclamation
	s = strings.ToUpper(s)

	fmt.Println(s)
}