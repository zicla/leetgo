package main

import (
	"fmt"
)

func isHappy(n int) bool {

	return isHappyRecursive(n, make(map[int]bool))
}

func isHappyRecursive(n int, visited map[int]bool) bool {
	if n == 1 {
		return true
	}
	if visited[n] {
		return false
	}

	visited[n] = true

	res := 0
	for n > 9 {
		res += (n % 10) * (n % 10)
		n = n / 10
	}
	res += n * n

	return isHappyRecursive(res, visited)
}

func main() {

	fmt.Printf("%v \n", isHappy(19))
	fmt.Printf("%v \n", isHappy(5))
	fmt.Printf("%v \n", isHappy(4))
	fmt.Printf("%v \n", isHappy(7))

}
