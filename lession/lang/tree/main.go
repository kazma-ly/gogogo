package main

import (
	"fmt"
	"lession/lang/tree/node"
)

type MyTreeNode struct {
	node *node.TreeNode
}

func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	(&MyTreeNode{myNode.node.Left}).postOrder()
	(&MyTreeNode{myNode.node.Right}).postOrder()
	myNode.node.Print()
}

func main() {
	var root node.TreeNode

	root = node.TreeNode{Value: 3}
	root.Left = &node.TreeNode{}
	root.Left = &node.TreeNode{}
	root.Right = &node.TreeNode{5, nil, nil}
	root.Right.Left = new(node.TreeNode)
	root.Left.Right = node.CreateNode(2)

	root.Right.Left.SetValue(10)
	fmt.Println(root.Right.Left.Value)

	root.Print()
	fmt.Println()

	var pRoot *node.TreeNode
	pRoot.SetValue(200)
	fmt.Println()

	root.Traverse()

	fmt.Println("=================")
	mt := MyTreeNode{&root}
	mt.postOrder()
}
