package main

import "fmt"

func isScramble(s1 string, s2 string) bool {

	return true
}

func main() {

	var arr [][]int
	for i := 0; i < 10; i++ {
		subArr1 := make([]int, 0, 10)
		for j := 0; j < 5; j++ {
			subArr1 = append(subArr1, j)
		}
		arr = append(arr, subArr1)
	}

	fmt.Printf("%v", arr)
}
