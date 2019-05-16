package main

import "fmt"

// 冒泡
func main() {
	arr := []int{1, 4, 6, 2, 4, 3, 11, 6, 87, 0}
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for k := i + 1; k < len(nums); k++ {
			if nums[k] > nums[i] {
				temp := nums[k]
				nums[k] = nums[i]
				nums[i] = temp
			}
		}
	}
}
