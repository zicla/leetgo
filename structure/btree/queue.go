package btree

import "fmt"

func Queue() {
	var queue []int
	// Push
	queue = append(queue, 1)

	// Top (just get next element, don't remove it)
	x := queue[0]
	fmt.Println(x)

	// Discard top element
	queue = queue[1:]
	// Is empty ?
	if len(queue) == 0 {
		fmt.Println("Queue is empty !")
	}
}
