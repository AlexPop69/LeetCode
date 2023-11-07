package main

//Given an integer array nums, return true if any value appears
//at least twice in the array, and return false if every element is distinct.
func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		count := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
				return true
			}
		}
	}
	return false
}
