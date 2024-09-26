package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	//Index loops
	for i := 0; i < len(nums); i++ {
		//Find remain each loop
		remain := target - nums[i]
		//If hashMap have remain val then return value of key
		if _, ok := hashMap[remain]; ok {
			return []int{hashMap[remain], i}
		}
		//Else must insert key,val into hashMap
		hashMap[nums[i]] = i
	}
	return nil
}
func main() {
	// Example inputs
	nums1 := []int{2, 7, 11, 15}
	target1 := 9

	nums2 := []int{3, 2, 4}
	target2 := 6

	nums3 := []int{3, 3}
	target3 := 6

	// Test twoSum with the inputs
	fmt.Println("Test Case 1:")
	fmt.Printf("Input: nums = %v, target = %d\n", nums1, target1)
	fmt.Printf("Output: %v\n", twoSum(nums1, target1))

	fmt.Println("Test Case 2:")
	fmt.Printf("Input: nums = %v, target = %d\n", nums2, target2)
	fmt.Printf("Output: %v\n", twoSum(nums2, target2))

	fmt.Println("Test Case 3:")
	fmt.Printf("Input: nums = %v, target = %d\n", nums3, target3)
	fmt.Printf("Output: %v\n", twoSum(nums3, target3))
}
