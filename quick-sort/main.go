package main

import "bradfield/ds-alg/quick-sort/qs"

func main() {
	slice := qs.GenerateSlice(100000, "random")
	qs.quicksort(slice)
}
