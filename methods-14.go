/*
	Nil interface values
	A nil interface value holds neither value nor concrete type.

	Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
*/

package main

import "fmt"

func main() {
	var i interface{}
	description(i)

	i = 42
	description(42)

	i = "Hello,world"
	description(i)
}

func description(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
