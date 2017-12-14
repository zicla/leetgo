package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {

	if len(s) == 0 && len(p) == 0 {
		return true
	} else if len(s) == 0 {
		if len(p) == 1 {
			return false
		} else {
			return p[1] == '*' && isMatch(s, p[2:])
		}
	} else if len(p) == 0 {
		return false
	} else if len(s) == 1 && len(p) == 1 {
		if p[0] == '.' {
			return true
		} else {
			return s[0] == p[0]
		}
	} else if len(s) == 1 {
		if p[1] == '*' {
			return isMatch(s, p[2:]) || ((s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p[2:]))
		} else {
			if s[0] == p[0] || p[0] == '.' {
				return isMatch(s[1:], p[1:])
			} else {
				return false
			}

		}
	} else if len(p) == 1 {
		return false
	} else {

		if p[0] == '.' {

			if p[1] == '*' {
				//一个不消耗
				b := isMatch(s, p[2:])
				i := 0
				for i < len(s) {
					b = b || isMatch(s[i+1:], p[2:])
					i++
					if b {
						return b
					}
				}

				return b

			} else {
				return isMatch(s[1:], p[1:])
			}

		} else {

			if p[1] == '*' {

				//一个都不消耗
				b := isMatch(s, p[2:])
				i := 0
				for i < len(s) && s[i] == p[0] {
					b = b || isMatch(s[i+1:], p[2:])
					i++
					if b {
						return b
					}
				}
				return b

			} else {
				//只能是字母，不可能是*
				if s[0] == p[0] {
					return isMatch(s[1:], p[1:])
				} else {
					return false
				}
			}

		}

	}

}
func main() {

	/*
		isMatch("aa","a") → false
	isMatch("aa","aa") → true
	isMatch("aaa","aa") → false
	isMatch("aa", "a*") → true
	isMatch("aa", ".*") → true
	isMatch("ab", ".*") → true
	isMatch("aab", "c*a*b") → true
	*/

	fmt.Printf("%v\n", isMatch("ab", ".*..c*")) //true
	fmt.Printf("%v\n", isMatch("a", ".*"))      //true
	fmt.Printf("%v\n", isMatch("a", "ab*"))     //true
	fmt.Printf("%v\n", isMatch("aa", ".*"))     //true
	fmt.Printf("%v\n", isMatch("ab", ".*"))     //true

	fmt.Printf("%v\n", isMatch("aa", "a"))   //false
	fmt.Printf("%v\n", isMatch("aa", "aa"))  //true
	fmt.Printf("%v\n", isMatch("aaa", "aa")) //false
	fmt.Printf("%v\n", isMatch("aa", "a*"))  //true

	fmt.Printf("%v\n", isMatch("aab", "c*a*b")) //true
}
