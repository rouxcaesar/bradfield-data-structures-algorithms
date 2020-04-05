package bst

import (
	"errors"
)

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	Root *Node
}

func (n *Node) Find(i int) (int, bool) {
	if n == nil {
		return i, false
	}

	switch {
	case i == n.value:
		return n.value, true
	case i < n.value:
		return n.left.Find(i)
	default:
		return n.right.Find(i)
	}
}

func (n *Node) Insert(i int) error {
	//fmt.Printf("Node - value i: %d\n", i)
	//fmt.Printf("Node - n.value: %d\n", n.value)
	if n == nil {
		return errors.New("cannot insert value into a nil tree")
	}

	switch {
	case i == n.value:
		return errors.New("value already exists in tree")
	case i < n.value:
		if n.left == nil {
			n.left = &Node{value: i}
			return nil
		}
		return n.left.Insert(i)
	case i > n.value:
		if n.right == nil {
			n.right = &Node{value: i}
			return nil
		}
		return n.right.Insert(i)
	}

	return nil
}

func (t *Tree) Find(i int) (int, bool) {
	if t.Root == nil {
		return i, false
	}

	return t.Root.Find(i)
}

func (t *Tree) Insert(i int) error {
	//fmt.Printf("Tree - value i: %d\n", i)
	if t.Root == nil {
		t.Root = &Node{value: i}
		return nil
	}

	return t.Root.Insert(i)
}
