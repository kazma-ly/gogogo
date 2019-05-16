package common

import (
	"math/rand"
	"time"
)

// Swap 交换数组
func Swap(index1, index2 int, arr []int) {
	temp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp
}

// Random 生成随机数
func Random(n int, val1, val2 int) []int {
	if val1 > val2 {
		return []int{}
	}

	var arr []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		v := rand.Int()%(val1-val2+1) + val1
		arr = append(arr, v)
	}
	return arr
}
