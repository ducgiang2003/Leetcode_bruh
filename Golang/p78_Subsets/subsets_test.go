package main

import (
	"reflect"
	"testing"
)

// Test function for the subsets function
func TestSubsets(t *testing.T) {
	// Test case 1: Example input [1, 2, 3]
	nums1 := []int{1, 2, 3}
	expected1 := [][]int{
		{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3},
	}

	result1 := subsets(nums1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test Case 1 Failed: expected %v, but got %v", expected1, result1)
	}

	// Test case 2: Single element [0]
	nums2 := []int{0}
	expected2 := [][]int{
		{}, {0},
	}

	result2 := subsets(nums2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test Case 2 Failed: expected %v, but got %v", expected2, result2)
	}

	// Test case 3: Empty input []

	// Test case 4: Input with duplicates [1, 2, 2]

}
