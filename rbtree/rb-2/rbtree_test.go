package rbtree

import (
	"fmt"
	"math/rand"
	"testing"
)

func flatten(node *treeNode) []int {
	if node == nil {
		return nil
	}
	var result []int
	result = append(result, flatten(node.L)...)
	result = append(result, node.value)
	result = append(result, flatten(node.R)...)
	return result
}

func validBstOrder(root *treeNode) bool {
	elements := flatten(root)
	for i := 0; i < len(elements)-1; i++ {
		if elements[i] > elements[i+1] {
			return false
		}
	}
	return true
}

func traversePaths(node *treeNode, count int, counts map[int]struct{}) {
	if node == nil {
		counts[count] = struct{}{}
		return
	}
	if node.color == black {
		count++
	}
	traversePaths(node.L, count, counts)
	traversePaths(node.R, count, counts)
}

// Perform DFS, ensure that every path from the root to an empty node has the
// same number of black nodes along the way.
func validRbtreePaths(root *treeNode) bool {
	counts := make(map[int]struct{})
	traversePaths(root, 0, counts)
	return len(counts) == 1
}

// Perform DFS, ensure that no red node has a red parent
func validRbtreeColors(node *treeNode, parentColor string) bool {
	if node == nil {
		return true
	}
	if node.color == red && parentColor == red {
		return false
	}
	return validRbtreeColors(node.L, node.color) &&
		validRbtreeColors(node.R, node.color)
}

func sorted(n int) []int {
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, i)
	}
	return result
}

func reversed(n int) []int {
	var result []int
	for i := n; i > 0; i-- {
		result = append(result, i)
	}
	return result
}

func randomized(n, upper int) []int {
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, rand.Intn(upper))
	}
	return result
}
func TestInsert(t *testing.T) {
	N := 10
	for _, testCase := range []struct {
		name   string
		values []int
	}{
		//{"sorted", sorted(N)},
		//{"reversed", reversed(N)},
		{"randomized 1", randomized(N, N)},
		//{"randomized 2", randomized(N, 10*N)},
	} {
		var root *treeNode
		for _, value := range testCase.values {
			root = insert(root, value)
		}
		fmt.Printf("\nTest case %q:\n", testCase.name)
		root.render(0)
		if !validBstOrder(root) {
			t.Fatalf("Invalid BST order for case %q", testCase.name)
		}

		// TODO: Uncomment these once you believe you've implemented a
		// working red-black tree.
		/*
			if !validRbtreeColors(root, black) {
				t.Fatalf("Found two red nodes in a row for case %q", testCase.name)
			}
			if !validRbtreePaths(root) {
				t.Fatalf("Found two paths with different black node counts for case %q", testCase.name)
			}
		*/
	}
}

func TestFind(t *testing.T) {
	N := 100

	var root *treeNode

	// Insert even integers
	for i := 0; i < N; i += 2 {
		root = insert(root, i)
	}

	for i := 0; i < N; i += 2 {
		if !find(root, i) {
			t.Fatalf("Expected to find %d", i)
		}
	}

	for i := 1; i < N; i += 2 {
		if find(root, i) {
			t.Fatalf("Expected not to find %d", i)
		}
	}
}
