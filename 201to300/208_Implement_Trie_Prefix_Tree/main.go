package main

import (
	"fmt"
)

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

func main() {
	obj := Constructor()
	obj.Insert("abc")
	param_2 := obj.Search("ab")
	param_3 := obj.StartsWith("ab")
	fmt.Println(param_2)
	fmt.Println(param_3)
}
