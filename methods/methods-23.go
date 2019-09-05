/*
	Exercise: rot13Reader
	A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

	For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

	Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

	The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
*/

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {
	switch {
	case x >= 65 && x <= 77, x >= 97 && x <= 109:
		x += 13
	case x >= 78 && x <= 90, x >= 110 && x <= 122:
		x -= 13
	}
	return x
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, error := r13.r.Read(b)
	for i := 0; i <= n; i++ {
		b[i] = rot13(b[i])
	}
	return n, error
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	//s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	// It should output "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
