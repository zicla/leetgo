package main

import (
	"math"
	"fmt"
)

func isPrime(n int, primeList []int) bool {

	a := int(math.Sqrt(float64(n)))
	for _, v := range primeList {
		if v > a {
			break
		}

		if n%v == 0 {
			return false
		}
	}

	return true
}

func countPrimes(n int) int {
	res := 0
	var primeList []int
	for i := 2; i < n; i++ {
		if isPrime(i, primeList) {
			primeList = append(primeList, i)
			res++
		}
	}
	return res
}
func main() {

	fmt.Printf("%v \n", countPrimes(5))
	fmt.Printf("%v \n", countPrimes(10))
	fmt.Printf("%v \n", countPrimes(1500000))

}
