package main

import (
	"fmt"
	"unsafe"
)

//思路1，找出所有可能的2^n
func isPowerOfTwo1(n int) bool {
	//只需要数1的个数即可。
	var bits = uint(unsafe.Sizeof(n)) * 8
	var i uint = 0
	for i = 0; i < bits; i++ {
		if 1<<i == n {
			return true
		}
	}
	return false
}

//思路2，位运算
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	return (n & (n - 1)) == 0
}

func main() {

	fmt.Printf("%v\n", isPowerOfTwo(1))
	fmt.Printf("%v\n", isPowerOfTwo(2))
	fmt.Printf("%v\n", isPowerOfTwo(4))
	fmt.Printf("%v\n", isPowerOfTwo(-4))
	fmt.Printf("%v\n", isPowerOfTwo(7))

}
