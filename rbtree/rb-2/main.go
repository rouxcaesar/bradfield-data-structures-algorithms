package main

import (
	"bradfield/ds/quick-sort/qs"
	"fmt"
)

func main() {
	values := []int{6, 4, 1, 2, 8, 7}
	fmt.Printf("values: %v\n", values)
	var root *rbtree.treeNode
	for _, value := range values {
		root = rbtree.insert(root, value)
	}
	root.render(0)
}
