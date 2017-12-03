package main

import "fmt"

func convertToTitle(n int) string {

	res := "";
	for n > 0 {
		n--;
		res = fmt.Sprintf("%c%s", n%26+'A', res)
		n /= 26;
	}
	return res;
}

func main() {

	fmt.Printf("%v \n", convertToTitle(27))
	fmt.Printf("%v \n", convertToTitle(29))
	fmt.Printf("%v \n", convertToTitle(703))
	fmt.Printf("%v \n", convertToTitle(1))

}
