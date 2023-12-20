package backtracking

import "sort"

// 给定一个整数集合candidates和一个目标数target, 找出candidates中所有可以使数字和为target的数字组合, candidates元素不允许重复使用
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	path := make([]int, 0, len(candidates))
	res := make([][]int, 0)
	backtrackSum2(candidates, target, 0, path, &res)
	return res
}

func backtrackSum2(candidates []int, target, start int, path []int, res *[][]int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			break // 剪枝
		}

		if i != start && candidates[i] == candidates[i-1] {
			continue // 树枝去重
		}

		path = append(path, candidates[i])
		backtrackSum2(candidates, target-candidates[i], i+1, path, res)
		path = path[:len(path)-1]
	}
}
