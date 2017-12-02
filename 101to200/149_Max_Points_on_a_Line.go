package main

import (
	"fmt"
	. "leetgo/structure/fraction"
)

type Point struct {
	X int
	Y int
}

//用小数存储会失败，必须使用分数存储。
func maxPoints(points []Point) int {

	//首先去除重复的点。
	var pointMap = make(map[int]map[int]int)
	var ps []Point
	for i := 0; i < len(points); i++ {
		p := points[i]

		if _, ok := pointMap[p.X]; !ok {
			pointMap[p.X] = make(map[int]int)
		}

		if _, ok := pointMap[p.X][p.Y]; !ok {
			pointMap[p.X][p.Y] = 0
			ps = append(ps, p)
		}
		pointMap[p.X][p.Y]++
	}

	//求两个点的斜率。
	var findK = func(p1, p2 Point) *Fraction {
		var upper int = p2.Y - p1.Y
		var lower int = p2.X - p1.X

		if lower == 0 {
			return NewMaxFraction()
		} else if upper == 0 {
			return NewZeroFraction()
		} else {
			return NewFraction(upper, lower)
		}
	}

	N := len(ps)
	res := 0
	//以每个点为终点，向其他点画直线，并且数一数直线上有多少点。
	for i := 0; i < N; i++ {
		p1 := ps[i]
		if pointMap[p1.X][p1.Y] > res {
			res = pointMap[p1.X][p1.Y]
		}
		for j := 0; j < N; j++ {

			if i == j {
				continue
			} else {
				p2 := ps[j]
				k := findK(p1, p2)
				num := pointMap[p1.X][p1.Y] + pointMap[p2.X][p2.Y]

				//寻找在k上的其他点。
				for m := 0; m < N; m++ {
					if m == i || m == j {
						continue
					}
					p3 := ps[m]
					fK := findK(p1, p3)

					if fK.Equal(k) {
						num += pointMap[p3.X][p3.Y]
					}
				}
				if num > res {
					res = num
				}

			}

		}
	}
	return res
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
