package main

import "fmt"

// 计数排序
func main() {

	arr := []int{1, 4, 6, 2, 4, 3, 11, 6, 87, 0}
	countSort(arr, 87)
	fmt.Println(arr)
}

func countSort(nums []int, maxNum int) {
	count := make([]int, maxNum+1)

	for _, num := range nums {
		count[num]++
	}
	var point = 0             // offset
	for i, c := range count { // i 是待排序的值 c是这个值出现了几次
		for k := 0; k < c; k++ {
			nums[point] = i
			point++
		}
	}
}
