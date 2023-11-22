package main

import (
	"strconv"
)

func isPalindromeNumber(x int) bool {
	strX := strconv.Itoa(x)
	var s []rune

	for i := len(strX) - 1; i >= 0; i-- {
		s = append(s, rune(strX[i]))
	}

	//	fmt.Println(string(s))
	if strX == string(s) {
		return true
	} else {
		return false
	}

}
