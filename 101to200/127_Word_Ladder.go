package main

import (
	"fmt"
	"math"
)

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

	//可以看作是走迷宫的题目，使用深度优先搜索可以快速获取最短路径，使用广度优先搜索可以获取具体每一条路的策略。
	//d表示到每个其他节点的距离。
	d := make([]int, N)
	for i := 0; i < N; i++ {
		d[i] = INF
	}
	visited := make([]bool, N)

	d[beginWordIndex] = 0
	visited[beginWordIndex] = true

	curPath := 0

	for i := 0; i <= N; i++ {
		if d[i] == curPath {
			i := beginWordIndex
			for j := 0; j < N; j++ {
				if !visited[j] {
					visited[j] = true
					if matrix[i][j] == 1 {
						d[j] = curPath + 1
					}
				}
			}
		}
	}

	return endWordIndex;
}

func main() {

	fmt.Printf("%v\n", ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
	fmt.Printf("%v\n", ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))

}
