package btree

//参考文章：http://www.cnblogs.com/skywang12345/p/3610187.html
//这里我们只研究最大堆，即根节点大于子节点。索引i的左孩子为 (2*i+1) 右孩子 (2*i+2)。 索引i的父节点 floor((i-1)/2)

//往堆中插入元素
func HeapInsert(arr []int, val int) []int {

	arr = append(arr, val)
	N := len(arr)

	HeapUp(arr, N)
	return arr
}

//向上整理。前N个元素
func HeapUp(arr []int, N int) {
	//准备从最后一个元素向上调整。
	//当前调整元素的索引
	start := N - 1
	//父节点的索引
	parent := (start - 1) / 2
	//当前节点的大小
	tmp := arr[start]

	for start > 0 {
		if arr[parent] >= tmp {
			break
		} else {
			arr[start] = arr[parent]
			start = parent
			parent = (start - 1) / 2
		}
	}

	arr[start] = tmp
}

//从堆中删除元素
func HeapRemove(arr []int, val int) []int {

	N := len(arr)
	if N == 0 {
		return arr
	}

	//获取val的索引。
	index := -1
	for i := 0; i < N; i++ {
		if val == arr[i] {
			index = i;
			break
		}
	}
	if index == -1 {
		return arr
	}

	//用最后的元素填补删掉的元素
	arr[index] = arr[N-1]
	arr = arr[:N-1]
	N = N - 1

	HeapDown(arr, N, index)

	return arr
}

//向下整理前N个元素,从index这个元素开始。
func HeapDown(arr []int, N int, index int) {

	//从index开始向下调整。
	c := index
	left := 2*c + 1
	tmp := arr[c]
	//最后一个元素的索引。
	for left < N {
		//left左孩子，left+1右孩子
		if left+1 < N && arr[left] < arr[left+1] {
			//左右孩子选择比较大的一个。
			left++
		}

		if tmp >= arr[left] {
			//调整结束
			break
		} else {
			arr[c] = arr[left]
			c = left
			left = 2*c + 1
		}
	}
	arr[c] = tmp
}

//将一个乱序的数组，整理成一个二叉堆。
func HeapCreate(arr []int) {
	N := len(arr)
	//最笨的办法，从第一个元素开始插入到最后一个元素。
	for i := 1; i <= N; i++ {
		HeapUp(arr, i)
	}
}

//将一个数组进行堆排序。
func HeapSort(arr []int) {
	N := len(arr)

	//先搞一个堆出来。
	HeapCreate(arr)

	//都和最后一个元素交换，交换后调整
	for i := N - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		HeapDown(arr, i, 0)
	}

}
