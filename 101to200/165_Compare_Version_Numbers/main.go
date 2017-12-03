package main

import (
	"strings"
	"fmt"
	"strconv"
)

func compareVersion(version1 string, version2 string) int {

	arr1 := strings.Split(version1, ".")
	arr2 := strings.Split(version2, ".")

	N1 := len(arr1)
	N2 := len(arr2)

	//全部转成数字
	list1 := make([]int, N1)
	list2 := make([]int, N2)
	for k, v := range arr1 {
		i, _ := strconv.Atoi(v)
		list1[k] = i
	}

	//末尾有几个零就去掉几个零。

	for k, v := range arr2 {
		i, _ := strconv.Atoi(v)
		list2[k] = i
	}

	if N1 > N2 {
		for i := 0; i < N2; i++ {
			if list1[i] > list2[i] {
				return 1
			} else if list1[i] < list2[i] {
				return -1
			} else {
				continue
			}
		}

		//判断一下剩下的是不是全部都为零。
		for i := N2; i < N1; i++ {
			if list1[i] != 0 {
				return 1
			}
		}
		return 0
	} else if N1 < N2 {
		for i := 0; i < N1; i++ {
			if list1[i] > list2[i] {
				return 1
			} else if list1[i] < list2[i] {
				return -1
			} else {
				continue
			}
		}

		//判断一下剩下的是不是全部都为零。
		for i := N1; i < N2; i++ {
			if list2[i] != 0 {
				return -1
			}
		}
		return 0
	} else {
		for i := 0; i < N1; i++ {
			if list1[i] > list2[i] {
				return 1
			} else if list1[i] < list2[i] {
				return -1
			} else {
				continue
			}
		}
		return 0
	}
}
func main() {

	//0.1 < 1.1 < 1.2 < 13.37
	fmt.Printf("%v \n", compareVersion("0.1", "0.1.2"))
	fmt.Printf("%v \n", compareVersion("0.1", "0.1.0"))
	fmt.Printf("%v \n", compareVersion("13.37", "0.1.2"))
	fmt.Printf("%v \n", compareVersion("13.37", "12.23.1"))
}
