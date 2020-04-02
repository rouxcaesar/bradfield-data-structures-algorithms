package is

import _ "fmt"

func InsertionSort(a []int) []int {

	for i := 0; i < len(a); i++ {
		// Choose the element in the collection to be sorted.
		element := a[i]
		j := i
		//fmt.Printf("chosen element: %v\n", element)

		// Loop over previously traversed elements and sort
		// based on chosen element.
		for ; j > 0 && element < a[j-1]; j-- {
			//fmt.Printf("element %v < %v\n", element, a[j-1])

			// Insert value into new position in the collection.
			a[j] = a[j-1]
		}

		// Reposition the element we chose to be sorted.
		a[j] = element
		//fmt.Printf("slice is now: %v\n", a)
	}

	return a
}
