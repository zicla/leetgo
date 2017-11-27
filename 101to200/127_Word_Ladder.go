package main

import (
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {

	//队列
	var queue []string

	var dict = make(map[string]bool)
	for _, v := range wordList {
		dict[v] = true
	}

	queue = append(queue, beginWord)
	//当前所在层数
	level := 1
	oldLevelItems := 1
	newLevelItems := 0

	dict[beginWord] = false

	for len(queue) > 0 {
		str := queue[0]
		queue = queue[1:]
		oldLevelItems--

		for i := 0; i < len(str); i++ {

			chs := []rune(str)
			for c := 'a'; c <= 'z'; c++ {
				chs[i] = c
				newStr := string(chs)
				if newStr == endWord && dict[newStr] {
					return level + 1
				}
				if dict[newStr] {
					queue = append(queue, newStr)
					dict[newStr] = false
					newLevelItems++
				}
			}
		}

		if oldLevelItems == 0 {
			oldLevelItems = newLevelItems
			newLevelItems = 0
			level++
		}

	}

	return 0;
}

func main() {

	fmt.Printf("%v\n", ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
	fmt.Printf("%v\n", ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))

}
