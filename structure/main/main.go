package main

import (
	"fmt"
	"leetgo/structure/avl"
	"leetgo/structure/btree"
	"leetgo/structure/graph"
	"leetgo/structure/link"
	"leetgo/structure/sort"
	"math"
	"regexp"
	"leetgo/structure/compiler"
)

const INF = math.MaxInt32

func main() {

	testShuntingYard()

}

func testShuntingYard() {

	compiler.CalExpression("1-1+1")
	compiler.CalExpression("(+9)+89+(12*3.0*0.1/0.45)+.34+1")
	compiler.CalExpression("3.1+4.5/(3-(2*1.51))")
	compiler.CalExpression("(+9)+1")
	compiler.CalExpression("(-9)")
	compiler.CalExpression("3 +-.1+ 3 * ( 6 - 1 * 2 ) / 4 - 1")
	compiler.CalExpression("2 + 3+(1+2)")

	compiler.CalExpression("3 + 3 * ( 6 - 1 * 2 ) / 4 - 1")

	compiler.CalExpression("3 +-+ 3 * ( 6 - 1 * 2 ) / 4 - 1")
	//compiler.ShuntingYard("3 +-+ 3.3.1 * ( 6 - 1 * 2 ) / 4 - 1")
}

func testRegex() {
	//传入string，返回string（更加方便）
	fmt.Println("------FindString------")

	// 查找以 hello 开头（忽略大小写），以 Go 结尾的字符串
	text := "/api/matter/fetch/275b3f0d-9cf6-4524-5eb1-dcb815ab9c93/杨过.png"
	reg := regexp.MustCompile(`^/api/matter/fetch/([^/]+)/([^/]+)$`)
	strs := reg.FindStringSubmatch(text)
	fmt.Printf("%q\n", strs)
	// ["Hello 世界！123 Go"]
}

func testLinkList() {
	list := link.CreateLinkList([]int{2, 1, 3, 4, 56, 2, 1})
	link.PrintLinkList(list)

}

func testAVLTree() {

	var root *avl.AVLTree

	avl.InsertAVL(&root, 50)
	avl.InsertAVL(&root, 30)
	avl.InsertAVL(&root, 70)
	avl.InsertAVL(&root, 60)
	avl.InsertAVL(&root, 80)

	avl.InsertAVL(&root, 55)

	avl.PrintLevelAVLTree(root)

	avl.DeleteAVL(&root, 50)
	avl.PrintLevelAVLTree(root)
	avl.DeleteAVL(&root, 30)
	avl.PrintLevelAVLTree(root)
	avl.DeleteAVL(&root, 70)
	avl.PrintLevelAVLTree(root)
	avl.DeleteAVL(&root, 60)
	avl.PrintLevelAVLTree(root)
	avl.DeleteAVL(&root, 80)
	avl.PrintLevelAVLTree(root)
	avl.DeleteAVL(&root, 55)

	avl.PrintLevelAVLTree(root)
}

func testBTree() {

	bt := btree.CreateBinaryTree([]int{1, 2, 3, 4}, 0)
	btree.PrintBinaryTreeGraph(bt)

}

func testDijkstra() {
	matrix := [][]int{
		{0, 12, INF, INF, INF, 16, 14},
		{12, 0, 10, INF, INF, 7, INF},
		{INF, 10, 0, 3, 5, 6, INF},
		{INF, INF, 3, 0, 4, INF, INF},
		{INF, INF, 5, 4, 0, 2, 8},
		{16, 7, 6, INF, 2, 0, 9},
		{14, INF, INF, INF, 8, 9, 0},
	}

	dist := graph.Dijkstra(matrix, 3)
	fmt.Printf("%v\n", dist)

	a := graph.Floyd(matrix, 3, 1)
	fmt.Printf("%v\n", a)
}

func testQuickSort() {
	fmt.Println("Quick Sort")

	a := []int{5, 4, 2, 1, 3, 100, 0, 200, 500}
	sort.QuickSort(a, 0, len(a)-1)

	fmt.Printf("%v \n", a)
}

func testBubbleSort() {

	fmt.Println("BubbleSort")

	a := []int{5, 4, 2, 1, 3, 100, 0, 200, 500}
	sort.BubbleSort(a)

	fmt.Printf("%v\n", a)

}

func testInsertSort() {

	fmt.Println("testInsertSort")

	a := []int{5, 4, 2, 1, 3, 100, 0, 200, 500}
	sort.InsertSort(a)

	fmt.Printf("%v\n", a)

}

func testHeap() {
	arr := []int{}
	arr = btree.HeapInsert(arr, 10)
	arr = btree.HeapInsert(arr, 40)
	arr = btree.HeapInsert(arr, 30)
	arr = btree.HeapInsert(arr, 60)
	arr = btree.HeapInsert(arr, 90)
	arr = btree.HeapInsert(arr, 70)
	arr = btree.HeapInsert(arr, 20)
	arr = btree.HeapInsert(arr, 50)
	arr = btree.HeapInsert(arr, 80)
	fmt.Printf("%v\n", arr)
	arr = btree.HeapInsert(arr, 85)
	fmt.Printf("%v\n", arr)
	arr = btree.HeapRemove(arr, 90)
	fmt.Printf("%v\n", arr)

	a := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	btree.HeapCreate(a)
	fmt.Printf("%v\n", a)

	b := []int{20, 10, 30, 40, 60, 160, 70, 80, 5}
	btree.HeapSort(b)
	fmt.Printf("%v \n", b)
}

func testMergeSort() {
	b := []int{20, 10, 30, 40, 60, 160, 70, 80, 5}
	sort.MergeSort(b, 0, len(b)-1)

	fmt.Printf("%v \n", b)
}

func testBinaryTree() {
	t := []int{1, 2, 3, 4, math.MaxInt64, 6, 7, 8, 9, math.MaxInt64, math.MaxInt64, 104, 104, math.MaxInt64, math.MaxInt64, 112, 112, 334}
	tree := btree.CreateBinaryTree(t, 0)
	btree.PrintLevelBinaryTree(tree)
}

func testBinaryTreePreOrder() {
	t := []int{1, 2, 3, 4, math.MaxInt64, 6, 7, 8, 9, math.MaxInt64, math.MaxInt64, 104, 104, math.MaxInt64, math.MaxInt64, 112, 112, 334}
	tree := btree.CreateBinaryTree(t, 0)
	btree.PrintLevelBinaryTree(tree)
	btree.BinaryTreePreOrder(tree)
	fmt.Println()
	btree.BinaryTreePreOrderUnify(tree)
	fmt.Println()

	btree.BinaryTreeInOrder(tree)
	fmt.Println()
	btree.BinaryTreeInOrderUnify(tree)
	fmt.Println()

	btree.BinaryTreePostOrder(tree)
	fmt.Println()
	btree.BinaryTreePostOrderUnify(tree)
	fmt.Println()

}
