/*
	From https://tour.golang.org/concurrency/8
*/

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkTree(t, ch)
	close(ch)
}

func WalkTree(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		WalkTree(t.Left, ch)
	}
	if t != nil {
		ch <- t.Value
	}
	if t.Right != nil {
		WalkTree(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for val1 := range ch1 {
		if val1 != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for val := range ch {
		fmt.Println(val)
	}

	sameCall := Same(tree.New(1), tree.New(1))
	fmt.Println(sameCall)

	differentCall := Same(tree.New(1), tree.New(2))
	fmt.Println(differentCall)

}
