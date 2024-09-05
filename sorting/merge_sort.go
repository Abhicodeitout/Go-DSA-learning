// Merge sort works on the principle of divide and merge
// memory consuption is more 
package main

import "fmt"

func merge_sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := merge_sort(arr[:len(arr)/2])
	second := merge_sort(arr[len(arr)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	res := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		res = append(res, a[i])
	}
	for ; j < len(b); j++ {
		res = append(res, b[j])
	}
	return res
}
func main() {
	arr := []int{10, 4, 5, 2, 3, 6, 9, 1, 7, 8}
	fmt.Println(merge_sort(arr))
}

🔗
Fast. Merge sort is much faster than bubble sort, being O(n*log(n)) instead of O(n^2).
Stable. Merge sort is also a stable sort which means that values with duplicate keys in the original list will be in the same order in the sorted list.
Cons
🔗
Extra memory. Most sorting algorithms can be performed using a single copy of the original array. Merge sort requires an extra array in memory to merge the sorted subarrays.
Recursive: Merge sort requires many recursive function calls, and function calls can have significant resource overhead.
If you need a sorting algorithm