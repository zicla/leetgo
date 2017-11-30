package btree

//参考文章：http://www.cnblogs.com/skywang12345/p/3610187.html
//这里我们只研究最大堆，即根节点大于子节点。索引i的左孩子为 (2*i+1) 右孩子 (2*i+2)。 索引i的父节点 floor((i-1)/2)

//往堆中插入元素
func HeapInsert(arr []int, val int) []int {

	N := len(arr)
	arr = append(arr, val)

	//准备从最后一个元素向上调整。
	//当前调整元素的索引
	start := N
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

	return arr
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

	return arr
}
