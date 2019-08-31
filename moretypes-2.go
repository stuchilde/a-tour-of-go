/*
	Structs
	A struct is a collection of fields.
 */
package main

import "fmt"

type Vertex2 struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex2{2, 5})
}