package main

import (
	"fmt"
	"strconv"
)

//辗转相除法。
func fractionToDecimalGCD(a, b int) int {
	//保证a>=b
	if a < b {
		a, b = b, a
	}
	if b <= 0 {
		panic("求公约数只能是正整数")
	}
	if a == b {
		return a
	}
	c := a % b
	if c == 0 {
		return b
	} else {
		return fractionToDecimalGCD(b, c)
	}
}

func fractionToDecimal(numerator int, denominator int) string {

	if numerator == 0 {
		return "0"
	}

	//保证两个数都是正数。
	positive := true
	if numerator > 0 && denominator > 0 {
		positive = true
	} else if numerator > 0 && denominator < 0 {
		positive = false
		denominator = -denominator
	} else if numerator < 0 && denominator > 0 {
		positive = false
		numerator = -numerator
	} else {
		positive = true
		numerator = -numerator
		denominator = -denominator
	}

	//利用辗转相除法，把分数化简。
	gcd := fractionToDecimalGCD(numerator, denominator)
	numerator = numerator / gcd
	denominator = denominator / gcd

	//首先判断整数部分。
	part1 := numerator / denominator
	remain := numerator % denominator
	if remain == 0 {
		if positive {
			return strconv.Itoa(part1)
		} else {
			return strconv.Itoa(-part1)
		}

	}

	//以下都是真分数情况了。开始按照小学的方式做除法了。

	//被除数
	factor := remain

	//商的结果集。
	var resArray []int
	//该余数是否以前出现过，如果出现过，出现的位置在哪里。
	remainMap := make(map[int]int)
	//开始循环的index
	repeatStart := -1
	repeatEnd := -1

	//那么初始的时候余数应放进去。这里相当于透支了restArray
	remainMap[factor] = 0

	factor = factor * 10
	for true {

		//判断被除数能否整除。
		shang := factor / denominator
		yu := factor % denominator

		//将商写入到结果集中去。
		resArray = append(resArray, shang)

		//这里已经除尽了。
		if yu == 0 {
			break
		} else {

			//判断被除数以前出现过没。
			if b, ok := remainMap[yu]; ok {
				//出现过，那么自然该就此打住了。
				repeatStart = b
				repeatEnd = len(resArray) - 1
				break
			} else {
				//收集好余数。
				remainMap[yu] = len(resArray)
				//余数*10继续下一步。
				factor = yu * 10
			}
		}
	}

	if repeatStart == -1 {
		//有限小数
		str := fmt.Sprintf("%d", part1)
		if !positive {
			str = fmt.Sprintf("-%d", part1)
		}

		//小数部分。
		factorPart := ""
		for _, v := range resArray {
			factorPart = fmt.Sprintf("%s%d", factorPart, v)
		}

		//合起来的部分。
		str = fmt.Sprintf("%s.%s", str, factorPart)
		return str

	} else {

		//无穷小数
		str := fmt.Sprintf("%d", part1)
		if !positive {
			str = fmt.Sprintf("-%d", part1)
		}

		//小数部分。
		factorPart := ""

		//循环之前的部分。
		for i := 0; i < repeatStart; i++ {
			factorPart = fmt.Sprintf("%s%d", factorPart, resArray[i])
		}

		factorPart = fmt.Sprintf("%s(", factorPart)

		//循环的部分。
		for i := repeatStart; i <= repeatEnd; i++ {
			factorPart = fmt.Sprintf("%s%d", factorPart, resArray[i])
		}

		factorPart = fmt.Sprintf("%s)", factorPart)

		//合起来的部分。
		str = fmt.Sprintf("%s.%s", str, factorPart)
		return str

	}

}
func main() {

	fmt.Printf("%v \n", fractionToDecimal(-2147483648, 1))
	fmt.Printf("%v \n", fractionToDecimal(22, 7))
	fmt.Printf("%v \n", fractionToDecimal(1, 7))

	fmt.Printf("%v \n", fractionToDecimal(1, 6))

	fmt.Printf("%v \n", fractionToDecimal(2, 3))
	fmt.Printf("%v \n", fractionToDecimal(0, 3))

	fmt.Printf("%v \n", fractionToDecimal(1, 2))
	fmt.Printf("%v \n", fractionToDecimal(2, 1))

}
