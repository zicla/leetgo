package btree

type Stack struct {
	N   int
	Arr []*TreeNode
}

func (this *Stack) Push(node *TreeNode) {
	if this.Arr == nil {
		this.Arr = make([]*TreeNode, 1<<7-1)
	}
	this.Arr[this.N] = node
	this.N++
}

func (this *Stack) Pop() *TreeNode {
	if this.N == 0 {
		return nil
	} else {
		node := this.Arr[this.N-1]
		this.N--
		return node
	}
}

func (this *Stack) Empty() bool {
	return this.N == 0
}
