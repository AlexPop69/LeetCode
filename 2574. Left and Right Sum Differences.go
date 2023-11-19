package main

//Given a 0-indexed integer array nums, find a 0-indexed integer array answer where:

//answer.length == nums.length.
//answer[i] = |leftSum[i] - rightSum[i]|.
//Where:
//leftSum[i] is the sum of elements to the left of the index i in the array nums.
//If there is no such element, leftSum[i] = 0.
//rightSum[i] is the sum of elements to the right of the index i in the array nums.
//If there is no such element, rightSum[i] = 0.
//Return the array answer.

//Example 1:
//Input: nums = [10,4,8,3]
//Output: [15,1,11,22]
//Explanation: The array leftSum is [0,10,14,22] and the array rightSum is [15,11,3,0].
//The array answer is [|0 - 15|,|10 - 11|,|14 - 3|,|22 - 0|] = [15,1,11,22].

func leftRightDifference(nums []int) []int {
	ans := make([]int, len(nums))
	leftSum, rightSum := make([]int, 1), make([]int, 1)

	for i := 1; i < len(nums); i++ {
		leftSum = append(leftSum, leftSum[i-1]+nums[i-1])
		rightSum = append(rightSum, rightSum[i-1]+nums[len(nums)-i])
	}

	for j := 1; j <= len(ans); j++ {
		if leftSum[j-1] > rightSum[len(nums)-j] {
			ans[j-1] = leftSum[j-1] - rightSum[len(nums)-j]
		} else {
			ans[j-1] = rightSum[len(nums)-j] - leftSum[j-1]
		}
	}
	return ans
}
