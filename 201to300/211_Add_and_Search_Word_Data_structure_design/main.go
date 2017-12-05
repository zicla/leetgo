package main

import (
	"fmt"
)

type WordDictionary struct {
	word  bool
	nodes []*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{nodes: make([]*WordDictionary, 26)}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	if word == "" {
		this.word = true
		return
	}

	ch := word[0]
	index := ch - 'a'
	if this.nodes[index] == nil {
		this.nodes[index] = &WordDictionary{nodes: make([]*WordDictionary, 26)}
	}
	this.nodes[index].AddWord(word[1:])
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if word == "" {
		return this.word
	}

	ch := word[0]

	if ch == '.' {

		for i := 0; i < 26; i++ {
			if this.nodes[i] != nil {
				b := this.nodes[i].Search(word[1:])
				if b {
					return true
				}
			}
		}
		return false

	} else {
		index := ch - 'a'
		if this.nodes[index] == nil {
			return false
		}

		return this.nodes[index].Search(word[1:])
	}

}

func main() {
	obj := Constructor()
	obj.AddWord("abc")
	param_2 := obj.Search("ab")
	param_3 := obj.Search("ab.")
	param_4 := obj.Search("ab.c")
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
}
