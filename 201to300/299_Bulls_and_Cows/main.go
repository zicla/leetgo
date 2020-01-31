package main

import "fmt"

func getHint(secret string, guess string) string {

	//统计多少个公牛
	var N int = len(secret)
	var bulls int = 0

	//装secret每种的个数
	var sList [10]int = [10]int{0}
	//装guess每种数的个数
	var gList [10]int = [10]int{0}

	for i := 0; i < N; i++ {

		sChar := secret[i]
		gChar := guess[i]

		if sChar == gChar {
			bulls++
		} else {
			sList[sChar-'0']++
			gList[gChar-'0']++
		}
	}

	var cows int = 0
	for i := 0; i < 10; i++ {
		if sList[i] < gList[i] {
			cows += sList[i]
		} else {
			cows += gList[i]
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)

}

func main() {

	a := getHint("1807", "7810")
	fmt.Println(a)

	b := getHint("1123", "0111")
	fmt.Println(b)

}
