package main

import (
	"fmt"
	"math/cmplx"
)

type (
	VoidCallback func(a, b int) int
)

func main() {
	euler()

	val := apply(func(a, b int) int {
		return a * b
	}, 2, 2)

	fmt.Println(val)
}

func apply(callback VoidCallback, a, b int) int {
	return callback(a, b)
}

func euler() {
	c := 3 + 4i // 复数
	fmt.Println(cmplx.Abs(c))
}
