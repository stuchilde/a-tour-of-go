/*
	Exercise: Slices
	Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

	The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

	(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

	(Use uint8(intValue) to convert between types.)
 */
package main
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	x := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		x[i] = make([]uint8, dx)
	}
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			switch  {
			case j % 4 == 0:
				x[i][j] = 240
			case j % 4 == 1:
				x[i][j] = 50
			case j % 4 == 2:
				x[i][j] = 150
			case j % 4 == 3:
				x[i][j] = 100
			}
		}
	}
	return x
}
func main() {
	pic.Show(Pic)
}
