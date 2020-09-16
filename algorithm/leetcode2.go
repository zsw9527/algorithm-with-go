package algorithm

import "algorithm-with-go/ds"

// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 示例：

// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

func addTwoNumbers(l1 *ds.ListNode, l2 *ds.ListNode) *ds.ListNode {
	res := &ds.ListNode{
		Val: -1,
	}
	p := res
	a := l1
	b := l2
	//进位
	carry := 0
	for a != nil || b != nil{
		var x, y int
		if a == nil {
			x = 0
		} else {
			x = a.Val
		}
		if b == nil {
			y = 0
		} else {
			y = b.Val
		}
		sum := carry + x + y
		p.Next = &ds.ListNode{
			Val: sum % 10,
		}
		carry = carry / 10
		p = p.Next
		if a != nil {
			a = a.Next    
		}
		if b != nil {
			b = b.Next
		}
	}
	if carry > 0 {
		p.Next = &ds.ListNode{
			Val: carry,
		}
	}
	return res.Next
}