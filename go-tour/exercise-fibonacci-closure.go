/*
	Go Tour Exercise: Fibonacci closure https://tour.golang.org/moretypes/26
*/
package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func() int {
	first := 0
	second := 1
	sum := 0
	return func() int {
		sum = first + second
		defer func() {
			first, second = second, sum
		}()
		return first
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}