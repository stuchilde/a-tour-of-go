/*
	Select
	The select statement lets a goroutine wait on multiple communication operations.

	A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
*/

package main

import "fmt"

func fibonacci(quit, ch chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
	fibonacci(quit, ch)
}
