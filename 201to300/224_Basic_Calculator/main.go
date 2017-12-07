package main

import "fmt"

//参考文章：https://www.cnblogs.com/grandyang/p/4570699.html
func calculate(s string) int {
	res := 0
	sign := 1
	n := len(s)
	var queue []int

	for i := 0; i < n; i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			num := 0
			for i < n && s[i] >= '0' && c <= '9' {
				num = 10*num + int(s[i]-'0')
				i++
			}
			res += sign * num
			i--
		} else if c == '+' {
			sign = 1
		} else if c == '-' {
			sign = -1
		} else if c == '(' {
			queue = append(queue, res)
			queue = append(queue, sign)
			res = 0
			sign = 1
		} else if c == ')' {
			//出栈符号
			res *= queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			//出栈之前的元素
			res += queue[len(queue)-1]
			queue = queue[:len(queue)-1]
		}
	}
	return res
}
func main() {

	fmt.Printf("%v \n", calculate("(1)"))
	fmt.Printf("%v \n", calculate("1 + 1"))
	fmt.Printf("%v \n", calculate(" 2-1 + 2 "))
	fmt.Printf("%v \n", calculate("(1+(4 +        521           +     2        )-3)+(6        +    878)"))

}
