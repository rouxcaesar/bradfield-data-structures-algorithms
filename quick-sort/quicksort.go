package main

import (
	"math/rand"
)

func main() {
	slice := generateSlice(100000)
	//fmt.Println("\n--- Unsorted --- \n\n", slice)
	quicksort(slice)
	//fmt.Println("\n--- Sorted --- \n\n", slice, "\n")
}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		//	slice[i] = rand.Intn(999) - rand.Intn(999)
		slice[i] = 155
	}
	return slice
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	//fmt.Printf("slice a: %v\n", a)

	left, right := 0, len(a)-1
	//fmt.Printf("left: %v\n", left)
	//fmt.Printf("a[left]: %v\n", a[left])
	//fmt.Printf("right: %v\n", right)
	//fmt.Printf("a[right]: %v\n", a[right])

	pivot := rand.Int() % len(a)
	//fmt.Printf("pivot: %v\n", pivot)

	a[pivot], a[right] = a[right], a[pivot]
	//fmt.Printf("a[pivot] after swap: %v\n", a[pivot])
	//fmt.Printf("a[right] after swap: %v\n", a[right])
	//fmt.Printf("slice a after swapping pivot, right: %v\n", a)

	for i, _ := range a {
		//fmt.Printf("i: %v\n", i)
		//fmt.Printf("a[i]: %v\n", a[i])
		//fmt.Printf("a[i] < a[right]: %v\n", a[i] < a[right])
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			//fmt.Printf("a[left] after swap: %v\n", a[left])
			//fmt.Printf("a[i] after swap: %v\n", a[i])
			//fmt.Printf("slice a after swapping i, left: %v\n", a)
			left++
			//fmt.Printf("left after increment: %v\n", left)
		}
	}

	a[left], a[right] = a[right], a[left]
	//fmt.Printf("a[left] after swap: %v\n", a[left])
	//fmt.Printf("a[right] after swap: %v\n", a[right])
	//fmt.Printf("slice a after swapping left, right: %v\n", a)
	//fmt.Println("------------------------------------------")

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
