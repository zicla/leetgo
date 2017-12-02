package main

import "fmt"

func isScramble(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	if s1 == "" && s2 == "" {
		return true
	}

	return recursive1(s1, s2)
}

func recursive1(s1 string, s2 string) bool {

	length := len(s1)

	if length == 1 {
		return s1 == s2
	}

	//通过判断各个字符的数量来进行剪枝
	map1 := make(map[byte]int)
	for i := 0; i < length; i++ {
		map1[s1[i]]++
	}

	for i := 0; i < length; i++ {
		map1[s2[i]]--
		if map1[s2[i]] < 0 {
			return false
		}
	}

	//划分两个字符串
	for i := 1; i < length; i++ {

		if recursive1(s1[0:i], s2[0:i]) && recursive1(s1[i:length], s2[i:length]) {
			return true
		}

		if recursive1(s1[0:i], s2[length-i:length]) && recursive1(s1[i:length], s2[0:length-i]) {
			return true
		}

	}

	return false
}

func main() {

	fmt.Printf("%v\n", isScramble("a", "a"))
	fmt.Printf("%v\n", isScramble("abc", "cba"))
	fmt.Printf("%v\n", isScramble("ccabcbabcbabbbbcbb", "bbbbabccccbbbabcba"))

}
