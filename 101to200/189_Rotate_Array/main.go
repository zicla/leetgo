package main

func rotate1Time(nums []int) {
	N := len(nums)
	tmp := nums[N-1]
	for i := N - 2; i >= 0; i-- {
		nums[i+1] = nums[i]
	}
	nums[0] = tmp
}
func rotate(nums []int, k int) {
	N := len(nums)
	if N == 0 || N == 1 {
		return
	}

	for i := 0; i < k; i++ {
		rotate1Time(nums)
	}

}
func main() {

}
