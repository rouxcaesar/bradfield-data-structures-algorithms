package rbt

// Please implement a red-black tree.
//
// This starter code contains a working binary search tree in which all nodes are
// simply colored RED.  Although this will work reasonably for random inputs,
// you'll need to implement the red-black tree invariants / balancing operations
// before it can handle more challenging inputs (e.g. already sorted).
//
// Good luck!

import (
	"fmt"
	"time"
)

const (
	red   = "RED"
	black = "BLACK"
)

type treeNode struct {
	value int
	color string
	L     *treeNode
	R     *treeNode
}

func newTreeNode(value int, color string, L *treeNode, R *treeNode) *treeNode {
	return &treeNode{
		value: value,
		color: color,
		L:     L,
		R:     R,
	}
}

// Helper function for printing a tree to the console.  You will need to
// turn your head sideways to look at the output :
func (node *treeNode) render(indent int) {
	if node.R != nil {
		node.R.render(indent + 1)
	}
	for i := 0; i < indent; i++ {
		fmt.Print("\t")
	}
	fmt.Printf("%d(%s)\n", node.value, node.color)
	if node.L != nil {
		node.L.render(indent + 1)
	}
}

func find(node *treeNode, value int) bool {
	if node == nil {
		return false
	} else if value < node.value {
		return find(node.L, value)
	} else if value > node.value {
		return find(node.R, value)
	} else {
		return true
	}
}

func insert(node *treeNode, value int) *treeNode {
	time.Sleep(5 * time.Second)
	fmt.Println("--------------------------------")
	fmt.Println("Inside insert func")
	fmt.Printf("value: %v\n", value)
	if node != nil {
		node.render(0)
	}

	if node == nil {
		return newTreeNode(value, red, nil, nil)
	}
	if value <= node.value {
		node.L = insert(node.L, value)
	} else {
		node.R = insert(node.R, value)
	}

	//time.Sleep(10 * time.Second)
	// balance tree here
	fmt.Println("Node before balancing")
	fmt.Printf("node is now: %v\n", node)
	node = balance(node)

	fmt.Println("Node after balancing")
	fmt.Printf("node is now: %v\n", node)
	node.render(0)

	return node
}

func balance(node *treeNode) *treeNode {

	fmt.Println("--------------------------------")
	fmt.Println("Inside balance func")
	fmt.Printf("node.value: %v\n", node.value)
	if node.L != nil {
		fmt.Printf("node.L.value: %v\n", node.L.value)
	}
	if node.R != nil {
		fmt.Printf("node.R.value: %v\n", node.R.value)
	}
	fmt.Println("Node before balancing")
	node.render(0)

	if node.L != nil && node.L.color == red {
		fmt.Printf("node.L.value: %d  node.L.color: %s\n", node.L.value, node.L.color)
		if node.L.L != nil && node.L.L.color == red {
			fmt.Printf("node.L.L detected for swap: %v\n", node.L.L.value)
			temp := node
			left := node.L
			childOfLeft := node.L.L

			fmt.Printf("temp: %d\nleft: %d\nchildOfLeft: %d\n", temp.value, left.value, childOfLeft.value)

			//// Swap nodes.
			//node = left
			//node.L = childOfLeft
			//node.R = temp

			//// Recolor nodes.
			//node.color = red
			//node.L.color = black
			//node.R.color = black

			//fmt.Println("Node after balancing")
			//node.render(2)
			//fmt.Println()
			return node
		} else if node.L.R != nil && node.L.R.color == red {
			fmt.Printf("node.L.R detected for swap: %v\n", node.L.R.value)
			temp := node
			left := node.L
			childOfLeft := node.L.R

			fmt.Printf("temp: %d\nleft: %d\nchildOfLeft: %d\n", temp.value, left.value, childOfLeft.value)

			//// Swap nodes.
			//node = childOfLeft
			//node.L = left
			//node.R = temp

			//// Recolor nodes.
			//node.color = red
			//node.L.color = black
			//node.R.color = black

			//fmt.Println("Node after balancing")
			//node.render(2)
			//fmt.Println()
			return node
		}
	} else if node.R != nil && node.R.color == red {
		fmt.Printf("node.R.value: %d  node.R.color: %s\n", node.R.value, node.R.color)
		if node.R.L != nil && node.R.L.color == red {
			fmt.Printf("node.R.L detected for swap: %v\n", node.R.L)
			parent := node
			right := node.R
			childOfRight := node.R.L
			var b *treeNode
			var c *treeNode

			if childOfRight.L != nil {
				b = childOfRight.L
				fmt.Printf("childOfRight.L: %v\n", childOfRight.L)
				fmt.Printf("b: %v\n", b)
			} else {
				b = nil
				fmt.Printf("childOfRight.L is NIL\n")
			}

			if childOfRight.R != nil {
				c = childOfRight.R
				fmt.Printf("childOfRight.R: %v\n", childOfRight.R)
				fmt.Printf("c: %v\n", c)
			} else {
				c = nil
				fmt.Printf("childOfRight.R is NIL\n")
			}

			fmt.Printf("parent: %d\nright: %d\nchildOfRight: %d\n", parent.value, right.value, childOfRight.value)
			fmt.Printf("parent: %v\nright: %v\nchildOfRight: %v\n", parent, right, childOfRight)
			fmt.Printf("SWAPPING\n")

			// Swap nodes.
			node = childOfRight
			fmt.Printf("node.value is now: %v\n", node.value)
			fmt.Printf("node is now: %v\n", node)

			node.L = parent
			fmt.Printf("node.L.value is now: %v\n", node.L.value)
			fmt.Printf("node.L is now: %v\n", node.L)

			node.R = right
			fmt.Printf("node.R.value is now: %v\n", node.R.value)
			fmt.Printf("node.R is now: %v\n", node.R)

			node.L.R = b
			//fmt.Printf("node.L.R.value is now: %v\n", node.L.R.value)
			fmt.Printf("node.L.R is now: %v\n", node.L.R)

			node.R.L = c
			//fmt.Printf("node.R.L.value is now: %v\n", node.R.L.value)
			fmt.Printf("node.R.L is now: %v\n", node.R.L)

			//if parent.L != nil {
			//	node.L.L = parent.L
			//}
			//if childOfRight.L != nil {
			//	fmt.Printf("childOfRight.L: %v\n", childOfRight.L)
			//	node.L.R = childOfRight.L
			//} else {
			//	node.L.R = nil
			//}

			//if childOfRight.R != nil {
			//	fmt.Printf("childOfRight.R: %v\n", childOfRight.R)
			//	node.R.L = childOfRight.R
			//} else {
			//	node.R.L = nil
			//}
			//if right.R != nil {
			//	node.R.R = right.R
			//}
			fmt.Printf("node: %d\nnode.L: %d\nnode.R: %d\n", node.value, node.L.value, node.R.value)

			fmt.Printf("RECOLORING\n")
			// Recolor nodes.
			node.color = red
			node.L.color = black
			node.R.color = black

			fmt.Println("Node after balancing")

			//	if node.L.L != nil {
			//		fmt.Printf("node.L.L: %d\n", node.L.L.value)
			//	}
			//	if node.L.R != nil {
			//		fmt.Printf("node.L.R: %d\n", node.L.R.value)
			//	}
			//	if node.R.L != nil {
			//		fmt.Printf("node.R.L: %d\n", node.R.L.value)
			//	}
			//	if node.R.R != nil {
			//		fmt.Printf("node.R.R: %d\n", node.R.R.value)
			//	}

			node.render(0)
			//fmt.Println()
			return node
		} else if node.R.R != nil && node.R.R.color == red {
			fmt.Printf("node.R.R detected for swap: %v\n", node.R.R)
			temp := node
			right := node.R
			childOfRight := node.R.R

			fmt.Printf("temp: %d\nright:%d\n childOfRight: %d\n", temp.value, right.value, childOfRight.value)

			//// Swap nodes.
			//node = right
			//node.L = temp
			//node.R = childOfRight

			//// Recolor nodes.
			//node.color = red
			//node.L.color = black
			//node.R.color = black

			//fmt.Println("Node after balancing")
			//node.render(2)
			//fmt.Println()
			return node
		}
	}
	return node
}
