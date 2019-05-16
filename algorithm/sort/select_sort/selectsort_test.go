package main

import (
	"algorithm/common"
	"testing"
)

func Test_sort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "2", args: args{common.Random(100, 0, 100)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortArr(tt.args.arr)
		})
	}
}
