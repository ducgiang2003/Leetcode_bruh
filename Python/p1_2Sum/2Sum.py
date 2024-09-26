#Given an array of integers nums and an integer target,
# then return indices of the two numbers such that they add up to target.

#Ex:nums = [2,7,11,15], target = 9
#Output : [0,1]
#Explanation:  Because nums[0] + nums[1] == 9, we return [0, 1]

#Intuitation:Using hashmap 
from typing import List

class Solution :
    #Assume we have a list [2,7,11,15]
     def twoSum(self, nums: List[int], target: int) -> List[int]:
         map = {}
         
         for index,value in enumerate(nums):
             #Assume target = 9 
             #In the first loop 
             #-> remain = 9 -2 =7
             remain = target-value
             #First checking will check remain = 7 
             #Cause there's no ele in map then it must insert first
             #After insert , check again 
             if remain in map :
                 #And return the answer
                 return [map[remain],index]
             #Insert a key and a value : 
             #K-> Key(value) : 2 , Index = 0
             map[value] = index

#Test function 
def test_two_sum():
    solution = Solution()

    # Test case 1
    nums1 = [2, 7, 11, 15]
    target1 = 9
    expected_output1 = [0, 1]
    assert solution.twoSum(nums1, target1) == expected_output1, f"Test case 1 failed: {solution.twoSum(nums1, target1)}"

    # Test case 2
    nums2 = [3, 2, 4]
    target2 = 6
    expected_output2 = [1, 2]
    assert solution.twoSum(nums2, target2) == expected_output2, f"Test case 2 failed: {solution.twoSum(nums2, target2)}"

    # Test case 3
    nums3 = [3, 3]
    target3 = 6
    expected_output3 = [0, 1]
    assert solution.twoSum(nums3, target3) == expected_output3, f"Test case 3 failed: {solution.twoSum(nums3, target3)}"

    print("All test cases passed!")

# Run the test function
test_two_sum()

             