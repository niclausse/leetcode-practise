package backtracking

import "sort"

func permute(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	res := make([][]int, 0)
	used := make(map[int]bool)
	dfsPermute(nums, path, &res, used)
	return res
}

func dfsPermute(nums []int, path []int, res *[][]int, used map[int]bool) {
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		path = append(path, nums[i])
		used[nums[i]] = true
		dfsPermute(nums, path, res, used)
		path = path[:len(path)-1]
		used[nums[i]] = false
	}
}

func permuteUnique(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	res := make([][]int, 0)
	indexUsed := make([]bool, len(nums))
	dfsPermuteUnique(nums, path, &res, indexUsed)
	return res
}

func dfsPermuteUnique(nums []int, path []int, res *[][]int, indexUsed []bool) {
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	used := make(map[int]bool) // 横向遍历时去重/剪枝 (可使用预排序替代)
	for i := 0; i < len(nums); i++ {
		// indexUsed用作纵向去重 used用作横向去重
		if used[nums[i]] || indexUsed[i] {
			continue
		}

		path = append(path, nums[i])
		indexUsed[i] = true
		used[nums[i]] = true
		dfsPermuteUnique(nums, path, res, indexUsed)
		path = path[:len(path)-1]
		indexUsed[i] = false
	}
}

func permuteUniqueBySort(nums []int) [][]int {
	sort.Ints(nums)
	path := make([]int, 0, len(nums))
	res := make([][]int, 0)
	indexUsed := make([]bool, len(nums))
	dfsPermuteUniqueBySort(nums, path, &res, indexUsed)
	return res
}

func dfsPermuteUniqueBySort(nums []int, path []int, res *[][]int, indexUsed []bool) {
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 纵向去重
		if indexUsed[i] {
			continue
		}

		// 横向去重 (剪枝)
		if i != 0 && nums[i] == nums[i-1] && !indexUsed[i-1] {
			continue
		}

		path = append(path, nums[i])
		indexUsed[i] = true
		dfsPermuteUniqueBySort(nums, path, res, indexUsed)
		path = path[:len(path)-1]
		indexUsed[i] = false
	}
}
