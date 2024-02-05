package main

import (
	"reflect"
	"testing"
)

func TestReadCSVKeys(t *testing.T) {
	tests := []struct {
		name           string
		file           string
		expected       []string
		expectedCounts map[string]int
	}{
		{
			name:           "Simple CSV",
			file:           "testdata/simple.csv",
			expected:       []string{"A", "B", "C", "D"},
			expectedCounts: map[string]int{"A": 1, "B": 1, "C": 1, "D": 1},
		},
		{
			name:           "CSV with duplicates",
			file:           "testdata/duplicates.csv",
			expected:       []string{"A", "B", "B", "C", "D"},
			expectedCounts: map[string]int{"A": 1, "B": 2, "C": 1, "D": 1},
		},
		{
			name:           "CSV with whitespace",
			file:           "testdata/whitespace.csv",
			expected:       []string{"A", "B", "C", "D"},
			expectedCounts: map[string]int{"A": 1, "B": 1, "C": 1, "D": 1},
		},
		{
			name:           "CSV with empty space",
			file:           "testdata/empty_space.csv",
			expected:       []string{},
			expectedCounts: map[string]int{},
		},

		{
			name:           "Given Dataset 1 Example CSV",
			file:           "testdata/dataset1.csv",
			expected:       []string{"A", "B", "C", "D", "D", "E", "F", "F"},
			expectedCounts: map[string]int{"A": 1, "B": 1, "C": 1, "D": 2, "E": 1, "F": 2},
		},
		{
			name:           "Given Dataset 2 Example CSV",
			file:           "testdata/dataset2.csv",
			expected:       []string{"A", "C", "C", "D", "F", "F", "F", "X", "Y"},
			expectedCounts: map[string]int{"A": 1, "C": 2, "D": 1, "F": 3, "X": 1, "Y": 1},
		},
		{
			name:           "More Columns CSV",
			file:           "testdata/more_columns.csv",
			expected:       []string{"A", "C", "D", "E"},
			expectedCounts: map[string]int{"A": 1, "C": 1, "D": 1, "E": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys, counts := readCSVKeys(tt.file)
			if !reflect.DeepEqual(keys, tt.expected) {
				t.Errorf("Expected keys: %v, got: %v", tt.expected, keys)
			}
			if !reflect.DeepEqual(counts, tt.expectedCounts) {
				t.Errorf("Expected counts: %v, got: %v", tt.expectedCounts, counts)
			}
		})
	}
}

func TestCountDistinctKeys(t *testing.T) {
	tests := []struct {
		name     string
		keys     []string
		expected int
	}{
		{
			name:     "Empty list",
			keys:     []string{},
			expected: 0,
		},
		{
			name:     "All unique keys",
			keys:     []string{"A", "B", "C"},
			expected: 3,
		},
		{
			name:     "Duplicate keys",
			keys:     []string{"A", "A", "B", "C", "C"},
			expected: 3,
		},

		{
			name:     "Dataset 1 example Duplicate keys",
			keys:     []string{"A", "B", "C", "D", "D", "E", "F", "F"},
			expected: 6,
		},
		{
			name:     "Dataset 2 example Duplicate keys",
			keys:     []string{"A", "C", "C", "D", "F", "F", "F", "X", "Y"},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distinct := countDistinctKeys(tt.keys)
			if distinct != tt.expected {
				t.Errorf("Expected %d distinct keys, got %d", tt.expected, distinct)
			}
		})
	}
}

func TestCountDistinctOverlap(t *testing.T) {
	tests := []struct {
		name     string
		keys1    []string
		keys2    []string
		expected int
	}{
		{
			name:     "No overlap",
			keys1:    []string{"A", "B", "C"},
			keys2:    []string{"D", "E", "F"},
			expected: 0,
		},
		{
			name:     "Partial overlap",
			keys1:    []string{"A", "B", "C", "D"},
			keys2:    []string{"C", "D", "E", "F"},
			expected: 2,
		},
		{
			name:     "Full overlap",
			keys1:    []string{"A", "B", "C"},
			keys2:    []string{"A", "B", "C"},
			expected: 3,
		},
		{
			name:     "Dataset 1 and 2 Count Duistinct keys",
			keys1:    []string{"A", "B", "C", "D", "D", "E", "F", "F"},
			keys2:    []string{"A", "C", "C", "D", "F", "F", "F", "X", "Y"},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			overlap := countDistinctOverlap(tt.keys1, tt.keys2)
			if overlap != tt.expected {
				t.Errorf("Expected %d distinct overlapping keys, got %d", tt.expected, overlap)
			}
		})
	}
}

func TestCalculateOverlapProduct(t *testing.T) {
	tests := []struct {
		name               string
		keys1              []string
		counts1            map[string]int
		keys2              []string
		counts2            map[string]int
		expectedSumProduct int
	}{
		{
			name:               "No overlap",
			keys1:              []string{"A", "B", "C"},
			counts1:            map[string]int{"A": 2, "B": 1, "C": 3},
			keys2:              []string{"D", "E", "F"},
			counts2:            map[string]int{"D": 1, "E": 4, "F": 2},
			expectedSumProduct: 0, // SumProduct of empty overlap is 0
		},
		{
			name:               "Partial overlap with different counts",
			keys1:              []string{"A", "B", "B", "C"},
			counts1:            map[string]int{"A": 1, "B": 2, "C": 1},
			keys2:              []string{"A", "B", "C", "C", "D"},
			counts2:            map[string]int{"A": 2, "B": 1, "C": 2, "D": 1},
			expectedSumProduct: 6, // (1 * 2) + (2 * 1) + (2 * 1)
		},
		{
			name:               "Full overlap with equal counts",
			keys1:              []string{"A", "B", "C"},
			counts1:            map[string]int{"A": 1, "B": 1, "C": 1},
			keys2:              []string{"A", "B", "C"},
			counts2:            map[string]int{"A": 1, "B": 1, "C": 1},
			expectedSumProduct: 3, //(1 * 1) + (1 * 1) + (1 * 1)
		},
		{
			name:               "Full overlap with different counts",
			keys1:              []string{"A", "B", "B", "C"},
			counts1:            map[string]int{"A": 1, "B": 2, "C": 3},
			keys2:              []string{"A", "A", "B", "C", "C", "C"},
			counts2:            map[string]int{"A": 2, "B": 1, "C": 3},
			expectedSumProduct: 13, // (1 * 2) + (2 * 1) + (3 * 3)
		},
		{
			name:               "Overlap using given Example",
			keys1:              []string{"A", "B", "C", "D", "D", "E", "F", "F"},
			counts1:            map[string]int{"A": 1, "B": 1, "C": 1, "D": 2, "E": 1, "F": 2},
			keys2:              []string{"A", "C", "C", "D", "F", "F", "F", "X", "Y"},
			counts2:            map[string]int{"A": 1, "C": 2, "D": 1, "F": 3, "X": 1, "Y": 1},
			expectedSumProduct: 11, // (1 * 1) + (1 * 1) + (2 * 1) + (2 * 3)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product := calculateOverlapProduct(tt.keys1, tt.counts1, tt.keys2, tt.counts2)
			if product != tt.expectedSumProduct {
				t.Errorf("Expected overlap product %d, got %d", tt.expectedSumProduct, product)
			}
		})
	}
}
