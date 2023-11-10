package main

// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.
func majorityElement(nums []int) int {
	result := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		result[nums[i]] += 1
		if result[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return 0
}
