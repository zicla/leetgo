package main

import (
	"strings"
	"fmt"
)

func calculate(s string) int {
	//去掉空格。
	s = strings.Replace(s, " ", "", -1)

	var queue []int
	var stack []byte

	//读取的指针p
	p := 0
	N := len(s)

	//能否接受数字。
	canNumber := true
	//当前的符号
	sign := 1

	for p < N {

		if canNumber {

			if s[p] == '-' {
				sign = -sign
			} else if s[p] == '+' {

			} else if s[p] >= '0' && s[p] <= '9' {

				var res = int(s[p] - '0')
				p++
				for p < N && s[p] >= '0' && s[p] <= '9' {
					res = res*10 + int(s[p]-'0')
					p++
				}
				p--

				canNumber = false

				//入队
				queue = append(queue, res*sign)
			}

		} else {
			if s[p] == '*' || s[p] == '/' {
				//一直出栈到空，或者遇到加减号。
				for len(stack) != 0 && (stack[len(stack)-1] == '*' || stack[len(stack)-1] == '/') {
					//出栈入队。入队时符号用负数表示。
					queue = append(queue, -int(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}

				//入栈
				stack = append(stack, s[p])
			} else if s[p] == '+' || s[p] == '-' {
				//当栈顶是空，或者栈顶也是加减号，入栈。
				if len(stack) == 0 {
					stack = append(stack, s[p])
				} else {

					//一直出栈到空
					for len(stack) != 0 {
						//出栈入队。入队时符号用负数表示。
						queue = append(queue, -int(stack[len(stack)-1]))
						stack = stack[:len(stack)-1]
					}
					//最后将自己入栈。
					stack = append(stack, s[p])

				}
			}
			canNumber = true
		}
		p++
	}

	//符号全部出栈入队。
	for len(stack) != 0 {
		//出栈入队。入队时符号用负数表示。
		queue = append(queue, -int(stack[len(stack)-1]))
		stack = stack[:len(stack)-1]
	}

	//计算逆波兰表达式。
	var res []int
	for i := 0; i < len(queue); i++ {

		val := queue[i]
		if val >= 0 {
			//入栈
			res = append(res, val)
		} else {
			if -val == '+' {
				res[len(res)-2] = res[len(res)-2] + res[len(res)-1]
				res = res[:len(res)-1]
			} else if -val == '-' {
				res[len(res)-2] = res[len(res)-2] - res[len(res)-1]
				res = res[:len(res)-1]
			} else if -val == '*' {
				res[len(res)-2] = res[len(res)-2] * res[len(res)-1]
				res = res[:len(res)-1]
			} else if -val == '/' {
				res[len(res)-2] = res[len(res)-2] / res[len(res)-1]
				res = res[:len(res)-1]
			}
		}

	}

	return res[0]
}

func main() {
	fmt.Printf("%s = %v \n", "1-1+1", calculate("1-1+1"))
	fmt.Printf("%s = %v \n", "100000000/1/2/3/4/5/6/7/8/9/10", calculate("100000000/1/2/3/4/5/6/7/8/9/10"))
	fmt.Printf("%s = %v \n", "3+2*2", calculate("3+2*2"))
	fmt.Printf("%s = %v \n", " 3/2 ", calculate(" 3/2 "))
	fmt.Printf("%s = %v \n", " 3+5 / 2 ", calculate(" 3+5 / 2 "))

}
