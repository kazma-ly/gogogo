package main

import (
	"algorithm/common"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test", args: args{nums: common.Random(100, 0, 1000)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubbleSort(tt.args.nums)
		})
	}
}
