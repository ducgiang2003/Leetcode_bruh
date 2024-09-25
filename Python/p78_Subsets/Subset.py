from typing import List
#Give a list without duplicate number and solve it into subsets
#[1,2,3] -> [[],[1],[2],[1,2],[3],[1,2,3]]
#Using backtracking
class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        #return result
        res = []
        subset = []
        def backtracking (i):
            if i>=len(nums):
                #Append subsets into res
                res.append(subset.copy())
                #then return
                return
            #Appeding nums[i] into subsets
            subset.append(nums[i])
            backtracking(i+1)
            #Backtracking
            subset.pop()
            backtracking(i+1)
        backtracking(0)
        return res
    
test_subsets()


#Test case for problem 
def test_subsets():
    s = Solution()  # Assuming the function is part of a class named Solution
    
    # Helper function to compare two lists of lists irrespective of order
    def compare_unordered_lists(list1, list2):
        return set(map(tuple, list1)) == set(map(tuple, list2))
    
    # Test case 1: Empty list
    assert compare_unordered_lists(s.subsets([]), [[]])
    
    # Test case 2: List with one element
    assert compare_unordered_lists(s.subsets([1]), [[], [1]])
    
    # Test case 3: List with two elements
    assert compare_unordered_lists(s.subsets([1, 2]), [[], [1], [2], [1, 2]])
    
    # Test case 4: List with three elements
    assert compare_unordered_lists(s.subsets([1, 2, 3]), [[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]])
    print("All test cases passed!")

    
    