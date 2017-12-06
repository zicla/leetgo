package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {

	sum := (C-A)*(D-B) + (G-E)*(H-F)
	//第一种情况，两个矩形不相交
	if E >= C || F >= D || B >= H || A >= G {
		return sum
	}
	//第二种情况，矩形相交，相交的矩形必然是 二者矩形中x较小值和y较小值。
	return sum - (min(G, C)-max(A, E))*(min(D, H)-max(B, F))
}

func main() {


}
