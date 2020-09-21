package algorithm

import "algorithm-with-go/ds"

//538. 把二叉搜索树转换为累加树
//给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。
//
//
//
//例如：
//
//输入: 原始二叉搜索树:
//	 5
//  / \
// 2   13
//
//输出: 转换为累加树:
//  18
// /  \
//20   13

//反序中序遍历（右中左）
func convertBST(root *ds.TreeNode) *ds.TreeNode {
	sum := 0
	var dfs func(*ds.TreeNode)
	dfs = func(root *ds.TreeNode) {
		if root != nil {
			dfs(root.Right)
			sum = sum + root.Val
			root.Val = sum
			dfs(root.Left)
		}
	}
	dfs(root)
	return root
}