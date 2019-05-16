package main

import (
	"algorithm/common"
	"fmt"
)

// 选择排序
func main() {

	arr1 := common.Random(100, 0, 100)
	fmt.Println(arr1)
	sortArr(arr1)
	fmt.Println(arr1)
}

func sortArr(arr []int) {
	length := len(arr)

	for i := 0; i < length; i++ {
		index := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[index] {
				index = j
			}
		}

		common.Swap(i, index, arr)
	}

}
