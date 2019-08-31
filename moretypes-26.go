/**
Exercise: Fibonacci closure
Let's have some fun with functions.

Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
*/
package main

import "fmt"

func fibonacci() func() int {
	left, right := 0, 1
	return func() int {
		left, right = right, left+right
		return right
	}
}

func main() {
	f := fibonacci()
	for i := 10; i > 0; i-- {
		fmt.Println(f())
	}
}
