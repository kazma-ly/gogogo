package main

import (
	"fmt"
)

// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
func main() {

	k := 10
	nums := []int{1, 3, 5, 20, 3, 4, 7}
	rotate(nums, k)

	fmt.Printf("%v, %d, %d", nums, cap(nums), len(nums))

}

func rotate(nums []int, k int) {
	if k < 0 {
		return
	}
	n := len(nums)
	if k == 0 || k == n {
		return
	}
	if k > n {
		k %= n
	}
	for i := 0; i < k; i++ {
		first := []int{nums[n-1]}
		right := nums[:n-1]
		copy(nums, append(first, right...))
	}
}

func rotateOther(nums []int, k int) {
	// 我已经假定 k >= 0

	n := len(nums)

	if k > n {
		k %= n
	}
	if k == 0 || k == n {
		return
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
