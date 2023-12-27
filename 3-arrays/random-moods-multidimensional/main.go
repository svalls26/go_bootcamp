package main

import (
	"fmt"
	"math/rand"
	"os"
)

const (
	usage = "[your name] [positive|negative]"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println(usage)
	}

	name, mood := args[0], args[1]
	moods := [2][3]string{
		{"good", "awsome", "spectacular"},
		{"bad", "terrible", "sad"},
	}

	var mi int
	if mood != "positive"{
		mi = 1
	}

	fmt.Printf("%s feels %s\n", name, moods[mi][rand.Intn(len(moods[0])-1)])


}