package main

import (
	"math"
	"fmt"
)

func countPrimes(n int) int {
	res := 0
	var notPrimeList = make([]bool, n)
	for i := 2; i < n; i++ {
		if !notPrimeList[i] {
			res++
			for j := 2; i*j < n; j++ {
				notPrimeList[i*j] = true
			}
		}
	}
	return res
}
func main() {

	fmt.Printf("%v \n", countPrimes(5))
	fmt.Printf("%v \n", countPrimes(10))
	fmt.Printf("%v \n", countPrimes(1500000))

}
