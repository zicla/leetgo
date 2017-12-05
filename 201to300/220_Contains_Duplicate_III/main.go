package main

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

	N := len(nums)
	if N == 0 || N == 1 {
		return false
	}

	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if (nums[i]-nums[j] <= t && nums[i]-nums[j] >= -t) && j-i <= k {
				return true
			}
		}
	}
	return false
}

func main() {


}
