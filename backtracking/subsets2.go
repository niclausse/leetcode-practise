package backtracking

import "sort"

func subsets2(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	res := make([][]int, 0)
	sort.Ints(nums) // 去重需要排序
	backtrack(nums, 0, path, &res)
	return res
}

func backtrack(nums []int, start int, path []int, res *[][]int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	*res = append(*res, tmp)

	for i := start; i < len(nums); i++ {
		if i != start && nums[i] == nums[i-1] { // 剔除重复元素
			continue
		}

		path = append(path, nums[i])
		backtrack(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}
