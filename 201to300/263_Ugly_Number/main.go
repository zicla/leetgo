package main

import (
	"fmt"
)

func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 {
		return true
	}

	if num%2 == 0 {
		return isUgly(num / 2)
	} else if num%3 == 0 {
		return isUgly(num / 3)
	} else if num%5 == 0 {
		return isUgly(num / 5)
	} else {
		return false
	}
}
func main() {

	fmt.Printf("%v\n", isUgly(14))
	fmt.Printf("%v\n", isUgly(6))
	fmt.Printf("%v\n", isUgly(8))
}
