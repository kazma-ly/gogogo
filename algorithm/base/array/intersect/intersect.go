package main

import "fmt"

// 给定两个数组，编写一个函数来计算它们的交集。
/**
	输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
	我们可以不考虑输出结果的顺序。
进阶:
	如果给定的数组已经排好序呢？你将如何优化你的算法？
	如果 nums1 的大小比 nums2 小很多，哪种方法更优？
	如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
 */
func main() {

	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(intersect(nums1, nums2))

	nums3 := []int{4, 9, 5}
	nums4 := []int{9, 4, 9, 8, 4}

	fmt.Println(intersect(nums3, nums4))
}

func intersect(nums1 []int, nums2 []int) []int {



	return []int{}
}
