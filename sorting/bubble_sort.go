// bubble sort  works on the mechanisam on comparing  the next element if its smaller swap them .
//Its worst-case time complexity is O(n²), and its best-case time complexity is O(n).

package main

import (
	"fmt"
)

func bubble_sort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{12, 45, 65, 34, 76}// input can be taken according to you 
	fmt.Println(bubble_sort(arr))
}
