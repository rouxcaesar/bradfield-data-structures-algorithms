package main

import (
	"bradfield/ds/quick-sort/qs"
	"fmt"
)

func main() {
	slice := qs.GenerateSlice(10, "random")
	fmt.Printf("---------Unsorted-----------\n")
	fmt.Println(slice)
	sortedSlice := qs.Quicksort(slice)
	//sortedSlice := qs.Quicksort2(slice)
	fmt.Printf("---------Sorted-----------\n")
	fmt.Println(sortedSlice)
}
