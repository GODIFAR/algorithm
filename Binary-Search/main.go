package main

import (
	"fmt"
	"math/rand"
)

func main() {
	list := make([]int, 1_000)
	for i := 0; i < len(list); i++ {
		list[i] = i + 1
	}

	//r := rand.New(rand.NewSource(1_000_000))
	for i := 0; i < 20; i++ {
		v := rand.Intn(1_000-1) + 1
		fmt.Printf("针对%d进行二分查找:\n", v)
		idx := binarySearch(list, v)
		fmt.Printf("%d的索引位置:[%d]\n", v, idx)
		fmt.Println("-------------------------------------")
	}
}
func binarySearch(list []int, targe int) int {

	low := 0
	high := len(list) - 1
	step := 0
	for {
		step++
		if low <= high {
			mid := (low + high) / 2
			guess := list[mid]
			if guess == targe {
				fmt.Printf("查找了%d次\n", step)
				return mid
			}
			if guess > targe {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
}
