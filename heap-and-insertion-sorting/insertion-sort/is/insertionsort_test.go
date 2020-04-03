package is_test

import (
	"bradfield/ds/2-sorting/insertion-sort/is"
	"testing"
)

func benchmarkInsertionSort(i int, s string, b *testing.B) {
	slice := is.GenerateSlice(i, s)
	for n := 0; n < b.N; n++ {
		is.InsertionSort(slice)
	}
}

//func BenchmarkInsertionSortRandom10(b *testing.B)     { benchmarkInsertionSort(100000, "random", b) }
//func BenchmarkInsertionSortRandom100(b *testing.B)    { benchmarkInsertionSort(100000, "random", b) }
//func BenchmarkInsertionSortRandom1000(b *testing.B)   { benchmarkInsertionSort(100000, "random", b) }
//func BenchmarkInsertionSortRandom10000(b *testing.B)  { benchmarkInsertionSort(100000, "random", b) }

func BenchmarkInsertionSortRandom100000(b *testing.B) { benchmarkInsertionSort(100000, "random", b) }
func BenchmarkInsertionSortEqual100000(b *testing.B)  { benchmarkInsertionSort(100000, "equal", b) }
func BenchmarkInsertionSortSorted100000(b *testing.B) { benchmarkInsertionSort(100000, "sorted", b) }
