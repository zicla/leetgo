package main

import "fmt"

func longestPalindrome(s string) string {

	N := len(s)
	if (N == 0) {
		return ""
	} else if (N == 1) {
		return s
	}

	palindrome := s[0:1]

	var left int
	var right int

	pivot := 1

	for pivot < 2*N-2 {

		if pivot%2 == 0 {
			left = pivot/2 - 1
			right = pivot/2 + 1
		} else {
			left = (pivot+1)/2 - 1
			right = (pivot + 1) / 2
		}

		for left >= 0 && right < N && s[left] == s[right] {

			if len(palindrome) < (right + 1 - left) {
				palindrome = s[left:right+1]
			}

			left --
			right ++
		}

		pivot++

	}

	return palindrome
}

func main() {

	str := "babad"

	fmt.Println(longestPalindrome(str))

	str = "cdbbbbbbbbbbbbbbd"

	fmt.Println(longestPalindrome(str))

	//fmt.Println(str[1:2])

}
