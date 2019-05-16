package main

import (
	"algorithm/common"
	"testing"
)

// 0.171s
func Test_sortArr(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "2", args: args{common.Random(100000, 0, 1000000)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortArr(tt.args.arr)
		})
	}
}

// 0.178s
func Test_sortArr2(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "3", args: args{common.Random(100000, 0, 1000000)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortArr2(tt.args.arr)
		})
	}
}
