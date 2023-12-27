package main

import (
	"fmt"
	"math/rand"
	"os"
)

const (
	usage = "usage [your name]"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println(usage)
		return
	}

	name := args[0]

	moods := [...]string{
		"happy 😀", "awsome 🥳", "fantastic 🎉",
		"bad ☹️", "terrible 😣",
	}

	fmt.Printf("%s feels %s\n", name, moods[rand.Intn(len(moods))])

}