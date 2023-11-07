package main

// Write a function that reverses a string. The input string is given as an array of characters s.
// You must do this by modifying the input array in-place with O(1) extra memory

//Runtime: 32ms (Beats 48.52%of users with Go)
//Memory: 6.50 (Beats 95.37%of users with Go)

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		temp := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = temp

	}
}
