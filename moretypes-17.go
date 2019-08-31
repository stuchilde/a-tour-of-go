/*
	Range continued
	You can skip the index or value by assigning to _.

	for i, _ := range pow
	for _, value := range pow
	If you only want the index, you can omit the second variable.

	for i := range pow
 */
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i:= range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
