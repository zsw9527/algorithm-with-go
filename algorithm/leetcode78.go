package algorithm

// 78. 子集
// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

// 说明：解集不能包含重复的子集。

// 示例:

// 输入: nums = [1,2,3]
// 输出:
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	for i := 0; i <= n; i++ {
		subsetsDFS(nums, 0, n, i, []int{}, &res)
	}
	return res
}

func subsetsDFS(nums []int, start, n, k int, cur []int, res *[][]int) {
	if len(cur) == k {
		newArr := make([]int, k)
		copy(newArr, cur)
		*res = append(*res, newArr)
		return
	}
	for i := start; i < n; i++ {
		cur = append(cur, nums[i])
		subsetsDFS(nums, i+1, n, k, cur, res)
		cur = cur[:len(cur) - 1]
	}
}