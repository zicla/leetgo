package main

import (
	"fmt"
)

func shortestPalindrome(s string) string {
	N := len(s)
	if N == 0 || N == 1 {
		return s
	}

	//中间的index
	mid := (N - 1) / 2
	//从mid开始往左边扩展。

	//反转到左边的字符串
	start1 := 1
	//继续呆在右边的字符串
	start2 := 0
	if N%2 == 1 {
		//奇数
		//中心点，可能是数字，可能是间隙。
		for i := 0; i <= N-1; i++ {
			if i%2 == 0 {
				//数字为中心
				p := mid - i/2 - 1
				q := mid - i/2 + 1
				for p >= 0 {
					if s[p] == s[q] {
						p--
						q++
					} else {
						break
					}
				}

				//这个中心点可以作为真正的中心点。
				if p < 0 {
					start1 = mid - i/2 + 1
					start2 = mid - i/2
					break
				}

			} else {
				//间隙为中心
				p := mid - i/2 - 1
				q := mid - i/2
				for p >= 0 {
					if s[p] == s[q] {
						p--
						q++
					} else {
						break
					}
				}

				//这个中心点可以作为真正的中心点。
				if p < 0 {
					start1 = mid - i/2
					start2 = mid - i/2
					break
				}
			}
		}
	} else {
		//偶数
		//中心点，可能是数字，可能是间隙。
		for i := 0; i <= N-1; i++ {
			if i%2 == 1 {
				//数字为中心
				p := mid - i/2 - 1
				q := mid - i/2 + 1
				for p >= 0 {
					if s[p] == s[q] {
						p--
						q++
					} else {
						break
					}
				}

				//这个中心点可以作为真正的中心点。
				if p < 0 {
					start1 = mid - i/2 + 1
					start2 = mid - i/2
					break
				}

			} else {
				//间隙为中心
				p := mid - i/2
				q := mid - i/2 + 1
				for p >= 0 {
					if s[p] == s[q] {
						p--
						q++
					} else {
						break
					}
				}

				//这个中心点可以作为真正的中心点。
				if p < 0 {
					start1 = mid - i/2 + 1
					start2 = mid - i/2 + 1
					break
				}
			}
		}
	}

	str1 := s[start1:N]
	N1 := len(str1)
	str2 := s[start2:N]
	N2 := len(str2)
	final := make([]byte, N1+N2)
	for i := 0; i < N1; i++ {
		final[i] = str1[N1-1-i]
	}
	for i := 0; i < N2; i++ {
		final[N1+i] = str2[i]
	}
	return string(final)
}

func main() {

	fmt.Printf("%v \n", shortestPalindrome("a"))
	fmt.Printf("%v \n", shortestPalindrome("abcd"))
	fmt.Printf("%v \n", shortestPalindrome("babcd"))
}
