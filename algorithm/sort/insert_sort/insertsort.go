package main

import (
	"algorithm/common"
	"fmt"
)

// 插入排序
// 近乎有序的数组，排序很快，因为break
func main() {
	arr := common.Random(10, 0, 100)
	fmt.Println(arr)
	sortArr2(arr)
	fmt.Println(arr)
}

// eg: 1
func sortArr(arr []int) {
	length := len(arr)

	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			// 当前这个数(也就是待对比的)和前一个数比较
			if arr[j] < arr[j-1] {
				t := j - 1
				common.Swap(j, t, arr)
			} else {
				break
			}
		}
	}

}

// eg: 2 减少数组交换
func sortArr2(arr []int) {
	length := len(arr)

	for i := 1; i < length; i++ {
		e := arr[i] // 需要比较和替换的数值
		p := i      // 需要替换的位置
		for j := i; j > 0; j-- {
			t := j - 1
			if e < arr[t] {
				arr[j] = arr[t]
				p = t
			} else {
				break
			}
		}
		arr[p] = e
	}
}
