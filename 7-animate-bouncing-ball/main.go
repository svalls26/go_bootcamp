package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/inancgumus/screen"
)

const (
	rows = 3
	cols = 3
) 

func main() {
	//Create the board
	board := make([][]string, rows)

	for i := range board {
		board[i] = make([]string, cols)
	}


	for {
		screen.Clear()
		screen.MoveTopLeft()
		//Get a random true
		r := rand.Intn(rows) 
		c := rand.Intn(cols)
		
		board[r][c] = "🥎"
		for i := range board {
			row := board[i]
			for _, v := range row {
				fmt.Printf("%-15s", v)
			}
			fmt.Println()
		}

		board[r][c] = ""
		time.Sleep(100 * time.Millisecond)
	}
	
}