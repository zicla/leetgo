package main

import "fmt"

func solve(board [][]byte) {
	//和墙壁联通的O都有活路，剩下的都是死路。有活路的O先标成P，最后O标成X,P标成O.

	N := len(board)
	if N == 0 {
		return
	}
	if len(board[0]) == 0 {
		return
	}

	for i := 0; i < N; i++ {
		if i == 0 || i == N-1 {
			for j := 0; j < len(board[0]); j++ {
				solveRecursive(board, i, j, N, len(board[0]))
			}
		} else {
			solveRecursive(board, i, 0, N, len(board[0]))
			solveRecursive(board, i, len(board[0])-1, N, len(board[0]))
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {

			if board[i][j] == 'P' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}

//这里是把活着的O标成P
func solveRecursive(board [][]byte, i int, j int, I int, J int) {
	if board[i][j] == 'X' || board[i][j] == 'P' {
		return
	} else {
		board[i][j] = 'P'

		if i > 0 {
			solveRecursive(board, i-1, j, I, J)
		}
		if i < I-1 {
			solveRecursive(board, i+1, j, I, J)
		}

		if j > 0 {
			solveRecursive(board, i, j-1, I, J)
		}
		if j < J-1 {
			solveRecursive(board, i, j+1, I, J)
		}

	}
}

func main() {

	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(board);
	for i := 0; i < len(board); i++ {
		if len(board[0]) > 0 {
			for j := 0; j < len(board[0]); j++ {
				fmt.Printf("%c ", board[i][j])
			}
		}
		fmt.Println()
	}
}
