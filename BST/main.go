package main

import (
	"bst/bst"
	"fmt"
)

func main() {
	s := &bst.BSTNode{}
	s.Insert(15)
	s.Insert(16)
	s.Insert(25)
	s.Insert(14)
	s.Insert(82)
	s.Insert(90)
	s.Insert(1)
	s.Insert(2)
	s.TraverseInOrder()
	s = s.Delete(2)
	s = s.Delete(16)
	s = s.Delete(90)
	fmt.Println("AFter DELETION")
	s.TraverseInOrder()
	s.Insert(90)
	s.Insert(1)
	s.Insert(2)
	fmt.Println(s.FindSuccessor(1))
	//fmt.Println(bst.IsTreeBST(&s))
	/*
		s.Insert(15)
		fmt.Println("HEIGHT", s.FindHeight())
		s.Insert(16)
		s.Insert(25)
		fmt.Println("HEIGHT", s.FindHeight())
		s.Insert(14)
		s.Insert(82)
		fmt.Println("HEIGHT", s.FindHeight())
		s.Insert(90)
		s.Insert(1)
		s.Traverse()
		fmt.Println("HEIGHT", s.FindHeight())

		fmt.Println(s.Search(14))
		fmt.Println(s.Search(25))
		fmt.Println(s.Search(30))
		fmt.Println(s.FindMaxIter())
		fmt.Println(s.FindMaxRecur())
		fmt.Println(s.FindMinIter())
		fmt.Println(s.FindMinRecur())*/
}
