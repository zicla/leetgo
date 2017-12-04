package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {

	N := len(grid)
	if N == 0 {
		return 0
	}
	M := len(grid[0])
	if M == 0 {
		return 0
	}

	visited := make([][]bool, N)
	for k, _ := range visited {
		visited[k] = make([]bool, M)
	}
	res := 0
	//直接深度优先搜索。
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {

			if !visited[i][j] && grid[i][j] == '1' {
				res++
				numIslandsDFS(grid, i, N, j, M, visited)
			}
		}
	}

	return res
}

func numIslandsDFS(grid [][]byte, r int, N int, c int, M int, visited [][]bool) {

	if grid[r][c] == '0' {
		return
	}
	if visited[r][c] {
		return
	}

	visited[r][c] = true

	if r > 0 {
		numIslandsDFS(grid, r-1, N, c, M, visited)
	}
	if c > 0 {
		numIslandsDFS(grid, r, N, c-1, M, visited)
	}

	if r < N-1 {
		numIslandsDFS(grid, r+1, N, c, M, visited)
	}

	if c < M-1 {
		numIslandsDFS(grid, r, N, c+1, M, visited)
	}
}

func main() {

	fmt.Printf("%v \n", numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))

}
