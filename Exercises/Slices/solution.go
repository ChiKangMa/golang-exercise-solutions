package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	//Create the slice of length dy
	a := make([][]uint8, dy)

	//Make each elemnt is a slice of dx 8-bit unsigned integers
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}

	//Choose a draw function for drawing picture
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			a[i][j] = draw(i, j)
		}
	}

	return a
}

func draw(x int, y int) uint8 {
	return uint8(x * y)
}

func draw2(x int, y int) uint8 {
	return uint8(x+y) / 2
}

func draw3(x int, y int) uint8 {
	return uint8(x ^ y)
}

func main() {
	pic.Show(Pic)
}
