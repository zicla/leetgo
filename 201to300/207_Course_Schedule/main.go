package main

import (
	"fmt"
	"math"
)

//这里是深度优先的算法
func canFinish1(numCourses int, prerequisites [][]int) bool {

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

//这里采用广度优先的算法。
func canFinish(numCourses int, prerequisites [][]int) bool {

	N := numCourses
	if N == 0 || N == 1 {
		return true
	}

	graph := make(map[int][]int)
	//记录每个元素的入度
	inDegree := make(map[int]int, N)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
		inDegree[i] = 0
	}

	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
		inDegree[v[0]]++
	}

	for len(inDegree) > 0 {
		//入度为0的节点。
		cur := math.MaxInt64
		for k, v := range inDegree {
			if v == 0 {
				cur = k
				break
			}
		}
		//还有元素，但是已经找不到入度为0的了。
		if cur == math.MaxInt64 {
			return false
		}

		//遍历该元素。
		for _, v := range graph[cur] {
			inDegree[v]--
		}
		delete(inDegree, cur)

	}

	return true

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
