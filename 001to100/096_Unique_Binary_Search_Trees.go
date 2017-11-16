package main

import "fmt"

func numTrees(n int) int {

	F := make([]int, n+1)

	for i := 1; i <= n; i++ {

		for k := 1; k <= i; k++ {



		}
		//if leftTrees == 0 && rightTrees == 0 {
		//	trees += 1
		//} else if leftTrees == 0 {
		//
		//	trees += rightTrees
		//
		//} else if rightTrees == 0 {
		//	trees += leftTrees
		//} else {
		//	trees += leftTrees * rightTrees
		//}

	}

	return F[n]
}

func numTreesRecursive(arr []int) int {

	var trees int
	for _, v := range arr {
		var newLeftArr []int
		var newRightArr []int
		for _, v1 := range arr {
			if v1 < v {
				newLeftArr = append(newLeftArr, v1)
			} else if v1 > v {
				newRightArr = append(newRightArr, v1)
			}
		}

		var leftTrees int
		if len(newLeftArr) > 0 {
			leftTrees = numTreesRecursive(newLeftArr)
		}
		var rightTrees int
		if len(newRightArr) > 0 {
			rightTrees = numTreesRecursive(newRightArr)
		}

		if leftTrees == 0 && rightTrees == 0 {
			trees += 1
		} else if leftTrees == 0 {

			trees += rightTrees

		} else if rightTrees == 0 {
			trees += leftTrees
		} else {
			trees += leftTrees * rightTrees
		}

	}
	return trees
}

func main() {
	trees := numTrees(19)
	fmt.Printf("%v", trees)
}
