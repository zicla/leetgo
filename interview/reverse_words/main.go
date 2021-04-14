package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	parts := strings.Split(str, " ")
	length := len(parts)

	for i := 0; i < length/2; i++ {
		//依次交换。
		temp := parts[i]
		parts[i] = parts[length-1-i]
		parts[length-1-i] = temp
	}

	return strings.Join(parts, " ")
}

func main() {

	fmt.Println("Hello World")

	fmt.Println(reverseWords("Hello world java"))
	fmt.Println(reverseWords("He llo world java"))
	fmt.Println(reverseWords(""))
}
