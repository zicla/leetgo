package snippet

import (
	"strconv"
	"fmt"
)

//把字符串转成数字。
func ConvertStringToDigital() {

	str := "56"

	//string到int
	d, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	fmt.Println(d)

	//string到int64
	d1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(d1)

	//int到string
	data := 54
	s1 := strconv.Itoa(data)
	fmt.Println(s1)

	//int64到string
	s2 := strconv.FormatInt(int64(data), 10)
	fmt.Println(s2)
}
