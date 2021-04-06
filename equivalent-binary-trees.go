package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
    if !ok1 || !ok2 {
			break
		}
		if v1 != v2 || ok1 != ok2 {
			return false
		}
	}
	return true
}

func main() {
	test1, test2 := Same(tree.New(1), tree.New(1)), Same(tree.New(1), tree.New(2))
	fmt.Println("Same(tree.New(1), tree.New(1)):", test1)
	fmt.Println("Same(tree.New(1), tree.New(2)):", test2)
	if test1 && !test2 {
		fmt.Println("Success!")
	} else {
		fmt.Println("Something's not right...")
	}
}
