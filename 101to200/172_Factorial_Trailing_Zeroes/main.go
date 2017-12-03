package main

import "fmt"

func trailingZeroes(n int) int {

	res := 0
	for n >= 5 {
		res += n / 5;
		n = n / 5;
	}
	return res
}
func main() {

	fmt.Printf("%v \n", trailingZeroes(5))
	fmt.Printf("%v \n", trailingZeroes(15))
	fmt.Printf("%v \n", trailingZeroes(1808548329))
}
