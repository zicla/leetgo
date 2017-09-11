package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	res := 0
	N := len(s)

	for p1 := 0; p1 < N; p1++ {

		dict := [256]uint8{}

		for p2 := p1; p2 < N; p2++ {

			bit := dict[s[p2]]

			if bit == 1 {
				if (p2 - p1) > res {
					res = p2 - p1
				}
				break
			} else {
				dict[s[p2]] = 1
			}

			if p2 == N-1 {
				if (p2 - p1 + 1) > res {
					res = p2 - p1 + 1
				}
			}
		}
	}

	return res

}
func main() {

	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("c"))

}
