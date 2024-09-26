package main

import (
	"reflect"
	"testing"
)

// Test function for twoSum
func TestTwoSum(t *testing.T) {
	// Test case 1: Example input with a valid pair
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	expected1 := []int{0, 1}

	result1 := twoSum(nums1, target1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Expected %v, but got %v", expected1, result1)
	}

	// Test case 2: Another valid pair
	nums2 := []int{3, 2, 4}
	target2 := 6
	expected2 := []int{1, 2}

	result2 := twoSum(nums2, target2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Expected %v, but got %v", expected2, result2)
	}

	// Test case 3: Pair of the same element
	nums3 := []int{3, 3}
	target3 := 6
	expected3 := []int{0, 1}

	result3 := twoSum(nums3, target3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Expected %v, but got %v", expected3, result3)
	}

	// Test case 4: No solution case
	nums4 := []int{1, 2, 3, 4}
	target4 := 8
	expected4 := []int(nil) // No pair adds to target

	result4 := twoSum(nums4, target4)
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Expected %v, but got %v", expected4, result4)
	}
}
