// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func quick_sort_start(arr []int) []int {
	return quick_start(arr, 0, len(arr)-1)
}
func quick_start(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quick_start(arr, low, p-1)
		arr = quick_start(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := 0
	for j := 0; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func main() {
	arr := []int{12, 7, 3, 9, 0, 2}
	fmt.Println(quick_sort_start(arr))
}
