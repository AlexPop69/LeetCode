package main

import (
	"fmt"
)

func main() {
	//1.TwoSum
	nums := []int{0, 4, 3, 0}
	target := 0
	fmt.Printf("TwoSum of %v with target %d is %v\n", nums, target, twoSum(nums, target))

	//9.Palindrome Numbder
	number := 1221
	fmt.Printf("Palindrome Number: %d is %v\n", number, isPalindromeNumber(number))

	//13. Roman to Integer
	fmt.Println(romanToInt("MCMXCIV")) // 1994

	//27. Remove Element
	var numbers = []int{1, 2, 3, 4, 5}
	el := 3
	fmt.Printf("Remove %v from %v. Result: %v\n", el, numbers, removeElement(numbers, el))

	//26. Remove Duplicates from Sorted Array
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	//28. Find the Index of the First Occurrence in a String
	haystack := "leetcode"
	needle := "leeto"
	fmt.Println(strStr(haystack, needle))

	//35. Search Insert Position
	numerals := []int{1, 3, 5, 6}
	searchIndex := 2
	fmt.Println(searchInsert(numerals, searchIndex))

	//58. Length of Last Word
	word := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(word))

	//66. Plus One
	var digits = []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	fmt.Println(plusOne(digits))

	//69. Sqrt(x)
	x := 8
	fmt.Printf("%d after sqrt is %d\n", x, mySqrt(x))

	//88. Merge Sorted Array
	mer1, merM, mer2, merN := []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3
	merge(mer1, merM, mer2, merN)
	fmt.Printf("88: Merged and sorted two arrays: %v\n", mer1)

	//125. Valid Palindrome
	strForPalindrome := ".a"
	fmt.Printf("125: \"%s\" is %v palindrome\n", strForPalindrome, isPalindrome(strForPalindrome))

	//136. Single Number
	digit := []int{1, 2, 4, -1, 3, 2, 4, 1, 3}
	fmt.Println(singleNumber(digit))

	//169. Majority Element
	numeral := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(numeral))

	//217. Contains Duplicate
	num1 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}    // true
	num2 := []int{1, 2, 3, 4}                      //false
	num3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 5, 4, 2} // true
	fmt.Println(containsDuplicate(num1))
	fmt.Println(containsDuplicate(num2))
	fmt.Println(containsDuplicate(num3))

	//242. Valid Anagram
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))

	//258. Add Digits
	addNumber := 38 //2
	fmt.Println(addDigits(addNumber))

	//283. Move Zeroes
	var forMoveZeroes = []int{0, 1, 0, 3, 12}
	moveZeroes(forMoveZeroes)
	fmt.Println(forMoveZeroes) //1,3,12,0,0

	//344. Reverse String
	reversStr := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(reversStr)
	fmt.Println(string(reversStr))

	//771. Jewels and Stones
	fmt.Println(numJewelsInStones("aA", "aAAbbbb")) //3

	//1281. Subtract the Product and Sum of Digits of an Integer
	fmt.Println(subtractProductAndSum(234))

	//1365. How Many Numbers Are Smaller Than the Current Number
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3})) //4,0,1,1,3

	//1431. Kids With the Greatest Number of Candies
	fmt.Printf("1431: %v\n", kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)) //[true,true,true,false,true]

	//1480. Running Sum of 1d Array
	fmt.Println(runningSum([]int{1, 2, 3, 4}))

	//1528. Shuffle String
	st := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	fmt.Println(restoreString(st, indices))

	//1816. Truncate Sentence (2 variants)
	string1 := "chopper is not a tanuki"
	k := 2
	fmt.Println(truncateSentence(string1, k))
	string1 = "Hello how are you Contestant"
	k = 4
	fmt.Println(truncateSentence1(string1, k))

	//1859. Sorting the Sentence
	fmt.Printf("1859. Sorting the Sentence: %s\n", sortSentence("sentence4 a3 is2 This1"))

	//1929. Concatenation of Array
	notAns := []int{1, 2, 1}
	fmt.Println(getConcatenation(notAns))

	//2011. Final Value of Variable After Performing Operations
	var aS = []string{"++x", "x++", "x--", "x++"}
	fmt.Println(finalValueAfterOperations(aS))

	//2114. Maximum Number of Words Found in Sentences
	fmt.Println(mostWordsFound([]string{
		"alice and bob love leetcode",
		"i think so too",
		"this is great thanks very much",
	})) //6

	//2413. Smallest Even Multiple
	fmt.Println(smallestEvenMultiple(6))

	//2469. Convert the Temperature
	celsius := 36.6
	fmt.Println(convertTemperature(celsius))

	//2574. Left and Right Sum Differences
	fmt.Println(leftRightDifference([]int{10, 4, 8, 3}))

	//2652. Sum Multiples
	toN := 13 //52
	fmt.Println(sumOfMultiples(toN))

}
