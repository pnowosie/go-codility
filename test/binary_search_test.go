package test

import (
	training "github.com/pnowosie/go-codility"
	"testing"
)

func TestBinarySearchExactMatch(t *testing.T) {
	A := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	tests := map[string]struct {
		x     int
		index int
	}{
		"Finding 1st element":    {0, 0},
		"Finding last element":   {9, 9},
		"Finding middle element": {5, 5},
		"Not found negative":     {-1, -1},
		"Not found positive":     {10, -1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			index := training.BinarySearch(A, test.x)
			if index != test.index {
				t.Errorf("Expected %d, got %d", test.index, index)
			}
		})
	}
}

func TestBinarySearchBounds(t *testing.T) {
	A := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}
	tests := map[string]struct {
		x     int
		index int
		left  int
		right int
	}{
		"Finding 1st element":         {0, 0, 0, 8},
		"Finding 1st element, upper":  {0, 0, 0, 4},
		"Finding 1st element, l==r":   {0, 0, 0, 0},
		"Finding 1st element, l>0":    {0, 1, 1, 8},
		"Finding last element":        {9, 8, 0, 8},
		"Finding last element, lower": {9, 8, 5, 8},
		"Finding last element, l==r":  {9, 8, 8, 8},
		"Finding last element, r<len": {9, 8, 1, 7},
		"Missing 5":                   {5, 5, 0, 8},
		"Missing 5, l<5":              {5, 5, 0, 4},
		"Missing 5, l>=5":             {5, 5, 5, 8},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			index := training.BinaryFindLR(A, test.x, test.left, test.right)
			if index != test.index {
				t.Errorf("Expected %d, got %d", test.index, index)
			}
		})
	}
}
