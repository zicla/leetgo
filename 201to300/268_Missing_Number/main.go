package main

func missingNumber(nums []int) int {

	sum := 0
	for _, v := range nums {
		sum += v
	}
	for i := 0; i <= len(nums); i++ {
		sum = sum - i
	}
	return -sum

}

func main() {

}
