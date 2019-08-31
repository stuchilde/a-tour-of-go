/*
	Slices of slices
	Slices can contain any type, including other slices.
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string {
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[0][2] = "X"
	board[1][2] = "X"
	board[1][0] = "O"
	board[2][2] = "O"

	for i:=0; i<len(board); i++ {
		fmt.Println(strings.Join(board[i], " "))
	}
}
