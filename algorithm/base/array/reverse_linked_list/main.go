package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}


// 反转一个单链表。
// https://leetcode-cn.com/problems/reverse-linked-list/
func main() {

	n5 := &ListNode{5, nil}
	n4 := &ListNode{4, n5}
	n3 := &ListNode{3, n4}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}

	node := reverseList(n1)

	log.Println(node)
}

func reverseList(head *ListNode) *ListNode {

	var p *ListNode
	var q *ListNode
	var r *ListNode

	head.Next = nil
	q = head.Next

	for q != nil {

	}

	return head
}
