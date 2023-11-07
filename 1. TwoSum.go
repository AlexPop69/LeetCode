package main

//not my solution
// func twoSum(nums []int, target int) []int {
// 	numsMap := make(map[int]int)

// 	for i := 0; i < len(nums); i++ {
// 		reqNum := target - nums[i]
// 		reqNumIdx, isPresent := numsMap[reqNum]

// 		if isPresent {
// 			return []int{i, reqNumIdx}
// 		}

// 		numsMap[nums[i]] = i
// 	}

// 	return []int{}
// }

func twoSum(nums []int, target int) []int {
	var res []int
	for i := range nums {
		for idx, val := range nums {
			if val+nums[i] == target && i != idx {
				res = append(res, idx)
			}
		}
	}
	return res
}
