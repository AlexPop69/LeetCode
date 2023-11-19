package main

//Given an integer array nums, move all 0's to the end of it while maintaining
//the relative order of the non-zero elements.
//Note that you must do this in-place without making a copy of the array.

//Example 1:
//Input: nums = [0,1,0,3,12]
//Output: [1,3,12,0,0]

func moveZeroes(nums []int) {
	arrZero := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			arrZero = append(arrZero, nums[i])
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	nums = append(nums, arrZero...)
}
