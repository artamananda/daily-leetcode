package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i)
				res = append(res, j)
				return res
			}
		}
	}
	return res
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}
