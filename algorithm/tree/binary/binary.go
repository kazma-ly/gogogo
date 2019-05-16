package main

import (
	"fmt"
)

// 二叉树

// Node 树节点
type Node struct {
	Val   int   // 节点的值
	Left  *Node // 左树
	Right *Node // 右树
}

// NewNode 新建一个节点
func NewNode(val int, left, right *Node) *Node {
	return &Node{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

func main() {

	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 9}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	n7 := &Node{Val: 7}
	n8 := &Node{Val: 8}
	n9 := &Node{Val: 4}
	n10 := &Node{Val: 10}

	/**
		  	       0
		        1      2
		    3      9
		  5   6  7   8
	         4
	**/

	root := NewNode(0, n1, n2)
	n1.Left = n3
	n1.Right = n4
	n3.Left = n5
	n3.Right = n6
	n4.Left = n7
	n4.Right = n8
	n6.Left = n9
	n9.Left = n10

	maxDep := findDepLength(root)
	fmt.Printf("寻找最大深度: %d\n", maxDep)

	maxDepVal := Node{}
	depQueue := []*Node{n4}
	findDepVal(depQueue, &maxDepVal)
	fmt.Printf("寻找最左边,最深的值为: %v\n", maxDepVal)

	var queue = []*Node{root}
	var maxVal = findMaxVal(queue, 0)
	fmt.Printf("寻找最大值: %d\n", maxVal)
}

func findDepLength(node *Node) int {
	if node == nil {
		return 0
	}

	left := findDepLength(node.Left)
	right := findDepLength(node.Right)

	if left >= right {
		return left + 1
	}
	return right + 1
}

func findDepVal(queue []*Node, maxLeft *Node) {
	if len(queue) <= 0 {
		return
	}

	val := queue[0]
	queue = queue[1:]
	if val.Left != nil {
		queue = append(queue, val.Left)
		*maxLeft = *val.Left
	}
	if val.Right != nil {
		queue = append(queue, val.Right)
	}

	findDepVal(queue, maxLeft)
}

func findMaxVal(queue []*Node, oldVal int) int {
	if len(queue) <= 0 {
		return oldVal
	}

	val := queue[0]
	queue = queue[1:]
	if val.Left != nil {
		queue = append(queue, val.Left)
	}
	if val.Right != nil {
		queue = append(queue, val.Right)
	}

	if oldVal < val.Val {
		oldVal = val.Val
	}
	return findMaxVal(queue, oldVal)
}
