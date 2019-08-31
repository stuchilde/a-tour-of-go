/*
	Exercise: Readers
	Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (reader MyReader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		bytes[i] = 65
	}
	return len(bytes), nil
}
func main() {
	reader.Validate(MyReader{})
}
