package structure

type Stack struct {
	N   int
	Arr []int
}

func (this *Stack) push(node int) {
	if this.Arr == nil {
		this.Arr = make([]int, 1<<7-1)
	}
	this.Arr[this.N] = node
	this.N++
}

func (this *Stack) pop() int {
	if this.N == 0 {
		return 0
	} else {
		return this.Arr[this.N]
	}
}

func (this *Stack) empty() bool {
	return this.N == 0
}
