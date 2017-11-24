package main

import (
	. "leetgo/structure/btree"
	"fmt"
	"math"
)

func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	if root != nil {
		pathSumRecursive(root, sum, []int{}, &res)
	}
	return res
}

func pathSumRecursive(root *TreeNode, sum int, path []int, res *[][]int) {
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			*res = append(*res, append(path, root.Val))
			return
		}
	} else if root.Left == nil {
		pathSumRecursive(root.Right, sum-root.Val, append(path, root.Val), res)

	} else if root.Right == nil {
		pathSumRecursive(root.Left, sum-root.Val, append(path, root.Val), res)

	} else {
		pathSumRecursive(root.Left, sum-root.Val, append(path, root.Val), res)
		pathSumRecursive(root.Right, sum-root.Val, append(path, root.Val), res)
	}
}

func main() {

	tree := CreateBinaryTree([]int{5, 4, 8, 11, math.MinInt64, 13, 4, 7, 2, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 5, 1}, 0)
	fmt.Println(pathSum(tree, 22))

	tree = CreateBinaryTree([]int{-260, -202, -903, -980, -570, -858, 218, 764, -300, 205, math.MinInt64, -35, math.MinInt64, math.MinInt64, -204, 950, -769, 258, -652, 614, -584, 76, 817, -192, math.MinInt64, math.MinInt64, -114, 880, math.MinInt64, -200, 71, 671, 344, 801, 193, -18, 876, -920, -730, 222, 679, math.MinInt64, -680, math.MinInt64, math.MinInt64, math.MinInt64, -859, 744, -261, 692, math.MinInt64, -341, -163, math.MinInt64, math.MinInt64, 482, -979, 205, math.MinInt64, 146, 165, 801, 100, -656, 714, -629, 995, 474, 307, -581, -150, -941, math.MinInt64, math.MinInt64, math.MinInt64, -937, -69, -23, 82, math.MinInt64, -139, -591, math.MinInt64, -453, -861, -370, math.MinInt64, math.MinInt64, math.MinInt64, 216, 233, math.MinInt64, 430, math.MinInt64, 5, -110, math.MinInt64, math.MinInt64, -660, 624, -510, -588, math.MinInt64, math.MinInt64, 381, math.MinInt64, 368, 559, math.MinInt64, 521, -301, math.MinInt64, 522, 379, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 456, 519, math.MinInt64, math.MinInt64, 482, 349, math.MinInt64, math.MinInt64, 19, math.MinInt64, math.MinInt64, 288, -811, math.MinInt64, -372, math.MinInt64, math.MinInt64, -536, math.MinInt64, -404, -457, -740, 860, math.MinInt64, math.MinInt64, -636, math.MinInt64, math.MinInt64, 342, -874, -462, -504, 781, 855, -392, math.MinInt64, math.MinInt64, math.MinInt64, 406, math.MinInt64, -758, 541, math.MinInt64, -947, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, -964, math.MinInt64, 600, -45, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, -194, math.MinInt64, math.MinInt64, math.MinInt64, -802, math.MinInt64, math.MinInt64, math.MinInt64, -3, math.MinInt64, -792, 672, 643, math.MinInt64, 14, math.MinInt64, math.MinInt64, 489, 457, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 412, math.MinInt64, 558, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, -846, 158, -146, math.MinInt64, math.MinInt64, -76, -650, math.MinInt64, -782, math.MinInt64, -127, math.MinInt64, -678, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, -464, -426, math.MinInt64, -366, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 81, -607, 716, math.MinInt64, math.MinInt64, -213, math.MinInt64, 379, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 644, 445, math.MinInt64, math.MinInt64, -419, -845, -720, math.MinInt64, math.MinInt64, -915, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, -686, 594, -243, math.MinInt64, 496, math.MinInt64, 907, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 579, 873, 702, math.MinInt64, math.MinInt64, math.MinInt64, -834, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, -300, -214, -466, math.MinInt64, math.MinInt64, 972, math.MinInt64, math.MinInt64, math.MinInt64, 814, math.MinInt64, -940, math.MinInt64, 763, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, -449, -844, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, -47}, 0)
	fmt.Println(pathSum(tree, -243))

}
