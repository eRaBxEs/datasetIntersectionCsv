package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func readCSVKeys(filename string) ([]string, map[string]int) {
	keys := []string{}
	counts := map[string]int{}
	knownHeader := "udprn"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, _ := reader.Read() // Read headers
	udprnIndex := -1
	udprnIndex = findColumnIndex(headers, knownHeader) // Find the index of the "udprn" column
	if udprnIndex == -1 {
		fmt.Printf("Known column header '%v' is not found in filepath:%v:", knownHeader, filename)
		os.Exit(1)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading CSV:", err)
			os.Exit(1)
		}

		key := strings.TrimSpace(record[udprnIndex]) // Extract the key from the "udprn" column
		keys = append(keys, key)
		counts[key]++
	}

	return keys, counts

}

func countDistinctKeys(keys []string) int {
	distinct := map[string]bool{}
	for _, key := range keys {
		distinct[key] = true // set only one instance to true and with a map you only get it once even if it occurs many times in slice
	}
	return len(distinct)
}

func countDistinctOverlap(keys1, keys2 []string) int {
	overlap := map[string]bool{}
	for _, key := range keys1 {
		if contains(keys2, key) {
			overlap[key] = true // sets overlapping once no matter the number of occurrence
		}
	}
	return len(overlap)
}

func calculateOverlapProduct(keys1 []string, counts1 map[string]int, keys2 []string, counts2 map[string]int) int {
	sum := 0

	// Find the common keys
	commonKeys := make(map[string]struct{})
	for _, key := range keys1 {
		if _, ok := counts2[key]; ok {
			commonKeys[key] = struct{}{}
		}
	}

	// Calculate the overlap product
	for key := range commonKeys {
		count1 := counts1[key]
		count2 := counts2[key]
		product := count1 * count2
		sum += product
	}

	return sum

}

func contains(slice []string, key string) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

func findColumnIndex(headers []string, column string) int {
	for i, header := range headers {
		if strings.EqualFold(header, column) {
			return i
		}
	}
	fmt.Println("Column", column, "not found in headers.")
	return -1 // Unreachable, but required for compilation
}
