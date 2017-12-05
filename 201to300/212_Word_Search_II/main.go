package main

import "fmt"

type Trie struct {
	word  bool
	nodes []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{nodes: make([]*Trie, 26)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if word == "" {
		this.word = true
		return
	}

	ch := word[0]
	index := ch - 'a'
	if this.nodes[index] == nil {
		this.nodes[index] = &Trie{nodes: make([]*Trie, 26)}
	}
	this.nodes[index].Insert(word[1:])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if word == "" {
		return this.word
	}

	ch := word[0]
	index := ch - 'a'
	if this.nodes[index] == nil {
		return false
	}

	return this.nodes[index].Search(word[1:])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return true
	}

	ch := prefix[0]
	index := ch - 'a'
	if this.nodes[index] == nil {
		return false
	}

	return this.nodes[index].StartsWith(prefix[1:])
}

func findWords(board [][]byte, words []string) []string {

	//处理边界条件
	R := len(board)
	M := len(words)
	if R == 0 || M == 0 {
		return []string{}
	}
	C := len(board[0])
	if C == 0 {
		return []string{}
	}

	trieI := Constructor()
	trie := &trieI
	for _, v := range words {
		trie.Insert(v)
	}

	visited := make([][]bool, R)
	for k, _ := range visited {
		visited[k] = make([]bool, C)
	}

	path := make([]byte, R*C)
	res := make(map[string]bool)

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			visited[i][j] = true
			findWordsRecursive(board, visited, R, C, words, trie, path, 0, i, j, res)
			visited[i][j] = false
		}
	}

	var strs []string
	for k, _ := range res {
		strs = append(strs, k)
	}

	return strs
}

//从i,j点开始，可以延伸出来的单词，此处保证 [i,j] 没有访问过，保证[i,j]不越界
func findWordsRecursive(board [][]byte, visited [][]bool, R int, C int, words []string, trie *Trie, path []byte, length int, i int, j int, res map[string]bool) {

	path[length] = board[i][j]
	length++

	word := string(path[:length])
	if trie.Search(word) {
		res[word] = true
	}
	if !trie.StartsWith(word) {
		return
	}

	if i > 0 {
		if !visited[i-1][j] {
			visited[i-1][j] = true
			findWordsRecursive(board, visited, R, C, words, trie, path, length, i-1, j, res)

			visited[i-1][j] = false
		}
	}
	if j > 0 {
		if !visited[i][j-1] {
			visited[i][j-1] = true
			findWordsRecursive(board, visited, R, C, words, trie, path, length, i, j-1, res)

			visited[i][j-1] = false
		}

	}
	if i < R-1 {
		if !visited[i+1][j] {
			visited[i+1][j] = true
			findWordsRecursive(board, visited, R, C, words, trie, path, length, i+1, j, res)

			visited[i+1][j] = false
		}
	}
	if j < C-1 {
		if !visited[i][j+1] {
			visited[i][j+1] = true
			findWordsRecursive(board, visited, R, C, words, trie, path, length, i, j+1, res)

			visited[i][j+1] = false
		}
	}

}

func main() {

	fmt.Printf("%v \n", findWords([][]byte{{'a', 'a'}}, []string{"aa"}))
	fmt.Printf("%v \n", findWords([][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"eat", "oath"}))
}
