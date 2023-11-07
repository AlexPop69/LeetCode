package main

// Given a sorted array of distinct integers and a target value,
// return the index if the target is found. If not, return the index
// where it would be if it were inserted in order.
// You must write an algorithm with O(log n) runtime complexity.

func searchInsert(nums []int, target int) int {
	var result int
	if nums[len(nums)-1] < target {
		result = len(nums)
	} else {
		for i, v := range nums {
			if v == target {
				result = i
				break
			} else if v > target {
				result = i
				break
			}
		}
	}
	return result
}
