// Insertion sort works on the mechanism of add nect element and checking whole series.
// Worst Case- O(n*n)
// Best Case- O(n) – When the array is already sorted
package main

func insertion_sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}

//func main() {
//	arr := []int{12, 67, 85, 45}
//	fmt.Println(insertion_sort(arr))
//}
