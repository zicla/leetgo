package main

import "fmt"


//快慢指针。具体分析参见：https://www.cnblogs.com/hiddenfox/p/3408931.html
func findDuplicate(nums []int) int {
	slow := 0
	fast := 0
	t := 0;
	for true {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	for true {
		slow = nums[slow]
		t = nums[t]
		if slow == t {
			break
		}
	}
	return slow
}
func main() {

	fmt.Printf("%v\n", findDuplicate([]int{1, 1}))
}
