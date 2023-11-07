package main

func mySqrt(x int) int {
	var start, end, temp, v, result int

	end = x
	for start <= end {
		temp = (start + end) / 2
		v = temp * temp

		if v == x {
			result = temp
			break
		}

		if v > x {
			end = temp - 1
		} else {
			result = temp
			start = temp + 1
		}
	}

	return result
}
