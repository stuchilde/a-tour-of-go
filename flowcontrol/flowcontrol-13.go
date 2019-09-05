/*
	Stacking defers
	Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

	To learn more about defer statements read this blog post.
*/
package main

import "fmt"

func main() {
	fmt.Println("count")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) //last-in-first-out
	}

	fmt.Println("done")
}
