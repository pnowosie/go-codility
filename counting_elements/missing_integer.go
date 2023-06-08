// Package counting_elements
// Problem [3] MissingInteger
// Copyright 2009–2023 by Codility Limited
package counting_elements

var _description3 = `
Write a function:

func Solution(A []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.
Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

- N is an integer within the range [1..100,000];
- each element of array A is an integer within the range [−1,000,000..1,000,000].
`

// Imports

// Code
func Solution3(A []int) int {
	exist := make([]bool, len(A))
	for _, v := range A {
		if v > 0 && v <= len(exist) {
			exist[v-1] = true
		}
	}
	for i := range exist {
		if !exist[i] {
			return i + 1
		}
	}

	return len(A) + 1
}

// Helpers
