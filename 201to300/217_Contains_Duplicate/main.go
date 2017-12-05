package main

func containsDuplicate(nums []int) bool {

	dict := make(map[int]bool)
	for _, v := range nums {
		if dict[v] {
			return true
		}
		dict[v] = true

	}
	return false
}

func main() {

}
