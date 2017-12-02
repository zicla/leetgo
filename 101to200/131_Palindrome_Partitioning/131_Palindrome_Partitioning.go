package main

import "fmt"

func partition(s string) [][]string {

	var res [][]string;
	N := len(s)
	if N != 0 {
		var path = make([]string, N)
		partitionRecursive(s, N, 0, path, 0, &res)
	}

	return res
}

func partitionIsPalindrome(s string) bool {

	N := len(s)
	if N == 0 || N == 1 {
		return true
	}
	p := 0
	q := N - 1
	for p < q {
		if s[p] != s[q] {
			return false
		}
		p++
		q--
	}
	return true
}

func partitionRecursive(s string, N int, already int, path []string, pathLength int, res *[][]string) {

	if already == N {
		//开始输出。
		tmp := make([]string, pathLength)
		copy(tmp, path[:pathLength])
		*res = append(*res, tmp)
		return
	}

	for i := already; i < N; i++ {
		tmpStr := s[already:i+1]
		if partitionIsPalindrome(tmpStr) {
			path[pathLength] = tmpStr
			partitionRecursive(s, N, already+len(tmpStr), path, pathLength+1, res)
		}
	}
}

func main() {

	fmt.Printf("%v\n", partition("aab"))
}
