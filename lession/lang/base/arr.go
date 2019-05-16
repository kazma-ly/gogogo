package main

import "fmt"

func main() {

	// 数组 array
	arr := [...]int{2, 4, 6, 8, 10, 11, 17}
	for _, a := range arr {
		fmt.Println(a)
	}

	var grid [4][5]bool // 4行5列
	fmt.Println(grid)

	// 切片 slice slice本身内部就是引用 传递时不需要再给*号来传递
	s1 := arr[2:6]
	s2 := arr[:]
	fmt.Println("update slice s1")
	updateSlice(s1)
	fmt.Printf("s1 = %d\n", s1)
	fmt.Println("update slice s2")
	updateSlice(s2)
	fmt.Printf("s2 = %d\n", s2)
	fmt.Printf("arr = %d\n", arr)
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))

	// 观察slice的情况
	var arrInfo []int
	for i := 0; i < 100; i++ {
		printSliceInfo(arrInfo)
		arrInfo = append(arrInfo, i)
	}
	fmt.Println(arrInfo)
}

func printSliceInfo(s []int) {
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func updateSlice(vals []int) {
	vals[0] = 100
}
