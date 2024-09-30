package main

import "fmt"

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
	nums1 := []int{1, 2, 3}

	nums2 := []int{2, 3, 4}
	fmt.Println("Subsets of [1, 2, 3]:")
	fmt.Println(subsets(nums1))

	fmt.Println("Subsets of [2,3,4]:")
	fmt.Println(subsets(nums2))
}
