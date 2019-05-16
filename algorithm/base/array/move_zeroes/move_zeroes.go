package main

import "fmt"

// 移动零
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]

// 说明:
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
func main() {

	nums := []int{0, 1, 0, 3, 12}
	moveZeroes1(nums)
	fmt.Println(nums)

	nums1 := []int{0, 0, 1}
	moveZeroes1(nums1)
	fmt.Println(nums1)

}

func moveZeroes(nums []int) {
	length := len(nums)
	count := 0
	for i := 0; i < length; i++ {
		val := nums[i]
		if val == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			if count == length-i {
				break
			}
			i--
			count ++
		}
	}
}

func moveZeroes1(nums []int) {
	s := 0
	f := 0
	for f < len(nums) {
		if nums[f] != 0 {
			nums[s], nums[f] = nums[f], nums[s]
			s++
		}
		f++
	}
}

// 先把不是零的往前移动，然后后面的全部填0 - -
func moveZeroes2(nums []int) {
	j := 0 // 记录填充的position
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j += 1
		}
	}
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}
