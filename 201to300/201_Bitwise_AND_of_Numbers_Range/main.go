package main

import (
	"fmt"
	"math"
)


//参考：https://www.cnblogs.com/grandyang/p/4431646.html 多种解法。
func rangeBitwiseAnd(m int, n int) int {

	d := math.MaxInt32
	for (m & d) != (n & d) {
		d <<= 1;
	}
	return m & d;
}

func main() {


	fmt.Printf("%v \n", rangeBitwiseAnd(5, 7))
	fmt.Printf("%v \n", rangeBitwiseAnd(8, 15))
	fmt.Printf("%v \n", rangeBitwiseAnd(4000000, 2147483646))

}
