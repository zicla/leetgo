package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	m := len(nums1)
	n := len(nums2)

	if (m+n)%2 == 0 {
		v1 := findKth(nums1, 0, m, nums2, 0, n, (m+n)/2)
		v2 := findKth(nums1, 0, m, nums2, 0, n, (m+n)/2+1)
		return float64(v1+v2) / 2.0
	} else {
		return float64(findKth(nums1, 0, m, nums2, 0, n, (m+n)/2+1))
	}

}

func min(a int, b int) int {

	if a < b {
		return a
	} else {
		return b
	}

}

//m is the length of array1.
func findKth(nums1 []int, startIndex1 int, m int, nums2 []int, startIndex2 int, n int, k int) int {

	//assume m < n.
	if m > n {
		return findKth(nums2, startIndex2, n, nums1, startIndex1, m, k)
	}

	if m == 0 {
		return nums2[startIndex2+k-1]
	}

	if k == 1 {
		return min(nums1[startIndex1], nums2[startIndex2])
	}

	pa := min(k/2, m)
	pb := k - pa

	pivot1 := nums1[startIndex1+pa-1]
	pivot2 := nums2[startIndex2+pb-1]
	if pivot1 < pivot2 {
		return findKth(nums1, startIndex1+pa, m-pa, nums2, startIndex2, n, k-pa)
	} else if pivot1 > pivot2 {
		return findKth(nums1, startIndex1, m, nums2, startIndex2+pb, n-pb, k-pb)
	} else {
		return pivot1
	}

	return 0

}
func main() {

	num1 := []int{1, 2}
	num2 := []int{3, 4}

	fmt.Println(findMedianSortedArrays(num1, num2))

	num3 := []int{1, 3, 8}
	num4 := []int{8}

	fmt.Println(findMedianSortedArrays(num3, num4))

}
