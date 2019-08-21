/*
	Zero values
	Variables declared without an explicit initial value are given their zero value.

	The zero value is:

	0 for numeric types,
	false for the boolean type, and
	"" (the empty string) for strings.
*/

package main

import "fmt"

func main() {
	var b bool
	var f float64
	var i int
	var s string
	fmt.Printf("%v %v %v %q", b, f, i, s)
}
