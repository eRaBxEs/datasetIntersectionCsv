package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go actions.go <file1.csv> <file2.csv>")
		os.Exit(1)
	}

	file1, file2 := os.Args[1], os.Args[2]

	// preparing keys and count ahead
	var keys1, keys2 []string
	counts1 := map[string]int{}
	counts2 := map[string]int{}

	var wg sync.WaitGroup
	wg.Add(2) // Increment the WaitGroup counter for two goroutines

	go func() {
		// process keys1 and counts1
		keys1, counts1 = readCSVKeys(file1)
		wg.Done()
	}()

	go func() {
		// process keys2 and counts2
		keys2, counts2 = readCSVKeys(file2)
		wg.Done()
	}()

	wg.Wait() // Wait for both goroutines to finish

	// distinct keys
	distinctKeys1 := countDistinctKeys(keys1)
	distinctKeys2 := countDistinctKeys(keys2)

	// distinct overlap
	distinctOverlap := countDistinctOverlap(keys1, keys2)

	// sum of overlap product
	overlapProduct := calculateOverlapProduct(keys1, counts1, keys2, counts2)

	// Display results
	fmt.Println("File 1:")
	fmt.Println("  Number of keys:", len(keys1))
	fmt.Println("  Number of distinct keys:", distinctKeys1)

	fmt.Println("File 2:")
	fmt.Println("  Number of keys:", len(keys2))
	fmt.Println("  Number of distinct keys:", distinctKeys2)

	fmt.Println("Distinct Overlap:", distinctOverlap)
	fmt.Println("Overlap Product:", overlapProduct)
}
