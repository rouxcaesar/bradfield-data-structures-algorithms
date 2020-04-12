package rbt

// Please implement a red-black tree.
//
// This starter code contains a working binary search tree in which all nodes are
// simply colored RED.  Although this will work reasonably for random inputs,
// you'll need to implement the red-black tree invariants / balancing operations
// before it can handle more challenging inputs (e.g. already sorted).
//
// Good luck!

import "fmt"

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
	if node == nil {
		return newTreeNode(value, red, nil, nil)
	}
	if value <= node.value {
		node.L = insert(node.L, value)
	} else {
		node.R = insert(node.R, value)
	}

	// balance tree here
	// balance(node)

	return node
}

func balance(node *treeNode) {
	if node.color == red {
		return
	}

	//fmt.Println("------------------------------")
	//fmt.Println("Node before balancing")
	//node.render(2)
	//fmt.Println()

	if node.L.color == red {
		if node.L.L.color == red {
			temp := node
			left := node.L
			childOfLeft := node.L.L

			// Swap nodes.
			node = left
			node.L = childOfLeft
			node.R = temp

			//fmt.Println("Node after balancing")
			//node.render(2)
			//fmt.Println()
		} else if node.L.R.color == red {
			temp := node
			left := node.L
			childOfLeft := node.L.R

			// Swap nodes.
			node = childOfLeft
			node.L = left
			node.R = temp

			//fmt.Println("Node after balancing")
			//node.render(2)
			//fmt.Println()
		}
		// else if node.L.R.color == red {

		//}
	}

	fmt.Println("------------------------------")

	return
}
