package main

import (
	"fmt"
	"math"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	N := numCourses
	var res []int
	if N == 0 {
		return res
	} else if N == 1 {
		return []int{0}
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
			return []int{}
		}

		//遍历该元素。
		for _, v := range graph[cur] {
			inDegree[v]--
		}
		delete(inDegree, cur)
		res = append(res, cur)
	}

	return res
}

func main() {
	fmt.Printf("%v \n", findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
