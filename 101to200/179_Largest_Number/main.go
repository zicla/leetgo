package main

import (
	"fmt"
	"strconv"
)

func greater(a, b int) int {

	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)
	strab := aStr + bStr
	strba := bStr + aStr
	if strab > strba {
		return 1
	} else if strab < strba {
		return -1
	} else {
		return 0
	}
}

//start:开始的index end:结束的index
func QuickSort(v []int, left int, right int) {

	if left < right {
		key := v[left];
		low := left;
		high := right;
		for low < high {
			for low < high && greater(key, v[high]) >= 0 {
				high--;
			}
			v[low] = v[high];
			for low < high && greater(v[low], key) >= 0 {
				low++;
			}
			v[high] = v[low];
		}
		v[low] = key;
		QuickSort(v, left, low-1);
		QuickSort(v, low+1, right);
	}


}

func largestNumber(nums []int) string {
	N := len(nums)
	QuickSort(nums, 0, N-1)
	if nums[0] == 0 {
		return "0"
	}
	str := ""
	for i := 0; i < N; i++ {
		str = fmt.Sprintf("%s%d", str, nums[i])
	}
	return str
}
func main() {

	fmt.Printf("%v \n", largestNumber([]int{1, 1, 1}))
	fmt.Printf("%v \n", largestNumber([]int{0, 0}))
	fmt.Printf("%v \n", largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Printf("%v \n", largestNumber([]int{121, 12}))
}
