package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	trimS := strings.TrimSpace(s)
	var arrayS []string = strings.Split(trimS, " ")
	return len(arrayS[len(arrayS)-1])
}
