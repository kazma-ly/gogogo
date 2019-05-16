package main

import (
	"fmt"
)

// 删除排序数组中的重复项
func main() {

	var arr = []int{1, 2, 2, 3, 3, 5, 6, 7, 10, 10, 11}
	var last = removeDuplicates(arr)
	fmt.Println(arr[:last])
}

func removeDuplicates(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			num1 := nums[:i-1]
			num2 := nums[i:]
			nums = append(num1, num2...)
			i--
		}
	}
	return len(nums)
}
