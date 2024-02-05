## Documentation

### 1. readCSVKeys(filename string) ([]string, map[string]int)
Purpose: Reads a CSV file and returns unique keys and their counts from a specified column.
Parameters:
filename: Path to the CSV file.
Returns:
keys: List of unique keys.
counts: Map of key counts.

### 2. countDistinctKeys(keys []string) int
Purpose: Counts the number of distinct keys in a list.
Parameters:
keys: List of keys.
Returns:
Number of unique keys.

### 3. countDistinctOverlap(keys1 []string, keys2 []string) int
Purpose: Counts the number of distinct keys present in both input lists.
Parameters:
keys1: First list of keys.
keys2: Second list of keys.
Returns:
Number of unique overlapping keys.

### 4. calculateOverlapProduct(keys1 []string, counts1 map[string]int, keys2 []string, counts2 map[string]int) int
Purpose: Calculates the sum of product of counts for overlapping keys in two key-count pairs.
Parameters:
keys1, counts1: First key-count pair.
keys2, counts2: Second key-count pair.
Returns:
Sum of Product of counts for overlapping keys.

