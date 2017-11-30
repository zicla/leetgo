package sort

//从上往下归并排序。也就是不停分割。
func MergeSort(a []int, start int, end int) {

	if start >= end {
		return
	}

	mid := (start + end) / 2

	MergeSort(a, start, mid)
	MergeSort(a, mid+1, end)

	Merge(a, start, mid, end)

}

//两个区间分别是 [start,mid] [mid+1 end]
func Merge(a []int, start int, mid int, end int) {
	//需要一个临时区域。
	N := end - start + 1
	tmp := make([]int, N)
	p1 := start
	p2 := mid + 1

	p := 0

	for p1 <= mid && p2 <= end {
		if a[p1] < a[p2] {
			tmp[p] = a[p1]
			p++
			p1++
		} else {
			tmp[p] = a[p2]
			p++
			p2++
		}
	}

	for p1 <= mid {
		tmp[p] = a[p1]
		p++
		p1++
	}
	for p2 <= end {
		tmp[p] = a[p2]
		p++
		p2++
	}

	//排序好了之后都整合到a中.
	for k, v := range tmp {
		a[start+k] = v
	}
}
