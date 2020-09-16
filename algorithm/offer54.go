package algorithm

import (
	"algorithm-with-go/ds"
)
// 剑指offer 54.二叉搜索树的第k大节点
// 给定一棵二叉搜索树，请找出其中第k大的节点。
// 示例 1:

// 输入: root = [3,1,4,null,2], k = 1
//    3
//   / \
//  1   4
//   \
//    2
// 输出: 4
// 示例 2:

// 输入: root = [5,3,6,2,4,null,null,1], k = 3
//        5
//       / \
//      3   6
//     / \
//    2   4
//   /
//  1
// 输出: 4

//思路：中序遍历二叉树将元素放入一个数组中，此时数组是升序的，
//然后取第len(数组)-k位置的元素即为第k大
func kthLargest(root *ds.TreeNode, k int) int {
	if root == nil || k < 0{
		return -1
	}
	var arr []int
	findKth(root, &arr)

	return arr[len(arr) - k]
}

//这里必须传数组的地址
func findKth(root *ds.TreeNode, arr *[]int)  {
	if root == nil {
		return
	}
	findKth(root.Left, arr)
	*arr = append(*arr, root.Val)
	findKth(root.Right, arr)
}