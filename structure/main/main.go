package main

import "leetgo/structure/avl"

func main() {

	var root *avl.AVLTree
	var hi bool

	avl.InsertAVL(&root, 6, &hi)
	avl.InsertAVL(&root, 4, &hi)
	avl.InsertAVL(&root, 7, &hi)
	avl.InsertAVL(&root, 3, &hi)
	avl.InsertAVL(&root, 5, &hi)

	avl.PrintLevelAVLTree(root)

	avl.DeleteAVL(&root, 7, &hi)

	avl.PrintLevelAVLTree(root)

}
