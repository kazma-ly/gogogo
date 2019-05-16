package main

import (
	"reflect"
	"testing"
)

func Test_readMaze(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"first", args{filename: "maze.in"}, [][]int{
			[]int{0, 1, 0, 0, 0},
			[]int{0, 0, 0, 1, 0},
			[]int{0, 1, 0, 1, 0},
			[]int{1, 1, 1, 0, 0},
			[]int{0, 1, 0, 0, 1},
			[]int{0, 1, 0, 0, 0},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readMaze(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readMaze() = %v, want %v", got, tt.want)
			}
		})
	}
}
