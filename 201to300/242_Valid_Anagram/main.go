package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	var dict = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		if dict[t[i]] == 0 {
			return false
		}
		dict[t[i]]--

	}
	return true

}

func main() {

	fmt.Printf("%v\n", isAnagram("anagram", "nagaram"))
	fmt.Printf("%v\n", isAnagram("rat", "car"))
}
