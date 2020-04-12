package main

import (
	"bradfield/ds/rbtree/rb-2/rbt"
	"fmt"
)

func main() {
	values := []int{6, 4, 1, 2, 8, 7}
	fmt.Printf("values: %v\n", values)
	var root *rbt.treeNode
	for _, value := range values {
		root = rbt.insert(root, value)
	}
	root.render(0)
}
