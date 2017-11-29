package main

import "fmt"

//我们使用一个双向链表来存储key-value.
type BiNode struct {
	Key   int
	Value int
	Pre   *BiNode
	Next  *BiNode
}

type LRUCache struct {
	Size     int
	Capacity int
	Head     *BiNode
	Tail     *BiNode
	//key -> BiNode
	Dict map[int]*BiNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Dict:     make(map[int]*BiNode),
	}
}

func (this *LRUCache) Get(key int) int {
	node := this.Dict[key]
	if node == nil {
		return -1
	}

	if node != this.Head {
		if node == this.Tail {
			this.Tail = this.Tail.Pre
			this.Tail.Next = nil
			node.Pre = nil
			node.Next = this.Head
			this.Head.Pre = node
			this.Head = node

		} else {
			node.Pre.Next = node.Next
			node.Next.Pre = node.Pre

			node.Pre = nil
			node.Next = this.Head
			this.Head.Pre = node
			this.Head = node
		}
	}

	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	if this.Capacity == 0 {
		return
	}

	node := this.Dict[key]
	if node == nil {
		//这时候需要删除优先级最低的。
		if this.Size == this.Capacity {

			deleteKey := this.Tail.Key

			if this.Size == 1 {
				this.Tail = nil
			} else {
				this.Tail = this.Tail.Pre
				this.Tail.Next = nil
			}

			delete(this.Dict, deleteKey)
			this.Size--
		}

		tmp := &BiNode{
			Key:   key,
			Value: value,
			Next:  this.Head,
		}
		this.Dict[key] = tmp

		this.Size++
		if this.Size == 1 {
			this.Tail = tmp
		}

		if this.Head != nil {
			this.Head.Pre = tmp
		}
		this.Head = tmp
	} else {
		node.Value = value

		//再调整优先级即可。
		if node != this.Head {
			if node == this.Tail {
				this.Tail = this.Tail.Pre
				this.Tail.Next = nil
				node.Pre = nil
				node.Next = this.Head
				this.Head.Pre = node
				this.Head = node

			} else {
				node.Pre.Next = node.Next
				node.Next.Pre = node.Pre

				node.Pre = nil
				node.Next = this.Head
				this.Head.Pre = node
				this.Head = node
			}
		}
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	var a int
	cache := Constructor(2)
	a = cache.Get(2);
	fmt.Println(a)
	cache.Put(2, 6);
	a = cache.Get(1);
	fmt.Println(a)
	cache.Put(1, 5);
	cache.Put(1, 2);
	a = cache.Get(1);
	fmt.Println(a)
	a = cache.Get(2);
	fmt.Println(a)
}
