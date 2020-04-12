package main

import (
	"bradfield/ds-alg/rb-trees-1/binary-search-tree/bst"
	"fmt"
	"log"
)

func main() {
	values := []int{2, 7, 9, 1, 4, 6}
	fmt.Printf("---------Values-----------\n")
	fmt.Println(values)
	fmt.Printf("---------Binary Search Tree-----------\n")
	tree := &bst.Tree{}
	for _, v := range values {
		err := tree.Insert(v)
		if err != nil {
			log.Fatalf("error inserting value %d into tree: %s\n", v, err)
		}
	}
	fmt.Printf("Tree: %#v\n", tree)
}
