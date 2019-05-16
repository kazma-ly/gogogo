package main

import (
	"fmt"
)

func main() {

	fmt.Println(reverseString("Helloi"))
}

// 编写一个函数，其作用是将输入的字符串反转过来。
func reverseString(s string) string {
	length := len(s)

	var resByte []byte
	for i := length - 1; i >= 0; i-- {
		resByte = append(resByte, s[i])
	}
	return string(resByte)
}
