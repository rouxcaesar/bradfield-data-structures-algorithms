package main

import (
	"bradfield/ds/2-sorting/heap-sort/hs"
	"fmt"
)

func main() {
	//slice := []int{3, 19, 1, 14, 8, 7}
	slice := hs.GenerateSlice(10, "sorted")
	fmt.Printf("---------Unsorted-----------\n")
	fmt.Println(slice)
	fmt.Println()
	sortedSlice := hs.HeapSort(slice)
	fmt.Println()
	fmt.Printf("---------Sorted-----------\n")
	fmt.Println(sortedSlice)
}
