package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {

	N := numCourses
	if N == 0 || N == 1 {
		return true
	}

	matrix := make([][]bool, N)
	visited := make([]bool, N)
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

func canFinishDFS(matrix [][]bool, N int, visited []bool, index int) bool {

	if visited[index] {
		return false
	}
	visited[index] = true

	res := true
	for i := 0; i < N; i++ {
		if index == i {
			continue
		}
		if matrix[index][i] {
			res = canFinishDFS(matrix, N, visited, i)
			if !res {
				break
			}
		}
	}

	visited[index] = false

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
