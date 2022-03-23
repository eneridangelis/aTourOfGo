package main

import "golang.org/x/tour/pic"

func showF(dx, dy int) [][]uint8 {
	ss := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		s := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			s[x] = uint8((x + y) / 2)
		}
		ss[y] = s
	}
	return ss
}

func Pic(dx, dy int) [][]uint8 {
	f := showF(dx, dy)
	return f
}

func main() {
	pic.Show(Pic)
}
