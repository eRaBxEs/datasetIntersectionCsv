package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go actions.go <file1.csv> <file2.csv>")
		os.Exit(1)
	}

	file1, file2 := os.Args[1], os.Args[2]

	// process keys1 and counts1
	keys1, counts1 := readCSVKeys(file1)
	// process keys2 and counts2
	keys2, counts2 := readCSVKeys(file2)

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
