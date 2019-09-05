/*
	Interface values
	Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

	(value, type)
	An interface value holds a value of a specific underlying concrete type.

	Calling a method on an interface value executes the method of the same name on its underlying type.
*/

package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}
type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func description(i I) {
	fmt.Printf("value: %v, type:%T\n", i, i)
}
func main() {
	var i I
	i = T{"Hello,world"}
	description(i)
	i.M()

	i = F(math.Pi)
	description(i)
	i.M()

}
