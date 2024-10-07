#Input (target = 7 and candicates or nums  [2,3,6,7])
#The element in output can be duplicate as long as the total of List == target
#Output :[[2,2,3],[7]]

from typing import List
class Solution:
    def combinationSum(self, candicates: List[int], target: int) -> List[List[int]]:
        result = []
        #Total is sum of ele in result 
        def backtracking(index,sub,total):
            if total == target :
                result.append(sub.copy())
            #If range of index greater than len(candicates) or total > target then return (back to previous case)
            if index>=len(candicates) or total > target:
                return
            #Appending candicates at index into sub 
            #i=0 -> [2]
            sub.append(candicates[index])
            #Inspecting next case
            backtracking(index,sub,total +candicates[index])
            sub.pop()
           # Skip the current candidate and move to the n 
            backtracking(index+1,sub,total)
        backtracking(0,[],0)
        return result
    
    
def test_combinationSum():
    sol = Solution()
    
    # Test case 1
    candidates = [2, 3, 6, 7]
    target = 7
    expected_output = [[2, 2, 3], [7]]
    assert sol.combinationSum(candidates, target) == expected_output, f"Test case 1 failed: {sol.combinationSum(candidates, target)}"
    
    # Test case 2
    candidates = [2, 3, 5]
    target = 8
    expected_output = [[2, 2, 2, 2], [2, 3, 3], [3, 5]]
    assert sol.combinationSum(candidates, target) == expected_output, f"Test case 2 failed: {sol.combinationSum(candidates, target)}"
    
    # Test case 3
    candidates = [2]
    target = 1
    expected_output = []
    assert sol.combinationSum(candidates, target) == expected_output, f"Test case 3 failed: {sol.combinationSum(candidates, target)}"
    
    # Test case 4
    candidates = [1]
    target = 1
    expected_output = [[1]]
    assert sol.combinationSum(candidates, target) == expected_output, f"Test case 4 failed: {sol.combinationSum(candidates, target)}"
    
    # Test case 5
    candidates = [1]
    target = 2
    expected_output = [[1, 1]]
    assert sol.combinationSum(candidates, target) == expected_output, f"Test case 5 failed: {sol.combinationSum(candidates, target)}"
    
    print("All test cases passed!")

# Run tests
test_combinationSum()
