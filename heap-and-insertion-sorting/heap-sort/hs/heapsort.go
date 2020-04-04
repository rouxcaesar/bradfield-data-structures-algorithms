package hs

func HeapSort(a []int) []int {
	// Restructure all of the elements in
	// the slice argument into  a heap data structure.
	a = buildMaxHeap(a, len(a))

	// Loop over collection to sort elements.
	for i := 0; i < len(a)-1; i++ {
		// Swap the element at index 0 with last
		// last element in the heap.
		a[0], a[len(a)-1-i] = a[len(a)-1-i], a[0]

		// We rebuild a heap with all elements excluding
		// the sorted elements positioned at the
		// end of the slice.
		//a = buildMaxHeap(a, len(a)-1-i)
		a = rebuildHeap(a, len(a)-1-i)
	}

	return a
}

func rebuildHeap(a []int, num int) []int {
	j := 0

	for i := 0; i < num; i++ {
		if ((2*i)+1) < num && a[j] < a[(2*i)+1] {
			a[i], a[(2*i)+1] = a[(2*i)+1], a[i]
			j++
			i++
			continue
		}

		// The second child node is found at index (2*i + 2).
		// If the parent node at index i is less than the child node,
		// swap the node values.
		if ((2*i)+2) < num && a[j] < a[(2*i)+2] {
			a[j], a[(2*i)+2] = a[(2*i)+2], a[j]
			j++
			i++
			continue
		}
	}

	return a
}

// buildMaxHeap takes a slice and reorganizes a subset of the slice
// elements into a heap data structure.
// This subset is identified as the elements from index 0 up to
// the index num argument (inclusive).
// Any elements at indexes greater than the num argument are
// the already sorted elements that should not be part of
// the heap data structure.
func buildMaxHeap(a []int, num int) []int {
	// Loop over collection to swap elements until no swaps occur.
	for {
		// noSwaps keeps track of whether any node values have been swapped.
		// If there have been no swaps, then we know that the elements have
		// been correctly reorganized into a heap data structure.
		// If there has been a swap, we need to loop over the elements again
		// to check if we've finished constructing the heap data structure.
		noSwaps := true
		for i := 0; i < num; i++ {

			// The first child node is found at index (2*i + 1).
			// If the parent node at index i is less than the child node,
			// swap the node values.
			if ((2*i)+1) < num && a[i] < a[(2*i)+1] {
				a[i], a[(2*i)+1] = a[(2*i)+1], a[i]
				noSwaps = false
			}

			// The second child node is found at index (2*i + 2).
			// If the parent node at index i is less than the child node,
			// swap the node values.
			if ((2*i)+2) < num && a[i] < a[(2*i)+2] {
				a[i], a[(2*i)+2] = a[(2*i)+2], a[i]
				noSwaps = false
			}
		}
		// If we have no swapping of node values, then our
		// heap has been created and we can exit the loop.
		if noSwaps {
			break
		}
	}

	return a
}
