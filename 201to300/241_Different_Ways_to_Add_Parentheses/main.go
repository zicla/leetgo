package main

import (
	"strconv"
	"fmt"
)

func diffWaysToCompute(input string) []int {
	var res []int
	for i := 0; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			left := diffWaysToCompute(input[:i]);
			right := diffWaysToCompute(input[i+1:]);
			for j := 0; j < len(left); j++ {
				for k := 0; k < len(right); k++ {
					if input[i] == '+' {
						res = append(res, left[j]+right[k])
					} else if input[i] == '-' {
						res = append(res, left[j]-right[k])
					} else {
						res = append(res, left[j]*right[k])
					}
				}
			}
		}
	}
	if len(res) == 0 {
		i, _ := strconv.Atoi(input)
		res = append(res, i)
	}
	return res;
}

func main() {

	fmt.Printf("%v\n", diffWaysToCompute("2*3-4*5"))
}
