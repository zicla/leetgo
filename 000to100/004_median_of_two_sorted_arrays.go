package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	m := len(nums1)
	n := len(nums2)

	//odd
	p1 := 0
	p2 := 0

	count := 1

	if (m+n)%2 == 1 {
		target := (m + n + 1) / 2
		for true {

			var current int;
			//nums1 end.
			if p1 >= m {
				current = nums2[p2]
				p2++
			} else if p2 >= n {
				current = nums1[p1]
				p1++
			} else {
				if nums1[p1] > nums2[p2] {
					current = nums2[p2]
					p2++
				} else {
					current = nums1[p1]
					p1++
				}
			}

			if count == target {
				return float64(current)
			}

			count++
		}

	} else {
		target := (m + n) / 2
		var seed1 int
		var seed2 int
		//even
		for true {

			var current int;
			//nums1 end.
			if p1 >= m {
				current = nums2[p2]
				p2++
			} else if p2 >= n {
				current = nums1[p1]
				p1++
			} else {
				if nums1[p1] > nums2[p2] {
					current = nums2[p2]
					p2++
				} else {
					current = nums1[p1]
					p1++
				}
			}

			if count == target {
				seed1 = current
			} else if count == target+1 {
				seed2 = current
				break
			}

			count++
		}
		return (float64(seed1) + float64(seed2)) / 2
	}

	return 0;
}
func main() {

	num1 := []int{1, 2}
	num2 := []int{3, 4}

	fmt.Println(findMedianSortedArrays(num1, num2))

	num3 := []int{1, 3}
	num4 := []int{8}

	fmt.Println(findMedianSortedArrays(num3, num4))

}
