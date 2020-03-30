package qs

import (
	"math/rand"
	"time"
)

func GenerateSlice(size int, s string) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		switch s {
		case "random":
			slice[i] = rand.Intn(999) - rand.Intn(999)
		case "equal":
			slice[i] = 155
		case "sorted":
			slice[i] = i
		default:
			slice[i] = rand.Intn(999) - rand.Intn(999)
		}
	}

	return slice
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
