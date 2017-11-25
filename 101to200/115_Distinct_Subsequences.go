package main

import (
	"fmt"
)

func numDistinct(s string, t string) int {

	total := 0
	numDistinctRecursive(s, len(s), 0, t, len(t), 0, &total)
	return total
}

func numDistinctRecursive(s string, sLen int, sIndex int, t string, tLen int, tIndex int, total *int) {

	if tIndex == tLen {
		*total++
		return
	}
	if sIndex == sLen {
		return
	}

	if s[sIndex] == t[tIndex] {
		//s继续前进，t按兵不动
		numDistinctRecursive(s, sLen, sIndex+1, t, tLen, tIndex, total)
		//s和t同时前进
		numDistinctRecursive(s, sLen, sIndex+1, t, tLen, tIndex+1, total)
	} else {
		//s继续前进，t按兵不动
		numDistinctRecursive(s, sLen, sIndex+1, t, tLen, tIndex, total)
	}
}

func main() {

	fmt.Println(numDistinct("bbbi", "bi"))
	fmt.Println(numDistinct("rabbbit", "rabbit"))

}
