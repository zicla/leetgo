package main

import "fmt"


//更加快捷的算法是从右上角开始找：从右上角开始, 比较target 和 matrix[i][j]的值. 如果小于target, 则该行不可能有此数,  所以i++; 如果大于target, 则该列不可能有此数, 所以j--. 遇到边界则表明该矩阵不含target.
//参考链接：http://blog.csdn.net/xudli/article/details/47015825
func searchMatrix(matrix [][]int, target int) bool {

	M := len(matrix)
	if M == 0 {
		return false
	}
	N := len(matrix[0])
	if N == 0 {
		return false
	}

	visited := make([][]bool, M)
	for k, _ := range visited {
		visited[k] = make([]bool, N)
	}

	return searchMatrixRecursive(matrix, M, N, 0, 0, visited, target)

}

func searchMatrixRecursive(matrix [][]int, M int, N int, m int, n int, visited [][]bool, target int) bool {

	if visited[m][n] {
		return false
	}
	visited[m][n] = true

	if target == matrix[m][n] {
		return true
	} else if target < matrix[m][n] {

		//向上
		if m > 0 {
			b := searchMatrixRecursive(matrix, M, N, m-1, n, visited, target)
			if b {
				return true
			}
		}

		//向左
		if n > 0 {
			b := searchMatrixRecursive(matrix, M, N, m, n-1, visited, target)
			if b {
				return true
			}
		}

	} else {
		//向下
		if m < M-1 {
			b := searchMatrixRecursive(matrix, M, N, m+1, n, visited, target)
			if b {
				return true
			}
		}

		//向右
		if n < N-1 {
			b := searchMatrixRecursive(matrix, M, N, m, n+1, visited, target)
			if b {
				return true
			}
		}
	}

	return false
}

func main() {

	matrix := [][]int{{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30}}
	fmt.Printf("%v\n", searchMatrix(matrix, 5))
	fmt.Printf("%v\n", searchMatrix(matrix, 20))
}
