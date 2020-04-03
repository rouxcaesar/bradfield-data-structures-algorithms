package hs_test

import (
	"bradfield/ds/2-sorting/heap-sort/hs"
	"testing"
)

func benchmarkHeapSort(i int, s string, b *testing.B) {
	slice := hs.GenerateSlice(i, s)
	for n := 0; n < b.N; n++ {
		hs.HeapSort(slice)
	}
}

//func BenchmarkHeapSortRandom10(b *testing.B)    { benchmarkHeapSort(10, "random", b) }
//func BenchmarkHeapSortRandom100(b *testing.B)   { benchmarkHeapSort(100, "random", b) }
//func BenchmarkHeapSortRandom1000(b *testing.B)  { benchmarkHeapSort(1000, "random", b) }
//func BenchmarkHeapSortRandom10000(b *testing.B) { benchmarkHeapSort(10000, "random", b) }

func BenchmarkHeapSortRandom100000(b *testing.B) { benchmarkHeapSort(100000, "random", b) }
func BenchmarkHeapSortEqual100000(b *testing.B)  { benchmarkHeapSort(100000, "equal", b) }
func BenchmarkHeapSortSorted100000(b *testing.B) { benchmarkHeapSort(100000, "sorted", b) }
