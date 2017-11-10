package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	p1 := m - 1
	p2 := n - 1
	for i := m + n - 1; i >= 0; i-- {

		if p2 < 0 {
			return
		}

		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[i] = nums1[p1]
			p1--
		} else {
			nums1[i] = nums2[p2]
			p2--
		}
	}
}

func main() {

	a := []int{1, 0}
	m := 1
	b := []int{2}
	n := 1
	merge(a, m, b, n)
	fmt.Printf("%v\n", a)


	a = []int{2, 0}
	m = 1
	b = []int{1}
	n = 1
	merge(a, m, b, n)
	fmt.Printf("%v\n", a)



}
