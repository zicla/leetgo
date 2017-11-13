package main

import "fmt"

func numDecodings(s string) int {
	N := len(s)
	if N == 0 {
		return 0
	}

	F := make([]int, N, N)
	for i := 0; i < N; i++ {

		if i == 0 {
			if s[0] == '0' {
				F[0] = 0
				return 0
			} else {
				F[0] = 1
				continue
			}
		} else if i == 1 {
			if s[i] == '0' {
				if s[i-1] == '1' || s[i-1] == '2' {
					F[i] = 1
				} else {
					return 0
				}
			} else {
				if s[i-1] == '1' || (s[i-1] == '2' && (s[i] == '1' || s[i] == '2' || s[i] == '3' || s[i] == '4' || s[i] == '5' || s[i] == '6')) {
					F[i] = 2
				} else {
					F[i] = 1
				}
			}
		} else {
			if s[i] == '0' {
				if s[i-1] == '1' || s[i-1] == '2' {
					//F[i] = F[i-2]*2 + F[i-1] - F[i-2] //因为F[i-2]*2 + F[i-1]正好重复了 F[i-2]
					F[i] = F[i-2]
				} else {
					return 0
				}
			} else {
				if s[i-1] == '0' {
					F[i] = F[i-1]
				} else if s[i-1] == '1' {
					F[i] = F[i-2] + F[i-1]
				} else if s[i-1] == '2' {
					if s[i] == '1' || s[i] == '2' || s[i] == '3' || s[i] == '4' || s[i] == '5' || s[i] == '6' {
						F[i] = F[i-2] + F[i-1]
					} else {
						F[i] = F[i-1]
					}
				} else {
					F[i] = F[i-1]
				}

			}
		}

	}
	return F[N-1]

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
