package qs_test

import (
	"bradfield/ds/quick-sort/qs"
	"testing"
)

func benchmarkQuicksort(i int, s string, b *testing.B) {
	slice := qs.GenerateSlice(i, s)
	for n := 0; n < b.N; n++ {
		qs.Quicksort(slice)
	}
}

//func BenchmarkQuicksortRandom10(b *testing.B)     { benchmarkQuicksort(100000, "random", b) }
//func BenchmarkQuicksortRandom100(b *testing.B)    { benchmarkQuicksort(100000, "random", b) }
//func BenchmarkQuicksortRandom1000(b *testing.B)   { benchmarkQuicksort(100000, "random", b) }
//func BenchmarkQuicksortRandom10000(b *testing.B)  { benchmarkQuicksort(100000, "random", b) }

func BenchmarkQuicksortRandom100000(b *testing.B) { benchmarkQuicksort(100000, "random", b) }
func BenchmarkQuicksortEqual100000(b *testing.B)  { benchmarkQuicksort(100000, "equal", b) }
func BenchmarkQuicksortSorted100000(b *testing.B) { benchmarkQuicksort(100000, "sorted", b) }
