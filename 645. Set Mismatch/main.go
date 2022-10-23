package main

import (
	"fmt"
	"sort"
)

func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	var res []int
	var numTmp int = 1
	var dupTmp int
	for i := 0; i < len(nums); i++ {
		if nums[i] == numTmp {
			numTmp++
		}
		if i < len(nums)-1 {
			if nums[i] == nums[i+1] {
				dupTmp = nums[i]
			}
		}
	}
	res = append(res, dupTmp)
	res = append(res, numTmp)
	return res
}

func main() {
	nums := []int{1, 5, 3, 2, 2, 7, 6, 4, 8, 9}
	fmt.Println(findErrorNums(nums))
}
