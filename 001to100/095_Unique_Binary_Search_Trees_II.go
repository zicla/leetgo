package main

import "fmt"

type TreeNode095 struct {
	Val   int
	Left  *TreeNode095
	Right *TreeNode095
}

func generateTrees(n int) []*TreeNode095 {

	arr := make([]int, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = i
	}
	//arr := []int{1, 2, 3}
	return generateTreesRecursive(arr)
}

func generateTreesRecursive(arr []int) []*TreeNode095 {

	var trees []*TreeNode095
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

		var leftTrees []*TreeNode095
		if len(newLeftArr) > 0 {
			leftTrees = generateTreesRecursive(newLeftArr)
		}
		var rightTrees []*TreeNode095
		if len(newRightArr) > 0 {
			rightTrees = generateTreesRecursive(newRightArr)
		}

		if len(leftTrees) == 0 && len(rightTrees) == 0 {
			node := &TreeNode095{Val: v}
			trees = append(trees, node)
		} else if len(leftTrees) == 0 {
			for _, rightTree := range rightTrees {
				node := &TreeNode095{Val: v}
				node.Right = rightTree
				trees = append(trees, node)
			}
		} else if len(rightTrees) == 0 {
			for _, leftTree := range leftTrees {
				node := &TreeNode095{Val: v}
				node.Left = leftTree
				trees = append(trees, node)
			}
		} else {
			for _, leftTree := range leftTrees {
				for _, rightTree := range rightTrees {
					node := &TreeNode095{Val: v}
					node.Left = leftTree
					node.Right = rightTree
					trees = append(trees, node)
				}
			}
		}

	}
	return trees

}

func main() {

	trees := generateTrees(3)
	fmt.Printf("%v", trees)
	fmt.Printf("%v", len(trees))
}
