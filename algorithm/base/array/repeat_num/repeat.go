package main

import "fmt"

// 给定一个整数数组，判断是否存在重复元素。
func main() {

	arr1 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	arr2 := []int{1, 2, 3, 4}
	arr3 := []int{1, 2, 3, 1}

	fmt.Println(containsDuplicate(arr1))
	fmt.Println(containsDuplicate(arr2))
	fmt.Println(containsDuplicate(arr3))

}

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}
