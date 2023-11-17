package main

//Given the array nums, for each nums[i] find out how many numbers
//in the array are smaller than it. That is,
//for each nums[i] you have to count the number of valid j's such that j != i and nums[j] < nums[i].
//Return the answer in an array.

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	for idx, val := range nums {
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] < val {
				count++
			}
		}
		result[idx] = count
	}
	return result
}
