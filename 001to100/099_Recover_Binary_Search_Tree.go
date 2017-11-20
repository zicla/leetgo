package main

import (
	"math"
	"fmt"
	. "leetgo/structure/btree"
)

func recoverTree(root *TreeNode) {

	//准备中序遍历。
	list := make([]*TreeNode, math.MaxInt16)
	N := []int{0}
	recoverTreeScan(root, list, N)
	correctArray(list, N)
}

func recoverTreeScan(root *TreeNode, list []*TreeNode, N []int) {

	if root == nil {
		return
	}
	recoverTreeScan(root.Left, list, N)

	list[N[0]] = root
	N[0]++

	recoverTreeScan(root.Right, list, N)

}

func correctArray(list []*TreeNode, N []int) {

	n := N[0]
	index1 := -1
	index2 := -1
	for i := 0; i < n-1; i++ {

		if list[i].Val > list[i+1].Val {
			if index1 == -1 {
				index1 = i

				index2 = i + 1
				for j := i + 2; j < n; j++ {
					if list[j].Val < list[index2].Val {
						index2 = j
					}
				}

			}

		}
	}
	if index1 == -1 {
		return
	}
	if index2 == -1 {
		index2 = index1 + 1
	}
	tmp := list[index1].Val
	list[index1].Val = list[index2].Val
	list[index2].Val = tmp
}

func printBTree(root *TreeNode) {
	if root == nil {
		return
	}
	printBTree(root.Left)
	fmt.Printf("%v ", root.Val)
	printBTree(root.Right)

}

func main() {
	//tree := &TreeNode{
	//	Val: 8,
	//	Left: &TreeNode{
	//		Val: 4,
	//		Left: &TreeNode{
	//			Val: 2,
	//			Left: &TreeNode{
	//				Val: 1},
	//			Right: &TreeNode{
	//				Val: 3}},
	//		Right: &TreeNode{
	//			Val: 6,
	//			Left: &TreeNode{
	//				Val: 7},
	//			Right: &TreeNode{
	//				Val: 5}}},
	//	Right: &TreeNode{
	//		Val: 10,
	//		Left: &TreeNode{
	//			Val: 9},
	//		Right: &TreeNode{
	//			Val: 11}}}

	tree := CreateBinaryTree([]int{8, 4, 10, 2, 6, 9, 11, 1, 3, 7, 5}, 0)
	printBTree(tree)
	fmt.Println()
	recoverTree(tree)
	printBTree(tree)
	fmt.Println()

	tree = CreateBinaryTree([]int{0, 1}, 0)
	printBTree(tree)
	fmt.Println()
	recoverTree(tree)
	printBTree(tree)
	fmt.Println()


	a := []int{4091, 846, 10073, 117, 979, 6182, 21113, -93, 390, 867, 2140, 4405, 6582, 14837, 21781, -649, 21, 264, 659, math.MinInt64, 882, 1359, 2671, 4197, 6092, 6410, 7887, 10122, 17272, 21596, 23688, math.MinInt64, -488, -41, 26, 134, 383, 552, 811, math.MinInt64, math.MinInt64, 1095, 2122, 2639, 3522, 4148, 4274, 4791, 6139, 6305, 6448, 7356, 9199, math.MinInt64, 14690, 15570, 18038, 21394, 21675, 22768, 23739, -636, -153, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 211, 324, 385, 484, 627, 716, math.MinInt64, 1017, 1298, 1942, math.MinInt64, 2335, math.MinInt64, 3131, 3864, math.MinInt64, math.MinInt64, math.MinInt64, 4338, 4474, 5663, math.MinInt64, math.MinInt64, 6262, 6359, math.MinInt64, 6569, 6795, 7813, 8704, 9390, 10333, 14836, 14972, 16537, 17346, 19282, 21191, 21475, math.MinInt64, 21772, 22317, 22890, 23697, 23767, math.MinInt64, -568, -341, math.MinInt64, 188, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 420, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 1251, math.MinInt64, 1396, 2044, 2250, 2353, 3072, 3174, 3834, 3923, math.MinInt64, 4349, math.MinInt64, 4559, 5292, 5776, math.MinInt64, 6278, math.MinInt64, 6398, 6500, math.MinInt64, 6721, 6931, 7735, math.MinInt64, 8504, 8840, 9343, 9921, 10158, 11686, 14806, math.MinInt64, 14895, 15199, 15985, 17001, math.MinInt64, 17892, 19030, 19865, 21154, 21336, math.MinInt64, 21562, math.MinInt64, math.MinInt64, 22183, 22707, 22814, 23507, math.MinInt64, math.MinInt64, math.MinInt64, 23894, math.MinInt64, -506, -424, -253, math.MinInt64, 202, math.MinInt64, math.MinInt64, 1174, math.MinInt64, math.MinInt64, 1751, 1994, 2119, 2232, math.MinInt64, math.MinInt64, 2561, 2682, math.MinInt64, math.MinInt64, 3274, 3676, math.MinInt64, math.MinInt64, 4003, math.MinInt64, math.MinInt64, math.MinInt64, 4705, 5155, 5555, 5753, 5906, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 6540, 6644, 6755, 6915, 7053, 7373, math.MinInt64, 8432, 8657, 8800, 9135, 9312, math.MinInt64, 9616, 10021, math.MinInt64, 10268, 11168, 11950, 14734, math.MinInt64, math.MinInt64, math.MinInt64, 15027, 15552, 15697, 16470, 16974, 17146, 17724, 17956, 18357, 19113, 19513, 19958, math.MinInt64, math.MinInt64, 21272, math.MinInt64, math.MinInt64, math.MinInt64, 21824, 22231, 22372, 22740, 22803, math.MinInt64, 22991, 23529, 23832, 24020, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, -236, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 1569, 1940, math.MinInt64, math.MinInt64, 2057, math.MinInt64, math.MinInt64, math.MinInt64, 2548, math.MinInt64, 2674, 2824, 3266, 3465, 3531, 3818, math.MinInt64, 4068, 4633, math.MinInt64, 4955, 5233, 5468, 5619, 5731, math.MinInt64, 5809, 5938, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 6883, math.MinInt64, 6957, 7261, math.MinInt64, 7584, 8161, 8452, 8646, 8687, math.MinInt64, 8807, 9079, 9179, 9275, math.MinInt64, 9534, 9824, math.MinInt64, math.MinInt64, 10209, math.MinInt64, 10400, 11569, 11858, 12704, math.MinInt64, 14755, math.MinInt64, 15035, 15507, math.MinInt64, 15649, 15856, 16003, math.MinInt64, 16690, math.MinInt64, 17049, 17175, 17467, 17779, math.MinInt64, math.MinInt64, 18258, 18884, math.MinInt64, 19196, 19457, 19832, 19888, 20594, math.MinInt64, math.MinInt64, math.MinInt64, 22097, 22187, 22253, 22338, 22657, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 22986, 23235, math.MinInt64, 23575, math.MinInt64, math.MinInt64, 23948, math.MinInt64, math.MinInt64, math.MinInt64, 1494, 1662, 1842, math.MinInt64, math.MinInt64, math.MinInt64, 2512, math.MinInt64, math.MinInt64, math.MinInt64, 2745, 2953, math.MinInt64, math.MinInt64, 3370, math.MinInt64, math.MinInt64, 3598, 3775, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 4824, 5085, math.MinInt64, math.MinInt64, 5391, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 6019, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 7227, 7339, 7521, 7684, 7997, 8337, math.MinInt64, 8470, 8511, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 8942, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 9460, math.MinInt64, 9680, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 10478, 11396, 11633, 11836, math.MinInt64, 12059, 13791, math.MinInt64, math.MinInt64, math.MinInt64, 15040, 15269, math.MinInt64, 15598, 15684, 15776, 15902, math.MinInt64, 16148, 16594, 16773, math.MinInt64, math.MinInt64, math.MinInt64, 17206, 17381, 17658, 17726, 17842, 18106, math.MinInt64, 18445, 18981, math.MinInt64, math.MinInt64, 19344, math.MinInt64, 19609, math.MinInt64, math.MinInt64, 19909, 20485, 20640, 21935, 22110, math.MinInt64, 14511, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 22566, math.MinInt64, math.MinInt64, math.MinInt64, 23146, 23319, math.MinInt64, 23645, math.MinInt64, 24006, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 2355, math.MinInt64, math.MinInt64, math.MinInt64, 2878, 3020, math.MinInt64, math.MinInt64, 3587, math.MinInt64, math.MinInt64, 3804, math.MinInt64, 4915, 5036, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 7158, math.MinInt64, math.MinInt64, math.MinInt64, 7431, math.MinInt64, math.MinInt64, 7705, 7955, 8008, 8253, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 8513, 8852, 8969, math.MinInt64, math.MinInt64, math.MinInt64, 9773, 10407, 10600, 11245, 11454, math.MinInt64, 11674, 11698, math.MinInt64, 12031, 12548, 13100, 13855, math.MinInt64, 15128, 15232, 15450, math.MinInt64, 15627, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 15953, 16049, 16339, 16584, 16655, math.MinInt64, 16857, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 17567, math.MinInt64, math.MinInt64, math.MinInt64, 17833, math.MinInt64, math.MinInt64, 18253, 18393, 18743, math.MinInt64, math.MinInt64, 19316, 19380, math.MinInt64, 19687, math.MinInt64, math.MinInt64, 20007, 20563, math.MinInt64, 20783, 21869, 22063, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 22377, 22567, 23011, math.MinInt64, 23316, 23385, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 2448, math.MinInt64, math.MinInt64, math.MinInt64, 3023, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 4838, math.MinInt64, 4965, math.MinInt64, 7064, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 8080, math.MinInt64, math.MinInt64, 8512, 8598, math.MinInt64, 8868, math.MinInt64, 8984, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 10530, 10937, math.MinInt64, 11298, math.MinInt64, 11470, math.MinInt64, math.MinInt64, math.MinInt64, 11774, math.MinInt64, math.MinInt64, 12073, 12626, 12795, 13580, math.MinInt64, 13935, math.MinInt64, math.MinInt64, 15214, math.MinInt64, 15358, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 16053, 16233, 16448, math.MinInt64, math.MinInt64, 16652, 16662, 16808, 16874, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 18211, 18254, math.MinInt64, math.MinInt64, 18512, 18826, math.MinInt64, 19341, math.MinInt64, math.MinInt64, math.MinInt64, 19758, 19982, 20106, math.MinInt64, math.MinInt64, 20668, 20854, math.MinInt64, math.MinInt64, 22034, 22065, math.MinInt64, 22473, math.MinInt64, math.MinInt64, math.MinInt64, 23100, math.MinInt64, math.MinInt64, math.MinInt64, 23431, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 10811, 11059, math.MinInt64, math.MinInt64, math.MinInt64, 11560, math.MinInt64, math.MinInt64, math.MinInt64, 12306, 12549, math.MinInt64, 12741, 12981, 13281, 13582, 13866, 14315, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 16073, 16153, 16244, 16437, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 18168, 18236, math.MinInt64, math.MinInt64, math.MinInt64, 18677, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 19787, math.MinInt64, math.MinInt64, math.MinInt64, 20391, 20652, 20747, math.MinInt64, 21023, 21959, math.MinInt64, math.MinInt64, 22071, 22446, math.MinInt64, 23072, math.MinInt64, math.MinInt64, math.MinInt64, 10638, 10890, 10980, 11143, math.MinInt64, math.MinInt64, 12157, 12500, math.MinInt64, math.MinInt64, 12740, math.MinInt64, 12883, 13045, 13139, 13361, math.MinInt64, 13709, math.MinInt64, 13912, 14135, 14566, math.MinInt64, math.MinInt64, math.MinInt64, 16182, math.MinInt64, 16332, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 18599, math.MinInt64, math.MinInt64, math.MinInt64, 20361, 20422, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 20984, 21064, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 10715, 10886, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 12255, 12366, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 13211, math.MinInt64, 13396, 13681, math.MinInt64, 13897, math.MinInt64, 14078, 14230, 14465, 14618, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 20193, math.MinInt64, math.MinInt64, math.MinInt64, 20920, 21010, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 10863, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 12411, math.MinInt64, math.MinInt64, math.MinInt64, 13484, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 14010, math.MinInt64, 14210, 14263, 14407, 14545, math.MinInt64, math.MinInt64, 20117, 20355, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 12465, 13477, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64, 14322, math.MinInt64, 22188, math.MinInt64, math.MinInt64, math.MinInt64, 20258}
	tree1 := CreateBinaryTree(a,0)
	printBTree(tree1)
	fmt.Println()
	recoverTree(tree1)
	printBTree(tree1)
	fmt.Println()

}
