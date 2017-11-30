package main

import (
	"leetgo/structure/avl"
	"leetgo/structure/link"
	"leetgo/structure/btree"
	"math"
	"leetgo/structure/graph"
	"fmt"
	"leetgo/structure/sort"
)

const INF = math.MaxInt32

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

func main() {

	testInsertSort()

}
