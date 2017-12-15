package main

func singleNumber(nums []int) []int {

	var dict = make(map[int]int)
	for _, v := range nums {
		dict[v]++
	}
	var res []int
	for k, v := range dict {
		if v == 1 {
			res = append(res, k)
		}

	}
	return res
}
func main() {

}
