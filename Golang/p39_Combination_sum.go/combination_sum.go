package main

import (
	"fmt"
	"sort"
)

// Input (target = 7 and candicates or nums  [2,3,6,7])
// The element in output can be duplicate as long as the total of List == target
// Output :[[2,2,3],[7]]

func combinationSum(candicates []int, target int) [][]int {
	res, sub := [][]int{}, []int{}
	//Sort array
	sort.Ints(candicates)

	var backtracking func(i int, sub []int, target int, candicates []int, res *[][]int, total int)
	backtracking = func(i int, sub []int, target int, candicates []int, res *[][]int, total int) {
		// if total == target then return output
		if total == target {
			temp := make([]int, len(sub))
			copy(temp, sub)
			*res = append(*res, temp)
			return
		}
		//Return back to previous case
		if i >= len(candicates) || total > target {
			return
		}
		sub = append(sub, candicates[i])
		backtracking(i, sub, target, candicates, res, total+candicates[i])
		//Pop
		sub = sub[:len(sub)-1]
		//Skip the current candidate and move to the n
		backtracking(i+1, sub, target, candicates, res, total)
	}
	backtracking(0, sub, target, candicates, &res, 0)
	return res
}
func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
