package bst_test

import (
	"bradfield/ds-alg/rb-trees-1/binary-search-tree/bst"
	"log"
	"testing"
)

func benchmarkBST(i int, s string, b *testing.B) {
	values := bst.GenerateSlice(i, s)

	for n := 0; n < b.N; n++ {
		tree := &bst.Tree{}

		for _, v := range values {
			err := tree.Insert(v)
			if err != nil {
				log.Fatalf("error inserting value %d into tree: %s\n", v, err)
			}
		}
	}
}

//func BenchmarkBSTRandom10(b *testing.B) { benchmarkBST(10, "random", b) }

//func BenchmarkBSTRandom100(b *testing.B)   { benchmarkBST(100, "random", b) }
//func BenchmarkBSTRandom1000(b *testing.B)  { benchmarkBST(1000, "random", b) }
//func BenchmarkBSTRandom10000(b *testing.B) { benchmarkBST(10000, "random", b) }

func BenchmarkBSTRandom100000(b *testing.B) { benchmarkBST(100000, "random", b) }

//func BenchmarkBSTEqual100000(b *testing.B)  { benchmarkBST(100000, "equal", b) }
func BenchmarkBSTSorted100000(b *testing.B) { benchmarkBST(100000, "sorted", b) }
