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
	fmt.Println("--------------------------------")
	fmt.Println("Inside insert func")
	fmt.Printf("value: %v\n", value)

	if node == nil {
		return newTreeNode(value, red, nil, nil)
	}
	if value <= node.value {
		node.L = insert(node.L, value)
	} else {
		node.R = insert(node.R, value)
	}

	time.Sleep(10 * time.Second)
	// balance tree here
	balance(node)

	fmt.Println("--------------------------------")
	return node
}

func balance(node *treeNode) {

	fmt.Println("--------------------------------")
	fmt.Println("Inside balance func")
	fmt.Printf("node.value: %v\n", node.value)
	if node.L != nil {
		fmt.Printf("node.L.value: %v\n", node.L.value)
	}
	if node.R != nil {
		fmt.Printf("node.R.value: %v\n", node.R.value)
	}
	//	node.render(0)
	//fmt.Println("------------------------------")
	//fmt.Println("Node before balancing")
	//node.render(2)
	//fmt.Println()

	if node.L != nil && node.L.color == red {
		if node.L.L != nil && node.L.L.color == red {
			fmt.Printf("node.L.L detected for swap: %v\n", node.L.L.value)
			//temp := node
			//left := node.L
			//childOfLeft := node.L.L

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
			return
		} else if node.L.R != nil && node.L.R.color == red {
			fmt.Printf("node.L.R detected for swap: %v\n", node.L.R.value)
			//temp := node
			//left := node.L
			//childOfLeft := node.L.R

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
			return
		} else if node.R != nil && node.R.color == red {
			if node.R.L != nil && node.R.L.color == red {
				fmt.Printf("node.R.L detected for swap: %v\n", node.R.L)
				//temp := node
				//right := node.R
				//childOfRight := node.R.L

				//// Swap nodes.
				//node = childOfRight
				//node.L = temp
				//node.R = right

				//// Recolor nodes.
				//node.color = red
				//node.L.color = black
				//node.R.color = black

				//fmt.Println("Node after balancing")
				//node.render(2)
				//fmt.Println()
				return
			} else if node.R.R != nil && node.R.R.color == red {
				fmt.Printf("node.R.R detected for swap: %v\n", node.R.R)
				//temp := node
				//right := node.R
				//childOfRight := node.R.R

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
				return
			}
		}
	}

	fmt.Println("--------------------------------")
	return
}
