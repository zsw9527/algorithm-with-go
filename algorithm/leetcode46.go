package algorithm

import "algorithm-with-go/common"

// 46. 全排列
// 给定一个 没有重复 数字的序列，返回其所有可能的全排列。

// 示例:

// 输入: [1,2,3]
// 输出:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

func Permute(nums []int) [][]int {
	var res [][]int
	permuteDFS(nums, 0, &res)
	return res
}

func permuteDFS(nums []int, start int, res *[][]int) {
	if start == len(nums) {
		//这个地方为什么要将nums数组copy一份呢？
		//如果不copy，每一次递归都会修改nums数组，造成之前已经append到res里的nums数据发生变化
		//比如res里有一个nums是[1,3,2],后面交换nums的时候，也将同时修改已放入到res里面的nums,
		//这样再次放进res就会造成res数组的元素都是重复的现象
        tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return 
	}
	for i := start; i < len(nums); i++ {
		common.Swap(nums, start, i)
		permuteDFS(nums, start+1, res)
		common.Swap(nums, start, i)
	}
}

