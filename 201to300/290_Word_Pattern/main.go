package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	dict := make(map[int32]string)
	dictReverse := make(map[string]int32)
	arr := strings.Split(str, " ")
	if len(arr) != len(pattern) {
		return false
	}
	for k, v := range pattern {
		if _, ok := dict[v]; ok {
			if dict[v] != arr[k] {
				return false
			}
		} else {

			if _, ok := dictReverse[arr[k]]; ok {
				return false
			} else {
				dict[v] = arr[k]
				dictReverse[arr[k]] = v
			}

		}
	}
	return true
}

func main() {

	fmt.Printf("%v\n", wordPattern("abba", "dog cat cat dog"))
	fmt.Printf("%v\n", wordPattern("abba", "dog cat cat fish"))
	fmt.Printf("%v\n", wordPattern("aaaa", "dog cat cat dog"))
	fmt.Printf("%v\n", wordPattern("abba", "dog dog dog dog"))
}
