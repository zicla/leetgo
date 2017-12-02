package main

import (
	"fmt"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	//队列
	var queue []string

	//dict标明第几层可以到达beginWord
	var dict = make(map[string]int)
	for _, v := range wordList {
		dict[v] = -1
	}

	queue = append(queue, beginWord)
	//当前所在层数
	level := 1

	oldLevelItems := 1
	newLevelItems := 0

	dict[beginWord] = 1

	//最终存储结果
	var res [][]string
	//最终层数应该是多少层。
	finalLevel := 0

	//记录自己前面是谁，告诉我们不能忘本啊。
	preNodeMap := make(map[string]map[string]bool)

	for len(queue) > 0 {
		str := queue[0]
		queue = queue[1:]
		oldLevelItems--

		if finalLevel == level {
			break
		}

		for i := 0; i < len(str); i++ {

			chs := []rune(str)
			for c := 'a'; c <= 'z'; c++ {
				chs[i] = c
				newStr := string(chs)

				if newStr == endWord && (dict[newStr] == level || dict[newStr] == -1) {
					dict[newStr] = level

					finalLevel = level + 1

					if preNodeMap[newStr] == nil {
						preNodeMap[newStr] = make(map[string]bool)
					}
					preNodeMap[newStr][str] = true

				} else {
					if dict[newStr] == level || dict[newStr] == -1 {
						queue = append(queue, newStr)
						dict[newStr] = level
						newLevelItems++

						if preNodeMap[newStr] == nil {
							preNodeMap[newStr] = make(map[string]bool)
						}
						preNodeMap[newStr][str] = true

					}
				}

			}
		}

		if oldLevelItems == 0 {
			oldLevelItems = newLevelItems
			newLevelItems = 0
			level++
		}

	}

	if len(preNodeMap[endWord]) > 0 {

		path := make([]string, finalLevel)
		//开始一次深度优先搜索。
		path[finalLevel-1] = endWord
		findLaddersRecursive(preNodeMap, &res, endWord, beginWord, finalLevel-2, path)
	}

	return res
}
func findLaddersRecursive(nodeMap map[string]map[string]bool, res *[][]string, cur string, beginWord string, index int, path []string) {
	if index == -1 {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	strs := nodeMap[cur]
	for k, _ := range strs {

		path[index] = k
		findLaddersRecursive(nodeMap, res, k, beginWord, index-1, path)
	}
}

func main() {

	fmt.Printf("%v\n", findLadders("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}))
	fmt.Printf("%v\n", findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
	fmt.Printf("%v\n", findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))

}
