package main

import "fmt"

func main() {
	arr := []int{5, 89, 4, 8, 1, 6, 48, 456, 4, 8, 5, 12, 3, 99, 87, 9, 4}
	fmt.Println(quickSort(arr))
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int
	for _, ele := range arr[1:] {
		if ele <= pivot {
			left = append(left, ele)
		} else {
			right = append(right, ele)
		}
	}
	return append(quickSort(left), append([]int{pivot}, quickSort(right)...)...)
}
