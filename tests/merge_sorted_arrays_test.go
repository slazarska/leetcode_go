package tests

import (
	"leetcode_go/easy"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
	}

	for _, test := range tests {
		copyNums1 := append([]int(nil), test.nums1...)
		easy.MergeSortedArrays(copyNums1, test.m, test.nums2, test.n)

		for i, v := range copyNums1 {
			if v != test.expected[i] {
				t.Errorf("Test failed for nums1 = %v, nums2 = %v. Expected %v but got %v", test.nums1, test.nums2, test.expected, copyNums1)
				break
			}
		}
	}
}
