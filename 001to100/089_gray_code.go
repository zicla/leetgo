package main

import "fmt"

func grayCode(n int) []int {
	total := 1 << uint(n)
	arr := make([]int, total, total)
	p := 0
	arr[p] = 0
	p++

	for i := 0; i < n; i++ {
		N := 1 << uint(i)
		for j := N - 1; j >= 0; j-- {
			num := arr[j]
			num = num | (1 << uint(i))
			arr[p] = num
			p++
		}
	}

	return arr

}

func main() {

	fmt.Printf("%v\n", grayCode(2))
	fmt.Printf("%v\n", grayCode(3))
	fmt.Printf("%v\n", grayCode(4))

}
