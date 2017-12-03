package main

func majorityElement(nums []int) int {

	major := nums[0]
	count := 0
	for _, v := range nums {

		if count == 0 {
			major = v
			count ++
		} else if major == v {
			count++
		} else {
			count--
		}

	}
	return major
}

func main() {

}
