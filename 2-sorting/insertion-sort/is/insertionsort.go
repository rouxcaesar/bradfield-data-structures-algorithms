package is

func InsertionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		// Choose the element in the collection to be sorted.
		element := a[i]
		j := i

		// Loop over previously traversed elements and sort
		// based on chosen element.
		for ; j > 0 && element < a[j-1]; j-- {

			// Insert value into new position in the collection.
			a[j] = a[j-1]
		}

		// Reposition the element we chose to be sorted.
		a[j] = element
	}

	return a
}
