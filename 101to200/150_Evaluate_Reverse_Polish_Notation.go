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

	var add = func(a, b int) int {
		return a + b
	}
	var minus = func(a, b int) int {
		return a - b
	}
	var times = func(a, b int) int {
		return a * b
	}
	var divide = func(a, b int) int {
		return a / b
	}

	operMap := make(map[string]func(a, b int) int)
	operMap["+"] = add
	operMap["-"] = minus
	operMap["*"] = times
	operMap["/"] = divide

	//准备一个栈。
	var stack = []string{tokens[0]}
	i := 1
	for i < N {

		str := tokens[i]
		i++

		var oper func(a, b int) int
		if str == "+" || str == "-" || str == "*" || str == "/" {
			oper = operMap[str]

			//弹出两个数字出来。
			str1 := stack[0]
			str2 := stack[1]
			stack = stack[2:]

			num1, _ := strconv.Atoi(str1)
			num2, _ := strconv.Atoi(str2)
			//注意此处顺序要反一下。
			num := oper(num2, num1)

			//入栈
			numStr := strconv.Itoa(num)
			stack = append([]string{numStr}, stack...)

		} else {

			//入栈
			stack = append([]string{str}, stack...)
		}

	}
	n, _ := strconv.Atoi(stack[0])
	return n
}

func main() {

	a := evalRPN([]string{"2", "1", "+", "3", "*"})
	b := evalRPN([]string{"4", "13", "5", "/", "+"})
	fmt.Println(a)
	fmt.Println(b)
	c := evalRPN([]string{"0", "3", "/"})
	fmt.Println(c)

}
