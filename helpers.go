package training

import "sort"

// BinaryFindLR finds the index of the first element in A that is greater or equal to x.
func BinaryFindLR(A []int, x int, left, right int) int {
	if left < 0 || right >= len(A) {
		panic("0 <= left <= right < len(A)")
	}

	for left <= right {
		mid := (left + right) / 2
		if A[mid] == x {
			return mid
		} else if A[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func BinarySearch(A []int, x int) int {
	idx := BinaryFindLR(A, x, 0, len(A)-1)
	if idx < len(A) && A[idx] == x {
		return idx
	}
	return -1
}

type IntSliceSorted []int

func (s *IntSliceSorted) Len() int {
	if s == nil {
		return 0
	}
	return len(*s)
}

func NewSortedSlice(A []int) *IntSliceSorted {
	b := make([]int, len(A))
	copy(b, A)
	s := IntSliceSorted(b)
	sort.Ints(s)
	return &s
}

func MapSortedSlice[T any](A []T, fn func(T) int) *IntSliceSorted {
	b := make([]int, len(A))
	for i, v := range A {
		b[i] = fn(v)
	}
	s := IntSliceSorted(b)
	sort.Ints(s)
	return &s
}

// by-key sorter
func Solution(A []int, K int) []int {
	negate := func(e1, e2 *sT) bool {
		return -(*e1) < -(*e2)
	}

	arr := make([]sT, len(A))
	for i, v := range A {
		arr[i] = sT(v)
	}

	By(negate).Sort(arr)

	for i, v := range arr {
		A[i] = int(v)
	}

	return A
}

type sT int
type Sortable interface {
	Less(j sT) bool
}

// Less type func
type By func(e1, e2 *sT) bool

func (by By) Sort(elts []sT) {
	sorter := &Sorter{
		elts: elts,
		by:   by,
	}
	sort.Sort(sorter)
}

type Sorter struct {
	elts []sT
	by   func(e1, e2 *sT) bool
}

func (s *Sorter) Len() int {
	return len(s.elts)
}

func (s *Sorter) Swap(i, j int) {
	s.elts[i], s.elts[j] = s.elts[j], s.elts[i]
}

func (s *Sorter) Less(i, j int) bool {
	return s.by(&s.elts[i], &s.elts[j])
}
