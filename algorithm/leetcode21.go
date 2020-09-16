package algorithm

import "algorithm-with-go/ds"

// 21. 合并两个有序链表

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 示例：

// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4

func mergeTwoLists(l1 *ds.ListNode, l2 *ds.ListNode) *ds.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	newHead := &ds.ListNode{
		Val: -1,
	}

	q := newHead
	p1 := l1
	p2 := l2

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			q.Next = p1
			p1 = p1.Next
		} else {
			q.Next = p2
			p2 = p2.Next
		}
		q = q.Next
	}

	if p1 != nil {
		q.Next = p1
	}
	if p2 != nil {
		q.Next = p2
	}
	return newHead.Next
}