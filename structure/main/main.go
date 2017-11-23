package main

import "leetgo/structure/avl"

func main() {

	var root *avl.AVLTree
	var hi bool

	avl.InsertAVL(&root, 50, &hi)
	avl.InsertAVL(&root, 30, &hi)
	avl.InsertAVL(&root, 70, &hi)
	avl.InsertAVL(&root, 60, &hi)
	avl.InsertAVL(&root, 80, &hi)

	avl.PrintLevelAVLTree(root)

	avl.InsertAVL(&root, 90, &hi)

	avl.PrintLevelAVLTree(root)

}
