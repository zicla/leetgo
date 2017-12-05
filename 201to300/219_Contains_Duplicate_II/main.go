package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {

	dict := make(map[int]int)
	for i, v := range nums {
		if dictValue, ok := dict[v]; ok {

			if i-dictValue <= k {

				return true
			}
		}
		dict[v] = i

	}
	return false
}

func main() {

	fmt.Printf("%v \n", containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))

}
