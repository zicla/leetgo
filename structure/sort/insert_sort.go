package sort

//插入排序,按照从小到大排列。
func InsertSort(v []int) {
	N := len(v)

	//操作N-1次即可。每次都是把v[i]安顿好。
	for i := 1; i < N; i++ {

		var j int
		//为v[i]找到合适的位置。
		for j = i - 1; j >= 0; j-- {
			if v[j] < v[i] {
				break;
			}
		}


		//如果是break出来的，那么a[i]就应该放到a[j]后面。对于本来a[i]就是最大的情况不用动,下面的情况可以处理 j=i-1的情况。
		tmp := v[i]
		for k := i - 1; k > j; k-- {
			v[k+1] = v[k]
		}
		//将a[i]放到正确的位置上
		v[j+1] = tmp

	}
}
