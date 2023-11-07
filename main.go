package main

import "fmt"

func main() {
	//1.TwoSum
	nums := []int{0, 4, 3, 0}
	target := 0
	fmt.Println(twoSum(nums, target))

	//9.Palindrome Numbder
	number := 1221
	fmt.Println(isPalindrome(number))

	//27. Remove Element
	var numbers = []int{1, 2, 3, 4, 5}
	el := 3
	fmt.Printf("%v, nums= %v", removeElement(numbers, el), numbers)

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
	fmt.Println(mySqrt(x))

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

	//344. Reverse String
	reversStr := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(reversStr)
	fmt.Println(string(reversStr))

	//258. Add Digits
	addNumber := 38 //2
	fmt.Println(addDigits(addNumber))
}
