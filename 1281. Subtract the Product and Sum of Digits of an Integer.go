package main

import "strconv"

//Given an integer number n, return the difference between
//the product of its digits and the sum of its digits.

//Example 1:
//Input: n = 234
//Output: 15
//Explanation:
//Product of digits = 2 * 3 * 4 = 24
//Sum of digits = 2 + 3 + 4 = 9
//Result = 24 - 9 = 15

func subtractProductAndSum(n int) int {
	strN := strconv.Itoa(n)
	sum := 0
	prod := 1
	//result := 0
	for _, val := range strN {
		num, _ := strconv.Atoi(string(val))
		sum += num
		prod *= num
	}
	return prod - sum
}
