package node

import "fmt"

// TreeNode node
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (node *TreeNode) SetValue(val int) {
	if node == nil {
		fmt.Println("Set value with nil. Ignored")
		return
	}
	node.Value = val
}

// Traverse 遍历
func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

// TraverseWithChannal 通过channel遍历
func (node *TreeNode) TraverseWithChannal() chan TreeNode {
	out := make(chan TreeNode)
	go func() {
		// out <- node
	}()
	return out
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func (node TreeNode) Print() {
	fmt.Println(node.Value)
}
