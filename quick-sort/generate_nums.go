package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	totalNums = 100000

	random = true

	equal = false

	sorted = false
)

func main() {
	// Seed with different numbers with each program execution.
	rand.Seed(time.Now().UnixNano())

	// Create file.
	f, err := os.Create("random_nums.txt")
	if err != nil {
		log.Fatalf("failed to create file")
	}

	// Loop through iterations and write  nums
	// to file.
	if random {
		for i := 0; i < totalNums; i++ {
			f.WriteString(strconv.Itoa(rand.Int()) + "\n")
		}
	} else if equal {
		chosenInt := strconv.Itoa(rand.Int())
		for i := 0; i < totalNums; i++ {
			f.WriteString(chosenInt + "\n")
		}
	} else if sorted {
		for i := 0; i < totalNums; i++ {
			f.WriteString(strconv.Itoa(i) + "\n")
		}
	}
}
