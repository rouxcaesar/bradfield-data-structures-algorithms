package hs

import "fmt"

func HeapSort(a []int) []int {
	a = buildMaxHeap(a)
	fmt.Printf("heap: %v\n", a)

	// Loop over collection to sort elements.
	//for i := 0; i < len(a); i++ {

	//heap = balanceHeap(heap, numElements)
	//fmt.Printf("balanced heap: %v\n", heap)
	//}

	return a
}

func buildMaxHeap(a []int) []int {

	// Loop over collection to swap elements until no swaps occur.
	for {
		noSwaps := true
		for i := 0; i < len(a); i++ {
			fmt.Printf("i: %v\n", i)

			if ((2*i)+1) < len(a) && a[i] < a[(2*i)+1] {
				a[i], a[(2*i)+1] = a[(2*i)+1], a[i]
				fmt.Printf("slice after swap w/ first node: %v\n", a)
				noSwaps = false
			}

			if ((2*i)+2) < len(a) && a[i] < a[(2*i)+2] {
				a[i], a[(2*i)+2] = a[(2*i)+2], a[i]
				fmt.Printf("slice after swap w/ second node: %v\n", a)
				noSwaps = false
			}
		}
		if noSwaps {
			break
		}
	}

	return a
}

func balanceHeap(a []int, num int) []int {
	return a
}
