package main

import (
	"strconv"
	"strings"
	"unicode"
)

/*A sentence is a list of words that are separated by a single space with no leading or trailing spaces.
Each word consists of lowercase and uppercase English letters.
A sentence can be shuffled by appending the 1-indexed word position to each word
then rearranging the wordsin the sentence.

#For example, the sentence "This is a sentence" can be shuffled as
"sentence4 a3 is2 This1" or "is2 sentence4 This1 a3".
Given a shuffled sentence s containing no more than 9 words, reconstruct and return the original sentence.*/

func sortSentence(s string) string {
	result := make([]string, strings.Count(s, " ")+1)
	var num int
	res := ""
	for _, v := range s {
		if unicode.IsDigit(v) {
			num, _ = strconv.Atoi(string(v))
		} else if v != 32 {
			res += string(v)
		}
		if v == 32 || v == rune(s[len(s)-1]) {
			result[num-1] = res
			res = ""
		}
	}
	return strings.Join(result, " ")
}

// func sortSentence(s string) string {
// 	s = "b2 c3 a1"
// 	result := make([]string, strings.Count(s, " ")+1)
// 	var num int
// 	res := ""
// 	for _, v := range s {
// 		switch {
// 		case unicode.IsDigit(v):
// 			num, _ = strconv.Atoi(string(v))
// 			result[num-1] = res
// 			res = ""
// 		case v == 32:
// 			continue
// 		default:
// 			res += string(v)
// 		}
// 	}
// 	return strings.Join(result, " ")
// }
