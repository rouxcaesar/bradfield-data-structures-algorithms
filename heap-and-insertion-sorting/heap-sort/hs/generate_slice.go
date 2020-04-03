package hs //_test

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
