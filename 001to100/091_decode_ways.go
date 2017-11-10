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

	ch1 := s[end-1]
	ch2 := s[end]

	if end-start >= 2 {

		if ch1 == '1' || (ch1 == '2' && (ch2 == '0' || ch2 == '1' || ch2 == '2' || ch2 == '3' || ch2 == '4' || ch2 == '5' || ch2 == '6')) {
			return numDecodingsRecursive(s, start, end-1) + 1
		} else {
			return numDecodingsRecursive(s, start, end-1)
		}

	} else {

		if s[end-1] == '0' {

		}

		if (ch1 == '1' || (ch1 == '2' && (ch2 == '0' || ch2 == '1' || ch2 == '2' || ch2 == '3' || ch2 == '4' || ch2 == '5' || ch2 == '6'))) && end-2 > start {
			return numDecodingsRecursive(s, start, end-1) + 1
		} else {
			return numDecodingsRecursive(s, start, end-1)
		}

	}

}

func main() {

	fmt.Printf("%v\n", numDecodings("121"))
	fmt.Printf("%v\n", numDecodings("0"))
	fmt.Printf("%v\n", numDecodings("10"))
	fmt.Printf("%v\n", numDecodings("11"))

}
