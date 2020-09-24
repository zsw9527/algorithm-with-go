package algorithm

import "algorithm-with-go/ds"

// 501. 二叉搜索树中的众数
// 给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

// 假定 BST 有如下定义：

// 结点左子树中所含结点的值小于等于当前结点的值
// 结点右子树中所含结点的值大于等于当前结点的值
// 左子树和右子树都是二叉搜索树
// 例如：
// 给定 BST [1,null,2,2],

//    1
//     \
//      2
//     /
//    2
// 返回[2].

func findMode(root *ds.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var countMap = make(map[int]int)
	treeToSlice(root, countMap)
	max := 0
	for _, v := range countMap {
		if v > max {
			max = v
		}
	}
	for k, v := range countMap {
		if v == max {
			res = append(res, k)
		}
	}
	return res
}

func treeToSlice(root *ds.TreeNode, countMap map[int]int) {
	if root == nil {
		return
	}
	treeToSlice(root.Left, countMap)
	if _, ok := countMap[root.Val]; ok {
		val := countMap[root.Val]
		countMap[root.Val] = val + 1
	} else {
		countMap[root.Val] = 1
	}
	treeToSlice(root.Right, countMap)
}