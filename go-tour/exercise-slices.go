/*
	Go Tour Exercise: Slices https://tour.golang.org/moretypes/18
*/
package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dx)
	for x := range image {
		image[x] = make([]uint8, dy)
		
		for y := range image {
			image[x][y] = uint8(x+dy^(x+y)/3)
		}
	}
	return image
}


func main() {
	pic.Show(Pic)
}