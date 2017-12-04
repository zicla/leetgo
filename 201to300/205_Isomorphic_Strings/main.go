package main

import (
	"fmt"
)

/*
返回true需要满足两个条件：
1、不能出现s中的一个字符对应到t中两个不同的字符
2、不能出现s中的两个不同字符对应到t中同一个字符
 */
func isIsomorphic(s string, t string) bool {
	N := len(s)
	if N != len(t) {
		return false
	}
	if N == 0 {
		return true
	}
	s2t := make([]int, 127)
	t2s := make([]int, 127)
	for k, _ := range s2t {
		s2t[k] = -1
		t2s[k] = -1
	}
	for i := 0; i < N; i++ {
		sChar := int(s[i])
		tChar := int(t[i])

		//s中同一个，却要对应t中两个不同的。
		if s2t[sChar] != -1 && s2t[sChar] != tChar {
			return false
		}

		//s中的两个不同字符对应到t中同一个字符
		if t2s[tChar] != -1 && t2s[tChar] != sChar {
			return false
		}


		s2t[sChar] = tChar
		t2s[tChar] = sChar
	}

	return true

}

func main() {

	fmt.Printf("%v \n", isIsomorphic("ab", "aa"))
	fmt.Printf("%v \n", isIsomorphic("egg", "add"))
	fmt.Printf("%v \n", isIsomorphic("foo", "bar"))
	fmt.Printf("%v \n", isIsomorphic("paper", "title"))
	fmt.Printf("%v \n", isIsomorphic("12", "23"))
	fmt.Printf("%v \n", isIsomorphic("13", "42"))
	fmt.Printf("%v \n", isIsomorphic("ab", "ca"))

}
