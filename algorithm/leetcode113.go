package algorithm

import (
	"algorithm-with-go/ds"
)

// 113. 路径总和 II
// 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

// 说明: 叶子节点是指没有子节点的节点。

// 示例:
// 给定如下二叉树，以及目标和 sum = 22，

//               5
//              / \
//             4   8
//            /   / \
//           11  13  4
//          /  \    / \
//         7    2  5   1
// 返回:

// [
//    [5,4,11,2],
//    [5,8,4,5]
// ]

func pathSum(root *ds.TreeNode, sum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	pathSumDFS(root, []int{}, &res, sum)
	return res
}

func pathSumDFS(root *ds.TreeNode, cur []int, res *[][]int, sum int) {
	sum -= root.Val
	cur = append(cur, root.Val)
	if root.Left == nil && root.Right == nil && sum == 0 {
		newArr := make([]int, len(cur))
		copy(newArr, cur)
		*res = append(*res, newArr)
		return
	}
	if root.Left != nil {
		pathSumDFS(root.Left, cur, res, sum)
	}
	if root.Right != nil {
		pathSumDFS(root.Right, cur, res, sum)
	}
	cur = cur[:len(cur) - 1]
}