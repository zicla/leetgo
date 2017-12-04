package main

import (
	"fmt"
	. "leetgo/structure/btree"
)

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	//典型的按层遍历，使用一个队列即可。
	var queue []*TreeNode

	//入队,nil用来分割每一层的
	queue = append(queue, root)
	queue = append(queue, nil)

	//记录每一层的最后一个元素
	var res []int
	var level = 1
	for true {
		//出队
		cur := queue[0]
		queue = queue[1:]

		if cur != nil {

			if len(res) == level {
				res[level-1] = cur.Val
			} else {
				res = append(res, cur.Val)
			}

			//装入下一层
			if cur.Left != nil {
				//入队
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				//入队
				queue = append(queue, cur.Right)
			}
		} else {
			//队列此时已经空了，那么就没了。
			if len(queue) == 0 {
				break
			} else {
				//nil入队，这里表示新增加了一层
				queue = append(queue, nil)
				level++
			}

		}

	}
	return res

}
func main() {

	fmt.Printf("%v \n", rightSideView(CreateBinaryTree([]int{2, 5, 8, 1, 3}, 0)))

}
