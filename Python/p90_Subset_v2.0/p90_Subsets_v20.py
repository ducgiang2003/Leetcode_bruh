from typing import List
import unittest
class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        res =[]
        #First must sort the arr(nums)
        nums.sort()
        
        def backtracking(start,subset):
            if start >= len(nums):
                res.append(subset[::])
                return
            #Appending nums[start] into ele of subset
            subset.append(nums[start])
            backtracking(start+1,subset)
            subset.pop()
            while start +1 >len(nums) and nums[start]==nums[start+1]:
                start+=1
            backtracking(start+1,subset)
        backtracking(0,[])
        return res 
    



if __name__ == "__main__":
    solution = Solution()
    nums = [1, 2, 2]
    result = solution.subsetsWithDup(nums)
    print(result)
    