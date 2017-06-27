package main

import (
	"fmt"

	. "golang_Algorithm/Algorithm"
)

func main() {
	p := fmt.Printf

	s := []int{4, 2, 3, 1}
	BubbleSort(s)
	p("bubble: %v \n", s)
	SelectSort(s)
	p("select...%v \n", s)

	bt := NewBinaryTree(s)
	preData := bt.PreOrder()
	p("PreOrder...%v \n", preData)
	p("InOrder...%v \n", bt.InOrder())
	p("PostOrder...%v \n", bt.PostOrder())
}
