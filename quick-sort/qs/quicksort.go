package qs

import (
	"fmt"
	"math/rand"
	_ "time"
)

func GenerateSlice(size int, s string) []int {
	slice := make([]int, size, size)
	//rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		switch s {
		case "random":
			slice[i] = rand.Intn(20) - rand.Intn(20)
			//slice[i] = rand.Int()
		case "equal":
			slice[i] = 155
		case "sorted":
			slice[i] = i
		default:
			slice[i] = rand.Intn(20) - rand.Intn(20)
			//slice[i] = rand.Int()
		}
	}

	return slice
}

func Quicksort2(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right, pivot := 0, len(a)-2, len(a)-1

	fmt.Printf("Before loop - left: %d, right: %d, pivot: %d\n", left, right, pivot)
	//a[pivot], a[right] = a[right], a[pivot]

	fmt.Printf("slice before looping: %v\n", a)

	for {
		fmt.Printf("In loop - a[left]: %d, a[right]: %d, a[pivot]: %d\n", a[left], a[right], a[pivot])
		if left >= right {
			a[left], a[pivot] = a[pivot], a[left]
			fmt.Printf("slice after swapping a[left] and a[pivot]: %v\n", a)
			break
		}

		if a[left] > a[pivot] && a[right] < a[pivot] {
			a[left], a[right] = a[right], a[left]
			fmt.Printf("slice after swapping a[left] and a[right]: %v\n", a)
			fmt.Printf("In loop - left: %d, right: %d, pivot: %d\n", left, right, pivot)
			continue
		}

		if a[left] < a[pivot] {
			left++
		}

		if a[right] > a[pivot] {
			right--
		}
		fmt.Printf("In loop - left: %d, right: %d, pivot: %d\n", left, right, pivot)
	}

	fmt.Printf("After loop - left: %d, right: %d, pivot: %d\n", left, right, pivot)
	fmt.Printf("slice after looping: %v\n", a)
	//for i, _ := range a {
	//	if a[i] < a[right] {
	//		a[left], a[i] = a[i], a[left]
	//		left++
	//	}
	//}

	//a[left], a[right] = a[right], a[left]

	Quicksort2(a[:left])
	Quicksort2(a[left+1:])

	return a
}

func Quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]
	fmt.Printf("Before loop - left: %d, right: %d, pivot: %d\n", left, right, pivot)
	fmt.Printf("slice before looping: %v\n", a)

	for i, _ := range a {
		fmt.Printf("slice in loop before check: %v\n", a)
		fmt.Printf("In loop - a[i]: %d, a[left]: %d, a[right]: %d, a[pivot]: %d\n", a[i], a[left], a[right], a[pivot])
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
			fmt.Printf("slice after swapping a[i] and a[left]: %v\n", a)
		}
	}

	a[left], a[right] = a[right], a[left]
	fmt.Printf("slice after swapping a[left] and a[right]: %v\n", a)

	Quicksort(a[:left])
	Quicksort(a[left+1:])

	return a
}
