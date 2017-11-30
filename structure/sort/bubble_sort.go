package sort

func BubbleSort(arr []int) {

	N := len(arr)
	for i := N - 1; i >= 1; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
