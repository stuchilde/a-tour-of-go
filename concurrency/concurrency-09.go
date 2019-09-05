/*
	sync.Mutex
	We've seen how channels are great for communication among goroutines.

	But what if we don't need communication? What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?

	This concept is called mutual exclusion, and the conventional name for the data structure that provides it is mutex.

	Go's standard library provides mutual exclusion with sync.Mutex and its two methods:

	Lock
	Unlock
	We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock as shown on the Inc method.

	We can also use defer to ensure the mutex will be unlocked as in the Value method.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (s *SafeCounter) Inc(key string) {
	s.mux.Lock()
	s.v[key]++
	s.mux.Unlock()
}

func (s *SafeCounter) Value(key string) int {
	s.mux.Lock()

	defer s.mux.Unlock()
	return s.v[key]
}

func main() {
	s := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go s.Inc("hello world")
	}
	time.Sleep(time.Second)
	fmt.Println(s.Value("hello world"))
}
