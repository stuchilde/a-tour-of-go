/*
	Struct Fields
	Struct fields are accessed using a dot.
 */
package main

import "fmt"

type Vertex3 struct {
	X int
	Y int
}

func main() {
	v := Vertex3{1, 4}
	v.X = 4
	fmt.Println(v.X)
}