package algorithm

import (
	"algorithm-with-go/ds"
)


func connect(root *ds.Node) *ds.Node {
	if root == nil {
		return root
	}
	//上一层的节点
	q := []*ds.Node{root}
	for len(q) > 0 {
		//tmp当前层
		tmp := q
		q = nil
		for i, node := range tmp {
			if i < len(tmp) - 1 {
				node.Next = tmp[i + 1]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}