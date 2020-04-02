package main

import (
	"bradfield/ds/2-sorting/insertion-sort/is"
	"fmt"
)

func main() {
	//slice := qs.GenerateSlice(10, "random")
	slice := []int{4, -31, 0, 99, 83, 1}
	fmt.Printf("---------Unsorted-----------\n")
	fmt.Println(slice)
	fmt.Println()
	sortedSlice := is.InsertionSort(slice)
	fmt.Println()
	fmt.Printf("---------Sorted-----------\n")
	fmt.Println(sortedSlice)
}
