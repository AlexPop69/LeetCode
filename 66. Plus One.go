package main

import (
	"math/big"
	"strconv"
)

func plusOne(digits []int) []int {
	str := ""
	for i := 0; i < len(digits); i++ {
		a := strconv.Itoa(digits[i])
		str += a
	}

	intD, _ := new(big.Int).SetString(str, 10)

	plusOne := new(big.Int).SetInt64(1)
	intD = intD.Add(intD, plusOne)

	strD := intD.Text(10)

	var newDigits []int

	for _, val := range strD {
		a, _ := strconv.Atoi(string(val))
		newDigits = append(newDigits, a)
	}

	return newDigits
}
