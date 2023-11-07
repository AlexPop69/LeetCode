package main

// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.

func majorityElement(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		count := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
		if count > len(nums)/2 {
			result = nums[i]
		}
	}
	return result
}

//the best solution, but it`s not my
// func majorityElement(nums []int) int {
// 	key := make(map[int]int)
// 	for _, n := range nums {
// 		key[n]++
// 	}

// 	res := 0
// 	max := 0
// 	for i, n := range key {
// 		if n > max {
// 			max = n
// 			res = i
// 		}
// 	}

// 	return res
// }
