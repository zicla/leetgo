package main

import (
	"fmt"
	"math"
)

func Dijkstra(matrix [][]int, vs int) []int {
	const INF = math.MaxInt32

	vexnum := len(matrix)
	dist := make([]int, vexnum)
	// flag[i]=true表示"顶点vs"到"顶点i"的最短路径已成功获取。用此变量可以区分S和U集合。
	flag := make([]bool, vexnum)

	// 初始化
	for i := 0; i < vexnum; i++ {
		// 顶点i的最短路径还没获取到。
		flag[i] = false

		// 顶点i的最短路径为"顶点vs"到"顶点i"的权。
		dist[i] = matrix[vs][i];
	}

	// 对"顶点vs"自身进行初始化
	flag[vs] = true
	dist[vs] = 0

	// 遍历G.vexnum-1次；每次找出一个顶点的最短路径。
	for i := 1; i < vexnum; i++ {
		// 寻找当前最小的路径；
		// 即，在未获取最短路径的顶点中，找到离vs最近的顶点 k。
		min := INF
		var k int
		for j := 0; j < vexnum; j++ {
			if flag[j] == false {
				if dist[j] < min {
					min = dist[j]
					k = j
				}
			}
		}
		// 标记"顶点k"为已经获取到最短路径
		flag[k] = true
		// 修正当前最短路径和前驱顶点
		// 即，当已经"顶点k的最短路径"之后，更新"未获取最短路径的顶点的最短路径和前驱顶点"
		for j := 0; j < vexnum; j++ {
			var tmp = INF
			if matrix[k][j] != INF {
				tmp = min + matrix[k][j]
			}
			if flag[j] == false {
				if tmp < dist[j] {
					dist[j] = tmp
				}
			}
		}
	}

	return dist
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

	const INF = math.MaxInt32

	//准备创建一个图出来。
	beginWordAppear := false
	endWordAppear := false
	beginWordIndex := 0
	endWordIndex := 0
	for i, v := range wordList {
		if v == beginWord {
			beginWordAppear = true
			beginWordIndex = i
		}
		if v == endWord {
			endWordAppear = true
			endWordIndex = i
		}
	}
	if !beginWordAppear {
		wordList = append(wordList, beginWord)
		beginWordIndex = len(wordList) - 1
	}
	if !endWordAppear {
		wordList = append(wordList, endWord)
		endWordIndex = len(wordList) - 1

		return 0
	}

	var distance = func(a string, b string) int {
		notSame := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				notSame++
			}
		}
		if notSame > 1 {
			return INF
		} else {
			return 1
		}
	}

	//从beginWordIndex到endWordIndex的最短距离
	N := len(wordList)
	matrix := make([][]int, N)
	for i := 0; i < N; i++ {
		matrix[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			if i == j {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = distance(wordList[i], wordList[j])
				matrix[j][i] = matrix[i][j]
			}
		}
	}


	//floyed开始你的表演 实时证明floyed会超时。
	//for k := 0; k < N; k++ {
	//	for i := 0; i < N; i++ {
	//		for j := 0; j < N; j++ {
	//			if matrix[i][j] > matrix[i][k]+matrix[k][j] {
	//				matrix[i][j] = matrix[i][k] + matrix[k][j];
	//			}
	//		}
	//	}
	//}
	//if matrix[beginWordIndex][endWordIndex] == INF {
	//	return 0
	//} else {
	//	return matrix[beginWordIndex][endWordIndex] + 1
	//}

	//Dijkstra开始你的表演
	list := Dijkstra(matrix, beginWordIndex)
	if list[endWordIndex] == INF {
		return 0
	} else {
		return list[endWordIndex] + 1
	}

}

func main() {

	fmt.Printf("%v\n", ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
	fmt.Printf("%v\n", ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))

}
