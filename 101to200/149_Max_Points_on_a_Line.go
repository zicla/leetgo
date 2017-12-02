package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

//三点共线的行列式：
/*

|x1 y1 1|
|x2 y2 1| = 0
|x3 y3 1|

*/
func maxPoints(points []Point) int {

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	res := 0;
	N := len(points)
	for i := 0; i < N; i++ {
		duplicate := 1;
		for j := i + 1; j < N; j++ {
			cnt := 0;
			x1 := points[i].X
			y1 := points[i].Y
			x2 := points[j].X
			y2 := points[j].Y
			if x1 == x2 && y1 == y2 {
				duplicate++
				continue
			}
			for k := 0; k < N; k++ {
				x3 := points[k].X
				y3 := points[k].Y
				if x1*y2+x2*y3+x3*y1-x3*y2-x2*y1-x1*y3 == 0 {
					cnt++
				}
			}
			res = max(res, cnt);
		}
		//这里为了防止所有点都在同一个位置。
		res = max(res, duplicate);
	}
	return res;
}

func main() {

	res2 := maxPoints([]Point{{X: 2, Y: 3}, {X: 3, Y: 3}, {X: -5, Y: 3}})
	fmt.Printf("%v \n", res2)

	res1 := maxPoints([]Point{{X: 0, Y: 0}, {X: 94911151, Y: 94911150}, {X: 94911152, Y: 94911151}})
	fmt.Printf("%v \n", res1)

	res := maxPoints([]Point{{X: 1, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 2}, {X: 2, Y: 2}, {X: 2, Y: 0}})
	fmt.Printf("%v \n", res)
	res = maxPoints([]Point{{X: 0, Y: 0}})
	fmt.Printf("%v \n", res)
	res = maxPoints([]Point{{X: 84, Y: 250}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -70}, {X: 0, Y: -70}, {X: 1, Y: -1}, {X: 21, Y: 10}, {X: 42, Y: 90}, {X: -42, Y: -230}})
	fmt.Printf("%v \n", res)

}
