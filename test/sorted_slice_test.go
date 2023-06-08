package test

import (
	training "github.com/pnowosie/go-codility"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedSliceSortsACopy(t *testing.T) {
	A := []int{9, 0, 8, 7, 6, 4, 5, 2, 3, 1}
	B := training.NewSortedSlice(A)

	// B can be iterated just like A
	for i, v := range *B {
		assert.Equal(t, i, v)
	}

	// A has not been modified
	assert.True(t, A[0] > A[8])

	// B has a Len method
	assert.Equal(t, 10, B.Len())
}

func TestMapSortedSlice(t *testing.T) {
	str := []string{"4444", "22", "", "333", "1"}
	s := training.MapSortedSlice(str, func(s string) int { return len(s) })
	for i, v := range *s {
		assert.Equal(t, i, v)
	}
}

func TestUseBinarySearchWithSortedSlice(t *testing.T) {
	s := training.NewSortedSlice([]int{6, 7, 8, 9, 0, 1, 2, 3, 4})

	find := training.BinarySearch(*s, 2)
	assert.Equal(t, find, 2)

	find = training.BinarySearch(*s, 9)
	assert.Equal(t, find, 8)

	find = training.BinarySearch(*s, 5)
	assert.Equal(t, find, -1)
}
