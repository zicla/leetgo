package main

import (
	"fmt"
)

//反转一个字符串
func reverse(s string) string {
	bs := []byte(s)
	N := len(s)
	for i := 0; i < N/2; i++ {
		bs[i], bs[N-1-i] = bs[N-1-i], bs[i]
	}
	return string(bs)
}

//判断是否是回文字符串
func valid(s string) bool {
	N := len(s)
	p := 0
	q := N - 1
	for p < N {
		if s[p] == s[q] {
			p++
			q--
		} else {
			return false
		}
	}
	return true
}
func shortestPalindrome(s string) string {
	N := len(s)
	if N == 0 || N == 1 {
		return s
	}

	for i := N - 1; i >= 0; i-- {
		if valid(s[:i+1]) {
			return reverse(s[i+1:]) + s
		}
	}

	return ""
}

func main() {

	fmt.Printf("%v \n", shortestPalindrome("aabbc"))
	fmt.Printf("%v \n", shortestPalindrome("a"))
	fmt.Printf("%v \n", shortestPalindrome("abcd"))
	fmt.Printf("%v \n", shortestPalindrome("babcd"))
}
