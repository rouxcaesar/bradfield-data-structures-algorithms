package qs

import (
	"math/rand"
	"time"
	_ "time"
)

func GenerateSlice(size int, s string) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		switch s {
		case "random":
			slice[i] = rand.Int()
		case "equal":
			slice[i] = 155
		case "sorted":
			slice[i] = i
		default:
			slice[i] = rand.Int()
		}
	}

	return slice
}

func Quicksort2(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right, pivot := 0, len(a)-2, len(a)-1

	//a[pivot], a[right] = a[right], a[pivot]

	for {
		if left >= right {
			a[left], a[pivot] = a[pivot], a[left]
			break
		}

		if a[left] > a[pivot] && a[right] < a[pivot] {
			a[left], a[right] = a[right], a[left]
			continue
		}

		if a[left] < a[pivot] {
			left++
		}

		if a[right] > a[pivot] {
			right--
		}
	}

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

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	Quicksort(a[:left])
	Quicksort(a[left+1:])

	return a
}
