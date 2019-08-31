/*
	Functions
	A function can take zero or more arguments.

	In this example, add takes two parameters of type int.

	Notice that the type comes after the variable name.

	(For more about why types look the way they do, see the article on Go's declaration syntax.)
*/
package main

import (
	"fmt"
)

func add1(x int, y int) int {
	return x + y
}
func main() {
	fmt.Println(add1(23, 34))
}
