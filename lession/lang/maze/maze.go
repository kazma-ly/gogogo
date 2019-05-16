package main

import (
	"fmt"
	"os"
)

type point struct {
	i int
	j int
}

// 行走的方向
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func main() {
	maze := readMaze("maze.in")
	fmt.Println("0可以走, 1不可以走")
	printTwoArray(maze)

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	fmt.Println("路径")
	printTwoArray(steps)
}

// 走迷宫 广度优先
func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze)) // 走过的步骤
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0] // 当前要走的步骤
		Q = Q[1:]
		if cur == end {
			break
		}

		for _, dir := range dirs { // 四个方向
			next := cur.add(dir)
			// 越界 或者撞墙了 不能走
			if val, ok := next.at(maze); !ok || val == 1 {
				continue
			}
			// 已经走过了
			if val, ok := next.at(steps); !ok || val != 0 {
				continue
			}
			// 不能回到起点
			if next == start {
				continue
			}
			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1
			Q = append(Q, next)
		}
	}
	return steps
}

// 读取迷宫
func readMaze(filename string) [][]int {
	var row, col int
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fmt.Fscanf(file, "%d %d\n", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			_, err := fmt.Fscanf(file, "%d", &maze[i][j])
			if err != nil {
				fmt.Println(i, j)
				panic(err)
			}
		}
		fmt.Fscanf(file, "\n")
	}
	return maze
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// at int表示1 死路 2 活路,  ok表示是否有值
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func printTwoArray(arr [][]int) {
	for _, vI := range arr {
		for _, vJ := range vI {
			fmt.Printf("%3d ", vJ)
		}
		fmt.Println()
	}
}
