package main

import (
	"sort"
)

func removeElement(nums []int, val int) int {
	count := len(nums)
	for i, v := range nums {
		if v == val {
			nums[i] = 101
			count--
		}
	}

	sort.Ints(nums)
	return count
}
