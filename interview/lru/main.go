package main

import (
	"fmt"
)

type Node struct {
	key   string
	value int
	pre   *Node
	next  *Node
}

type MyCache struct {
	capacity int

	//TODO:
	head      *Node
	tail      *Node
	container map[string]*Node
}

//构造函数
func NewMyCache(capacity int) *MyCache {
	//TODO:

	myCache := &MyCache{
		capacity:  capacity,
		head:      &Node{},
		tail:      &Node{},
		container: make(map[string]*Node),
	}

	myCache.head.next = myCache.tail
	myCache.tail.pre = myCache.head

	return myCache
}

func (myCache *MyCache) Get(key string) (val int, ok bool) {
	//TODO:

	if node, ok := myCache.container[key]; ok {

		//将对象移动到队首。
		node.pre.next = node.next
		node.next.pre = node.pre

		node.next = myCache.head.next
		myCache.head.next.pre = node
		node.pre = myCache.head
		myCache.head.next = node

		return node.value, true
	} else {
		return -1, false
	}

}

func (myCache *MyCache) Set(key string, value int) {
	//TODO:

	if len(myCache.container) != 0 && len(myCache.container) >= myCache.capacity {
		//淘汰最后一个
		lastNode := myCache.tail.pre
		myCache.tail.pre = lastNode.pre
		lastNode.pre.next = myCache.tail

		//从container中删除
		delete(myCache.container, lastNode.key)
	}

	node := &Node{
		key:   key,
		value: value,
	}
	//加入到map中
	myCache.container[key] = node

	//放到队首去。
	node.next = myCache.head.next
	myCache.head.next.pre = node
	node.pre = myCache.head
	myCache.head.next = node

}

func main() {

	myCache := NewMyCache(3)

	fmt.Println(myCache.Get("d")) //预期结果：-1 false
	myCache.Set("a", 1)
	myCache.Set("b", 2)
	myCache.Set("c", 3)
	fmt.Println(myCache.Get("a")) //预期结果：1 true
	myCache.Set("d", 4)
	fmt.Println(myCache.Get("b")) //预期结果：-1 false
}
