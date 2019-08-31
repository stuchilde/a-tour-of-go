package main

import (
	"fmt"
	"math"
)

/*
	Pointers and functions
	Here we see the Abs and Scale methods rewritten as functions.

	Again, try removing the * from line 16. Can you see why the behavior changes? What else did you need to change for the example to compile?

	(If you're not sure, continue to the next page.)
*/

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
