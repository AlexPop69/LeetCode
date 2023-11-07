package main

import (
	"math/big"
	"strconv"
	"strings"
)

func plusOne(digits []int) []int {
	var s []string
	for i := 0; i < len(digits); i++ {
		s = append(s, strconv.Itoa(digits[i]))
	}

	text := strings.Join(s, "")

	intD, _ := new(big.Int).SetString(text, 10)

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
