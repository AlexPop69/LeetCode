package main

import (
	"strings"
	"unicode"
)

/*A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and
removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters
include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.*/

func isPalindrome(s string) bool {
	newS := ""
	if s != "" {
		for _, symbol := range s {
			if !unicode.IsLetter(rune(symbol)) && !unicode.IsNumber(rune(symbol)) {
				s = strings.Replace(s, string(symbol), "", -1)
				continue
			}
			newS = string(symbol) + newS
		}
	}
	return strings.EqualFold(strings.ToLower(s), strings.ToLower(newS))
}
