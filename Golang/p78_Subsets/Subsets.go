package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Subsets function using Backtracking(DFS)
func subsets(nums []int) [][]int {
	res, subset := [][]int{}, []int{}
	// Pass agurments into function
	dfs(0, &res, subset, nums)
	return res
}

func dfs(i int, res *[][]int, subset []int, nums []int) {
	if i == len(nums) {
		//Making a tmporary variable to copy subset
		temp := make([]int, len(subset))
		copy(temp, subset)
		//Appending subset into res
		*res = append(*res, temp)
		return
	}
	//Appending nums[i] into subset
	subset = append(subset, nums[i])
	dfs(i+1, res, subset, nums)
	// Backtracking
	subset = subset[:len(subset)-1]
	dfs(i+1, res, subset, nums)
}

// Testing function is created by chatGPT
func main() {
	// Example inputs to test manually
	inputs := [][]int{
		{1, 2},
		{1, 2, 3},
		{},
		{1, 2, 3, 4},
	}

	// Print the subsets for each input
	for _, nums := range inputs {
		fmt.Printf("Subsets of %v: %v\n", nums, subsets(nums))
	}
}

// Testing function
func TestSubsets(t *testing.T) {
	// Test case 1
	test1 := []int{1, 2}
	expected1 := [][]int{{}, {1}, {2}, {1, 2}}

	// Test case 2
	test2 := []int{1, 2, 3}
	expected2 := [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}

	// Test case 3
	test3 := []int{}
	expected3 := [][]int{{}}

	// Test case 4
	test4 := []int{1, 2, 3, 4}
	expected4 := [][]int{
		{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3},
		{4}, {1, 4}, {2, 4}, {1, 2, 4}, {3, 4}, {1, 3, 4}, {2, 3, 4}, {1, 2, 3, 4},
	}

	// Run the test cases
	t.Run("Test with input []int{1, 2}", func(t *testing.T) {
		result := subsets(test1)
		if !reflect.DeepEqual(result, expected1) {
			t.Errorf("Expected %v, but got %v", expected1, result)
		}
	})

	t.Run("Test with input []int{1, 2, 3}", func(t *testing.T) {
		result := subsets(test2)
		if !reflect.DeepEqual(result, expected2) {
			t.Errorf("Expected %v, but got %v", expected2, result)
		}
	})

	t.Run("Test with empty input []int{}", func(t *testing.T) {
		result := subsets(test3)
		if !reflect.DeepEqual(result, expected3) {
			t.Errorf("Expected %v, but got %v", expected3, result)
		}
	})

	t.Run("Test with input []int{1, 2, 3, 4}", func(t *testing.T) {
		result := subsets(test4)
		if !reflect.DeepEqual(result, expected4) {
			t.Errorf("Expected %v, but got %v", expected4, result)
		}
	})
}
