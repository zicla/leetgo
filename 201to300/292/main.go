package main

import "fmt"

func canWinNim(n int) bool {

	var remain int = n % 4

	if remain == 0 {
		return false
	} else {
		return true
	}

}

func main() {

	fmt.Println(canWinNim(5))

}
