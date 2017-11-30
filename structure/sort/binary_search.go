package sort

import "fmt"

//要求传入的arr为从小到达排列。返回所在位置，如果找不到，返回-1
func BinarySearch(arr []int, start int, end int, target int) int {

	if start > end {
		return -1
	}
	standardIndex := start + (end-start)/2
	if arr[standardIndex] == target {
		return standardIndex
	} else if arr[standardIndex] < target {
		return BinarySearch(arr, start+1, end, target)
	} else {
		return BinarySearch(arr, start, end-1, target)
	}
}

func main() {
	fmt.Println("Binary Search")

	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	b := BinarySearch(a, 0, len(a)-1, 3)

	fmt.Println(b)
}
