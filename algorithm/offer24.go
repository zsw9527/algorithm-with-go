package algorithm

import (
	"algorithm-with-go/ds"
)

// 剑指 Offer 24. 反转链表
// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

 

// 示例:

// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL

//迭代版
func reverseList(head *ds.ListNode) *ds.ListNode {
	if head == nil {
		return head
	}
	newHead := new(ds.ListNode)

	p := head
	for p != nil {
		q := p.Next
		p.Next = newHead.Next
		newHead.Next = p
		p = q
	}

	return newHead.Next
}

//递归版
func reverseListRecursion(head *ds.ListNode) *ds.ListNode {
	if head == nil {
		return head
	}
	next := head.Next
	if next == nil {
		return head
	}
	newHead := reverseListRecursion(head.Next)
	next.Next = head
	head.Next = nil
	return newHead
}