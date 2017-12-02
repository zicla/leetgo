package main

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var stack []*TreeNode
	visitedMap := make(map[*TreeNode]bool)

	//根节点入栈
	stack = append(stack, root)
	visitedMap[root] = false
	var visited bool
	for len(stack) != 0 {
		root := stack[0]
		stack = stack[1:]
		visited = visitedMap[root]
		if root == nil {
			continue;
		}
		if visited {
			res = append(res, root.Val)
		} else {

			stack = append([]*TreeNode{root.Right}, stack...)
			stack = append([]*TreeNode{root.Left}, stack...)

			//调整此处的顺序可以获取到前序中序和后续遍历
			stack = append([]*TreeNode{root}, stack...)
			visitedMap[root] = true

		}
	}
	return res
}

func main() {

}
