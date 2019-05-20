/*
	Go Tour Exercise: Equivalent Binary Trees https://tour.golang.org/concurrency/7
*/
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

var length = 10

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if(t != nil){
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i:=0;i<length;i++ {
		if(<-ch1 != <-ch2){
			return false
		}
	}
	return true
}

func TestSame(){
	isSame := false
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			isSame = Same(tree.New(i), tree.New(j))
			fmt.Printf("Same(tree.New(%d), tree.New(%d)): %t\n", i, j, isSame)
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i:=0;i<length;i++ {
		fmt.Println(<-ch)
	}
	TestSame()
}
