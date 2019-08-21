/*
	Variables with initializers
	A var declaration can include initializers, one per variable.

	If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
*/
package main

import (
	"fmt"
)

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
