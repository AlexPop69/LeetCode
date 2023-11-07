package main

import (
	"sort"
)

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.

// Runtime: 4ms (Beats 73.04%of users with Go);
// Memory: 3.06mb (Beats 25.32%of users with Go)
func isAnagram(s string, t string) bool {
	arrayS := []byte(s)
	arrayT := []byte(t)
	sort.Slice(arrayS, func(i, j int) bool { return arrayS[i] < arrayS[j] })
	sort.Slice(arrayT, func(i, j int) bool { return arrayT[i] < arrayT[j] })
	return string(arrayS) == string(arrayT)
}
