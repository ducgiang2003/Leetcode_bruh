package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	subset := []int{}
	//First must sort the arr(nums)
	sort.Ints(nums)
	var backtracing func(nums []int, start int, subset []int, res *[][]int)
	backtracing = func(nums []int, start int, subset []int, res *[][]int) {
		//Check if start > len(nums) then return back final res
		if start >= len(nums) {
			temp := make([]int, len(subset))
			copy(temp, subset)
			*res = append(*res, temp)
			return
		}
		//Appending subset
		subset = append(subset, nums[start])
		backtracing(nums, start+1, subset, res)
		//Pop
		subset = subset[:len(subset)-1]
		for start+1 < len(nums) {
			//Check if nums[i] == nums[i+1] will skip(i+1) then consider next case
			if nums[start] != nums[start+1] {
				break
			}
			//Ski[ ]
			start += 1
		}
		backtracing(nums, start+1, subset, res)

	}
	//Agurments
	backtracing(nums, 0, subset, &res)
	return res
}
func main() {
	nums1 := []int{1, 2, 2}
	fmt.Println("Subsets for [1, 2, 2]:")
	fmt.Println(subsetsWithDup(nums1))

	// Test case 2: [2, 2, 2]
	nums2 := []int{2, 2, 2}
	fmt.Println("\nSubsets for [2, 2, 2]:")
	fmt.Println(subsetsWithDup(nums2))

	// Test case 3: [1, 2, 3]
	nums3 := []int{1, 2, 3}
	fmt.Println("\nSubsets for [1, 2, 3]:")
	fmt.Println(subsetsWithDup(nums3))
}
