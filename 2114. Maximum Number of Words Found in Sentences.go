package main

import "strings"

//A sentence is a list of words that are separated by a single space with no leading or trailing spaces.
//You are given an array of strings sentences, where each sentences[i] represents a single sentence.
//Return the maximum number of words that appear in a single sentence.

func mostWordsFound(sentences []string) int {
	numberOfWords := 0
	for _, val := range sentences {
		words := strings.Split(val, " ")
		if len(words) > numberOfWords {
			numberOfWords = len(words)
		}
	}
	return numberOfWords
}
