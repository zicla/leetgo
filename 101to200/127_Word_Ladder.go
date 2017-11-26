package main

import (
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {

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
	}

	//从beginWordIndex到endWordIndex的最短距离
	N := len(wordList)

	return beginWordIndex + endWordIndex + N
}

func main() {

	fmt.Printf("%v\n", ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))

}
