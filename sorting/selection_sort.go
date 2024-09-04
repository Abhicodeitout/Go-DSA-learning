// selection sort works om mechanism of slection of min and swap the values.
// Selection Sort has a time complexity of O ( n 2 ) in all cases (best, average, worst). 
//This makes the algorithm inefficient for large datasets as its performance deteriorates with the increase in the size of the input.
package main

import (
	"fmt"
)

func selection_sort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		// finding the min if min i is greater than any term 
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}

		}
		//swap code 
		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
	return arr
}

func main() {
	arr := []int{13, 46, 24, 52, 20, 9}
	fmt.Println(selection_sort(arr))

}
