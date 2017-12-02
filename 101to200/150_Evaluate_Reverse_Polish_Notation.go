package main

import (
	"strconv"
	"fmt"
)

func evalRPN(tokens []string) int {

	N := len(tokens)
	if N == 0 {
		return 0
	}

	//准备一个栈。
	var stack = []int{}

	for _, str := range tokens {

		if str == "+" || str == "-" || str == "*" || str == "/" {

			//弹出两个数字出来。
			num1 := stack[0]
			num2 := stack[1]
			stack = stack[2:]

			var num int
			//注意此处顺序要反一下。
			switch str {
			case "+":
				num = num2 + num1
			case "-":
				num = num2 - num1
			case "*":
				num = num2 * num1
			case "/":
				num = num2 / num1
			}
			//入栈
			stack = append([]int{num}, stack...)

		} else {

			//入栈
			a, _ := strconv.Atoi(str)
			stack = append([]int{a}, stack...)
		}

	}
	return stack[0]
}

func main() {

	a := evalRPN([]string{"2", "1", "+", "3", "*"})
	b := evalRPN([]string{"4", "13", "5", "/", "+"})
	fmt.Println(a)
	fmt.Println(b)
	c := evalRPN([]string{"0", "3", "/"})
	fmt.Println(c)

}
