package main

func majorityElement(nums []int) int {
	dict := make(map[int]int)
	N := len(nums)
	for _, v := range nums {
		dict[v]++
		if dict[v] > N/2 {
			return v
		}
	}
	return 0
}

func main() {

}
