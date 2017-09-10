package main

import "fmt"

func QuickSort(arr []int, start int, end int) {

	if start >= end {
		return
	}

	headIndex := start
	tailIndex := end

	standarIndex := headIndex
	for headIndex < tailIndex {

		if standarIndex == headIndex {
			if arr[standarIndex] > arr[tailIndex] {
				temp := arr[standarIndex]
				arr[standarIndex] = arr[tailIndex]
				arr[tailIndex] = temp

				standarIndex = tailIndex
				headIndex++
			} else {
				tailIndex--
			}
		} else if standarIndex == tailIndex {

			if arr[standarIndex] < arr[headIndex] {
				temp := arr[standarIndex]
				arr[standarIndex] = arr[headIndex]
				arr[headIndex] = temp

				standarIndex = headIndex
				tailIndex--

			} else {
				headIndex++
			}

		}

	}

	QuickSort(arr, start, standarIndex-1)
	QuickSort(arr, standarIndex+1, end)

}

func main() {
	fmt.Println("Quick Sort")

	a := []int{5, 4, 2, 1, 3, 100, 0, 200, 500}
	QuickSort(a, 0, len(a)-1)

	fmt.Printf("%v", a)
}
