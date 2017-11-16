package main

import "fmt"

func numTrees(n int) int {

	F := make([]int, n+1)

	for i := 1; i <= n; i++ {
		for k := 1; k <= i; k++ {
			leftNodes := k - 1
			rightNodes := i - 1 - (k - 1)
			if leftNodes == 0 && rightNodes == 0 {
				F[i] += 1
			} else if leftNodes == 0 {
				F[i] += F[rightNodes]
			} else if rightNodes == 0 {
				F[i] += F[leftNodes]
			} else {
				F[i] += F[leftNodes] * F[rightNodes]
			}
		}
	}
	return F[n]
}

func main() {
	trees := numTrees(19)
	fmt.Printf("%v", trees)
}
