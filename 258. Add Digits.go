package main

//Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.
func addDigits(num int) int {
	if num >= 10 {
		for num > 9 {
			num = num%10 + num/10
		}
	}
	return num
}
