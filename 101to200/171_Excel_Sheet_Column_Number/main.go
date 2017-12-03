package main

import "fmt"

func titleToNumber(s string) int {

	N := len(s)

	res := 0
	factor := 1
	for i := N - 1; i >= 0; i-- {
		res += factor * int(s[i]-'A'+1)
		factor *= 26
	}
	return res
}

func main() {

	fmt.Printf("%v \n", titleToNumber("A"))
	fmt.Printf("%v \n", titleToNumber("AA"))
	fmt.Printf("%v \n", titleToNumber("AAA"))
	fmt.Printf("%v \n", titleToNumber("AB"))
}
