package sort

//start:开始的index end:结束的index
func QuickSort(v []int, left int, right int) {

	if left < right {
		key := v[left]
		low := left
		high := right
		for low < high {
			for low < high && v[high] >= key {
				high--
			}
			v[low] = v[high]
			for low < high && v[low] <= key {
				low++
			}
			v[high] = v[low]
		}
		v[low] = key
		QuickSort(v, left, low-1)
		QuickSort(v, low+1, right)
	}

}
