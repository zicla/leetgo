package main

import "fmt"

func numDecodings(s string) int {
	return numDecodingsRecursive(s, 0, len(s)-1)
}

func numDecodingsRecursive(s string, start int, end int) int {
	if start > end {
		return 0
	}
	if start == end {
		if s[start] == '0' {
			return 0
		} else {
			return 1
		}
	}

	ch0 := s[start]
	ch1 := s[end-1]
	ch2 := s[end]

	if ch0 == '0' {
		return 0
	}

	if ch1 == '0' && ch2 == '0' {
		return 0
	}

	if ch2 == '0' {

		if ch1 == '1' || ch1 == '2' {

			if end-start == 1 {
				return 1
			}

			pre := numDecodingsRecursive(s, start, end-2)
			if pre == 0 {
				return 0
			}
			return pre
		} else {
			return 0
		}

	} else if ch1 == '1' || (ch1 == '2' && (ch2 == '1' || ch2 == '2' || ch2 == '3' || ch2 == '4' || ch2 == '5' || ch2 == '6')) {
		pre := numDecodingsRecursive(s, start, end-1)
		if pre == 0 {
			return 0
		}
		return pre + 1
	} else {
		pre := numDecodingsRecursive(s, start, end-1)
		if pre == 0 {
			return 0
		}
		return pre
	}

}

func main() {

	fmt.Printf("%v\n", numDecodings("121"))
	fmt.Printf("%v\n", numDecodings("0"))
	fmt.Printf("%v\n", numDecodings("10"))
	fmt.Printf("%v\n", numDecodings("11"))
	fmt.Printf("%v\n", numDecodings("00"))
	fmt.Printf("%v\n", numDecodings("01"))
	fmt.Printf("%v\n", numDecodings("012"))
	fmt.Printf("%v\n", numDecodings("100"))
	fmt.Printf("%v\n", numDecodings("110"))
	fmt.Printf("%v\n", numDecodings("230"))
	fmt.Printf("%v\n", numDecodings("12120"))
	fmt.Printf("%v\n", numDecodings("1212"))

}
