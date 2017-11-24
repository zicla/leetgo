package main

import "leetgo/structure/avl"

func main() {

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
