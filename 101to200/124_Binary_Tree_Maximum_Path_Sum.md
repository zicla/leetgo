package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	N := len(s)
	if N == 0 {
		return true
	}

	var isEqual = func(a byte, b byte) bool {

		if a == b {
			return true
		} else {
			if ((a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z')) && ((b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')) {

				if a > b {
					return a-b == 'a'-'A'
				} else {
					return b-a == 'a'-'A'
				}

			} else {
				return false
			}
		}

	}

	p := 0
	q := N - 1
	for p < q {

		for p < N && !((s[p] >= 'a' && s[p] <= 'z') || (s[p] >= 'A' && s[p] <= 'Z') || (s[p] >= '0' && s[p] <= '9')) {
			p++
		}
		for q >= 0 && !((s[q] >= 'a' && s[q] <= 'z') || (s[q] >= 'A' && s[q] <= 'Z') || (s[q] >= '0' && s[q] <= '9')) {
			q--
		}

		if p > q {
			break
		}

		if isEqual(s[p], s[q]) {
			p++
			q--
		} else {
			return false
		}

	}
	return true

}

func main() {

	fmt.Printf("%v\n", isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Printf("%v\n", isPalindrome(".,"))
	fmt.Printf("%v\n", isPalindrome("0P"))
	fmt.Printf("%v\n", isPalindrome("race a car"))

}
