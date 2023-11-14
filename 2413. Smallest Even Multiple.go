package main

//Given a positive integer n, return the smallest positive integer that is a multiple of both 2 and n.
func smallestEvenMultiple(n int) int {
	k := n % 2
	if k == 0 {
		return n
	}
	return n * 2
}
