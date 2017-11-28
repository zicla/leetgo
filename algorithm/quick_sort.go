package main

import "fmt"

//start:开始的index end:结束的index
func QuickSort(v []int, left int, right int) {

	if left < right {
		key := v[left];
		low := left;
		high := right;
		for low < high {
			for low < high && v[high] > key {
				high--;
			}
			v[low] = v[high];
			for low < high && v[low] < key {
				low++;
			}
			v[high] = v[low];
		}
		v[low] = key;
		QuickSort(v, left, low-1);
		QuickSort(v, low+1, right);
	}

}

func main() {
	fmt.Println("Quick Sort")

	a := []int{5, 4, 2, 1, 3, 100, 0, 200, 500}
	QuickSort(a, 0, len(a)-1)

	fmt.Printf("%v", a)
}
