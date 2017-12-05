package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {

	N := numCourses
	if N == 0 || N == 1 {
		return true
	}

	matrix := make([][]bool, N)
	visited := make([]int, N)
	for k, _ := range matrix {
		matrix[k] = make([]bool, N)
	}

	for _, v := range prerequisites {
		matrix[v[1]][v[0]] = true
	}

	for i := 0; i < N; i++ {
		b := canFinishDFS(matrix, N, visited, i)
		if !b {
			return false
		}
	}

	return true

}

//visited 0表示没有访问过，1表示访问过了，并且没问题。 -1 表示临时给这条路加个锁。
func canFinishDFS(matrix [][]bool, N int, visited []int, index int) bool {

	if visited[index] == 1 {
		return true
	}
	if visited[index] == -1 {
		return false
	}
	visited[index] = -1

	res := true
	for i := 0; i < N; i++ {
		if matrix[index][i] {
			res = canFinishDFS(matrix, N, visited, i)
			if !res {
				return false
			}
		}
	}

	visited[index] = 1

	return res

}

func main() {

	fmt.Printf("%v \n", canFinish(3, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	fmt.Printf("%v \n", canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Printf("%v \n", canFinish(3, [][]int{{2, 0}, {2, 1}}))
	fmt.Printf("%v \n", canFinish(2, [][]int{{0, 1}}))
	fmt.Printf("%v \n", canFinish(2, [][]int{{1, 0}}))
	fmt.Printf("%v \n", canFinish(4, [][]int{{1, 0}, {2, 1}, {3, 2}}))
	fmt.Printf("%v \n", canFinish(4, [][]int{{1, 0}, {2, 1}, {3, 2}, {1, 3}}))
}
