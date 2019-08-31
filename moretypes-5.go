/*
	Struct Literals
	A struct literal denotes a newly allocated struct value by listing the values of its fields.

	You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

	The special prefix & returns a pointer to the struct value.
 */
package main

import "fmt"

type Vertex5 struct {
	X, Y int
}
func main() {
	v1 := Vertex5{1, 2}
	v2 := Vertex5{X: 0}
	v3 := Vertex5{}
	p := &Vertex5{1, 5}
	fmt.Println(v1, v2, v3, p)
}