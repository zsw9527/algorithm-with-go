package algorithm

import (
	"algorithm-with-go/ds"
)

// 404. 左叶子之和
// 计算给定二叉树的所有左叶子之和。

// 示例：

//     3
//    / \
//   9  20
//     /  \
//    15   7

// 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

func sumOfLeftLeaves(root *ds.TreeNode) int {
	if root == nil {
        return 0
    }
	sum := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil{
		sum += root.Left.Val
	}
	return sum + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}