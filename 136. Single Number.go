package main

import (
	"sort"
)

// Given a non-empty array of integers nums, every element
// appears twice except for one. Find that single one.
// You must implement a solution with a linear runtime complexity and use only constant extra space.
func singleNumber(nums []int) int {
	var temp int
	for i := 0; i < len(nums); i++ {
		temp = nums[i]
		for j := i + 1; j < len(nums); j++ {
			if temp == nums[j] {
				nums[i], nums[j] = 2147483648, 2147483648
			}
		}
	}
	sort.Ints(nums)
	return nums[0]
}
