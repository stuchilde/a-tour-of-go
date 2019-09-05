/*
	Numeric Constants
	Numeric constants are high-precision values.

	An untyped constant takes the type needed by its context.

	Try printing needInt(Big) too.

	(An int can store at maximum a 64-bit integer, and sometimes less.)
*/
package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}
