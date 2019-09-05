/*
Exercise: Equivalent Binary Trees
1. Implement the Walk function.

2. Test the Walk function.

The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.

Create a new channel ch and kick off the walker:

	go Walk(tree.New(1), ch)
	Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

	3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

	4. Test the Same function.

	Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

	The documentation for Tree can be found here.
*/

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Walker(t *tree.Tree) <-chan int {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := Walker(t1), Walker(t2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))

}
